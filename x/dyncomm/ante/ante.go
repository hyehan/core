package ante

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	dyncommkeeper "github.com/classic-terra/core/v3/x/dyncomm/keeper"
)

// DyncommDecorator checks for EditValidator and rejects
// edits that do not conform with dyncomm
type DyncommDecorator struct {
	dyncommKeeper dyncommkeeper.Keeper
	stakingKeeper *stakingkeeper.Keeper
	cdc           codec.BinaryCodec
}

func NewDyncommDecorator(cdc codec.BinaryCodec, dk dyncommkeeper.Keeper, sk *stakingkeeper.Keeper) DyncommDecorator {
	return DyncommDecorator{
		dyncommKeeper: dk,
		stakingKeeper: sk,
		cdc:           cdc,
	}
}

func (dd DyncommDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if simulate {
		return next(ctx, tx, simulate)
	}

	msgs := tx.GetMsgs()
	err := dd.FilterMsgsAndProcessMsgs(ctx, msgs...)
	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func (dd DyncommDecorator) FilterMsgsAndProcessMsgs(ctx sdk.Context, msgs ...sdk.Msg) (err error) {
	for _, msg := range msgs {

		switch msg := msg.(type) {
		case *stakingtypes.MsgEditValidator:
			err = dd.ProcessEditValidator(ctx, msg)
		case *authz.MsgExec:
			messages, msgerr := msg.GetMessages()
			if msgerr == nil {
				err = dd.FilterMsgsAndProcessMsgs(ctx, messages...)
			}
		case *channeltypes.MsgRecvPacket:
			var data icatypes.InterchainAccountPacketData
			err = icatypes.ModuleCdc.UnmarshalJSON(msg.Packet.GetData(), &data)
			if err != nil {
				continue
			}
			if data.Type != icatypes.EXECUTE_TX {
				continue
			}
			messages, msgerr := icatypes.DeserializeCosmosTx(dd.cdc, data.Data)
			if msgerr == nil {
				err = dd.FilterMsgsAndProcessMsgs(ctx, messages...)
			}
		default:
			continue
		}

		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, err.Error())
		}

	}
	return nil
}

func (dd DyncommDecorator) ProcessEditValidator(ctx sdk.Context, msg sdk.Msg) (err error) {
	msgEditValidator := msg.(*stakingtypes.MsgEditValidator)

	// no update of CommissionRate provided
	if msgEditValidator.CommissionRate == nil {
		return nil
	}

	operator := msgEditValidator.ValidatorAddress
	newIntendedRate := msgEditValidator.CommissionRate
	dynMinRate := dd.dyncommKeeper.GetDynCommissionRate(ctx, operator)

	if newIntendedRate.LT(dynMinRate) {
		return fmt.Errorf("commission for %s must be at least %f", operator, dynMinRate.MustFloat64())
	}

	return nil
}
