// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/feeshare/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// params are the feeshare module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// FeeShare is a slice of active registered contracts for fee distribution
	FeeShare []FeeShare `protobuf:"bytes,2,rep,name=fee_share,json=feeShare,proto3" json:"fee_share"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6781086762d348, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeeShare() []FeeShare {
	if m != nil {
		return m.FeeShare
	}
	return nil
}

// Params defines the feeshare module params
type Params struct {
	// enable_feeshare defines a parameter to enable the feeshare module
	EnableFeeShare bool `protobuf:"varint,1,opt,name=enable_fee_share,json=enableFeeShare,proto3" json:"enable_fee_share,omitempty"`
	// developer_shares defines the proportion of the transaction fees to be
	// distributed to the registered contract owner
	DeveloperShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=developer_shares,json=developerShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_shares"`
	// allowed_denoms defines the list of denoms that are allowed to be paid to
	// the contract withdraw addresses. If said denom is not in the list, the fees
	// will ONLY be sent to the community pool.
	// If this list is empty, all denoms are allowed.
	AllowedDenoms []string `protobuf:"bytes,3,rep,name=allowed_denoms,json=allowedDenoms,proto3" json:"allowed_denoms,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6781086762d348, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableFeeShare() bool {
	if m != nil {
		return m.EnableFeeShare
	}
	return false
}

func (m *Params) GetAllowedDenoms() []string {
	if m != nil {
		return m.AllowedDenoms
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "terra.feeshare.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "terra.feeshare.v1beta1.Params")
}

func init() {
	proto.RegisterFile("terra/feeshare/v1beta1/genesis.proto", fileDescriptor_0b6781086762d348)
}

var fileDescriptor_0b6781086762d348 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x3b, 0x70, 0x43, 0x60, 0xb8, 0x22, 0x69, 0x8c, 0x69, 0x58, 0x94, 0x86, 0x88, 0xe9,
	0x86, 0x69, 0xc0, 0xad, 0x2b, 0x24, 0xea, 0xd2, 0x94, 0x95, 0x6e, 0xc8, 0xb4, 0x3d, 0x14, 0x62,
	0xcb, 0x34, 0x33, 0x23, 0xea, 0x5b, 0xe8, 0xbb, 0xf8, 0x10, 0x2c, 0x59, 0x1a, 0x17, 0xc4, 0xc0,
	0x8b, 0x98, 0x4e, 0x0b, 0xb8, 0x90, 0x55, 0x3b, 0x7f, 0xbe, 0xf3, 0x9d, 0x93, 0xfc, 0xf8, 0x4c,
	0x02, 0xe7, 0xd4, 0x19, 0x03, 0x88, 0x09, 0xe5, 0xe0, 0xcc, 0xbb, 0x1e, 0x48, 0xda, 0x75, 0x42,
	0x98, 0x81, 0x98, 0x0a, 0x92, 0x70, 0x26, 0x99, 0x7e, 0xaa, 0x28, 0xb2, 0xa5, 0x48, 0x4e, 0x35,
	0x4e, 0x42, 0x16, 0x32, 0x85, 0x38, 0xe9, 0x5f, 0x46, 0x37, 0xda, 0x07, 0x9c, 0xbb, 0x71, 0x85,
	0xb5, 0xde, 0x11, 0xfe, 0x7f, 0x93, 0xad, 0x19, 0x4a, 0x2a, 0x41, 0xbf, 0xc4, 0xa5, 0x84, 0x72,
	0x1a, 0x0b, 0x03, 0x59, 0xc8, 0xae, 0xf6, 0x4c, 0xf2, 0xf7, 0x5a, 0x72, 0xa7, 0xa8, 0xfe, 0xbf,
	0xc5, 0xaa, 0xa9, 0xb9, 0xf9, 0x8c, 0x7e, 0x85, 0x2b, 0x63, 0x80, 0x91, 0x22, 0x8d, 0x82, 0x55,
	0xb4, 0xab, 0x3d, 0xeb, 0x90, 0xe0, 0x1a, 0x60, 0x98, 0x06, 0xb9, 0xa2, 0x3c, 0xce, 0xdf, 0xad,
	0x0f, 0x84, 0x4b, 0x99, 0x5d, 0xb7, 0x71, 0x1d, 0x66, 0xd4, 0x8b, 0x60, 0xb4, 0xd7, 0xa6, 0x77,
	0x95, 0xdd, 0x5a, 0x96, 0x6f, 0x25, 0xfa, 0x3d, 0xae, 0x07, 0x30, 0x87, 0x88, 0x25, 0xc0, 0x33,
	0x50, 0x18, 0x05, 0x0b, 0xd9, 0x95, 0x3e, 0x49, 0xf5, 0x5f, 0xab, 0xe6, 0x79, 0x38, 0x95, 0x93,
	0x27, 0x8f, 0xf8, 0x2c, 0x76, 0x7c, 0x26, 0x62, 0x26, 0xf2, 0x4f, 0x47, 0x04, 0x8f, 0x8e, 0x7c,
	0x4d, 0x40, 0x90, 0x01, 0xf8, 0xee, 0xf1, 0xce, 0xa3, 0xcc, 0x42, 0x6f, 0xe3, 0x1a, 0x8d, 0x22,
	0xf6, 0x0c, 0xc1, 0x28, 0x80, 0x19, 0x8b, 0x85, 0x51, 0xb4, 0x8a, 0x76, 0xc5, 0x3d, 0xca, 0xd3,
	0x81, 0x0a, 0xfb, 0xb7, 0x8b, 0xb5, 0x89, 0x96, 0x6b, 0x13, 0x7d, 0xaf, 0x4d, 0xf4, 0xb6, 0x31,
	0xb5, 0xe5, 0xc6, 0xd4, 0x3e, 0x37, 0xa6, 0xf6, 0x40, 0x7e, 0x6f, 0x8e, 0xa8, 0x10, 0x53, 0xbf,
	0x93, 0xd5, 0xe3, 0x33, 0x0e, 0xce, 0xcb, 0xbe, 0x25, 0x75, 0x85, 0x57, 0x52, 0xdd, 0x5c, 0xfc,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x95, 0x2d, 0x58, 0xfd, 0x18, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeShare) > 0 {
		for iNdEx := len(m.FeeShare) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeShare[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedDenoms) > 0 {
		for iNdEx := len(m.AllowedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedDenoms[iNdEx])
			copy(dAtA[i:], m.AllowedDenoms[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.AllowedDenoms[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DeveloperShares) > 0 {
		i -= len(m.DeveloperShares)
		copy(dAtA[i:], m.DeveloperShares)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DeveloperShares)))
		i--
		dAtA[i] = 0x12
	}
	if m.EnableFeeShare {
		i--
		if m.EnableFeeShare {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeeShare) > 0 {
		for _, e := range m.FeeShare {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableFeeShare {
		n += 2
	}
	l = len(m.DeveloperShares)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AllowedDenoms) > 0 {
		for _, s := range m.AllowedDenoms {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeShare", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeShare = append(m.FeeShare, FeeShare{})
			if err := m.FeeShare[len(m.FeeShare)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableFeeShare", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableFeeShare = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeveloperShares = github_com_cosmos_cosmos_sdk_types.Dec(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedDenoms = append(m.AllowedDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
