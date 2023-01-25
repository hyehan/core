package rest

import (
	"fmt"
	"net/http"

	"github.com/terra-money/core/x/oracle/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/gorilla/mux"
)

func registerTxHandlers(cliCtx client.Context, rtr *mux.Router) {
	rtr.HandleFunc(fmt.Sprintf("/oracle/voters/{%s}/feeder", RestVoter), newDelegateHandlerFunction(cliCtx)).Methods("POST")
	rtr.HandleFunc(fmt.Sprintf("/oracle/voters/{%s}/aggregate_prevote", RestVoter), newAggregatePrevoteHandlerFunction(cliCtx)).Methods("POST")
	rtr.HandleFunc(fmt.Sprintf("/oracle/voters/{%s}/aggregate_vote", RestVoter), newAggregateVoteHandlerFunction(cliCtx)).Methods("POST")
}

type (
	delegateReq struct {
		BaseReq rest.BaseReq   `json:"base_req" yaml:"base_req"`
		Feeder  sdk.AccAddress `json:"feeder" yaml:"feeder"`
	}

	aggregatePrevoteReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Hash          string `json:"hash" yaml:"hash"`
		ExchangeRates string `json:"exchange_rates" yaml:"exchange_rates"`
		Salt          string `json:"salt" yaml:"salt"`
	}

	aggregateVoteReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		ExchangeRates string `json:"exchange_rates" yaml:"exchange_rates"`
		Salt          string `json:"salt" yaml:"salt"`
	}
)

func newDelegateHandlerFunction(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req delegateReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		voterAddr, ok := checkVoterAddressVar(w, r)
		if !ok {
			return
		}

		// create the message
		msg := types.NewMsgDelegateFeedConsent(voterAddr, req.Feeder)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newAggregatePrevoteHandlerFunction(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req aggregatePrevoteReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		feederAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		voterAddr, ok := checkVoterAddressVar(w, r)
		if !ok {
			return
		}

		var hash types.AggregateVoteHash

		// If hash is not given, then retrieve hash from exchange_rate and salt
		if len(req.Hash) == 0 && (len(req.ExchangeRates) > 0 && len(req.Salt) > 0) { //nolint:gocritic
			_, err := types.ParseExchangeRateTuples(req.ExchangeRates)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			hash = types.GetAggregateVoteHash(req.Salt, req.ExchangeRates, voterAddr)
		} else if len(req.Hash) > 0 {
			hash, err = types.AggregateVoteHashFromHexString(req.Hash)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "must provide Hash or (ExchangeRates & Salt)")
			return
		}

		// create the message
		msg := types.NewMsgAggregateExchangeRatePrevote(hash, feederAddr, voterAddr)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newAggregateVoteHandlerFunction(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req aggregateVoteReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		feederAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		voterAddr, ok := checkVoterAddressVar(w, r)
		if !ok {
			return
		}

		// Check validation of tuples
		_, err = types.ParseExchangeRateTuples(req.ExchangeRates)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgAggregateExchangeRateVote(req.Salt, req.ExchangeRates, feederAddr, voterAddr)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
