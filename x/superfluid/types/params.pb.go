// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the superfluid module
type Params struct {
	// refresh epoch identifier for staked amount
	RefreshEpochIdentifier string `protobuf:"bytes,1,opt,name=refresh_epoch_identifier,json=refreshEpochIdentifier,proto3" json:"refresh_epoch_identifier,omitempty" yaml:"refresh_epoch_identifier"`
	// the risk_factor is to be cut on OSMO equivalent value of lp tokens for superfluid staking, default: 5%
	MinimumRiskFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=minimum_risk_factor,json=minimumRiskFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minimum_risk_factor" yaml:"minimum_risk_factor"`
	// unbonding duration of superfluid delegation
	UnbondingDuration time.Duration `protobuf:"bytes,3,opt,name=unbonding_duration,json=unbondingDuration,proto3,stdduration" json:"unbonding_duration" yaml:"unbonding_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0985261dfaf2a82e, []int{0}
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

func (m *Params) GetRefreshEpochIdentifier() string {
	if m != nil {
		return m.RefreshEpochIdentifier
	}
	return ""
}

func (m *Params) GetUnbondingDuration() time.Duration {
	if m != nil {
		return m.UnbondingDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.superfluid.Params")
}

func init() { proto.RegisterFile("osmosis/superfluid/params.proto", fileDescriptor_0985261dfaf2a82e) }

var fileDescriptor_0985261dfaf2a82e = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xc7, 0x93, 0x5e, 0x28, 0xdc, 0xdc, 0x55, 0x73, 0x2f, 0x97, 0xb4, 0x8b, 0x4c, 0x89, 0x28,
	0xdd, 0x34, 0x03, 0x75, 0xe7, 0xb2, 0x54, 0x41, 0xe8, 0x42, 0xb2, 0x14, 0x24, 0xe4, 0x63, 0x92,
	0x0e, 0xcd, 0xe4, 0x84, 0x99, 0x0c, 0x58, 0xf0, 0x21, 0x5c, 0xfa, 0x48, 0x5d, 0x76, 0x25, 0xe2,
	0x22, 0x4a, 0xfb, 0x06, 0x7d, 0x02, 0xe9, 0x24, 0xad, 0x05, 0x75, 0x35, 0x73, 0xce, 0xef, 0x3f,
	0xff, 0x39, 0x1f, 0x06, 0x02, 0xc1, 0x40, 0x50, 0x81, 0x85, 0x2c, 0x08, 0x4f, 0x32, 0x49, 0x63,
	0x5c, 0x04, 0x3c, 0x60, 0xc2, 0x2d, 0x38, 0x94, 0x60, 0x9a, 0x8d, 0xc0, 0xfd, 0x14, 0xf4, 0xfe,
	0xa5, 0x90, 0x82, 0xc2, 0x78, 0x77, 0xab, 0x95, 0x3d, 0x3b, 0x05, 0x48, 0x33, 0x82, 0x55, 0x14,
	0xca, 0x04, 0xc7, 0x92, 0x07, 0x25, 0x85, 0xbc, 0xe6, 0xce, 0x73, 0xcb, 0x68, 0xdf, 0x28, 0x6b,
	0xf3, 0xce, 0xb0, 0x38, 0x49, 0x38, 0x11, 0x33, 0x9f, 0x14, 0x10, 0xcd, 0x7c, 0x1a, 0x93, 0xbc,
	0xa4, 0x09, 0x25, 0xdc, 0xd2, 0xfb, 0xfa, 0xe0, 0xf7, 0xf8, 0x64, 0x5b, 0x21, 0xb4, 0x08, 0x58,
	0x76, 0xe1, 0xfc, 0xa4, 0x74, 0xbc, 0xff, 0x0d, 0xba, 0xdc, 0x91, 0xeb, 0x03, 0x30, 0x1f, 0x8c,
	0xbf, 0x8c, 0xe6, 0x94, 0x49, 0xe6, 0x73, 0x2a, 0xe6, 0x7e, 0x12, 0x44, 0x25, 0x70, 0xab, 0xa5,
	0x9c, 0xa7, 0xcb, 0x0a, 0x69, 0xaf, 0x15, 0x3a, 0x4b, 0x69, 0x39, 0x93, 0xa1, 0x1b, 0x01, 0xc3,
	0x91, 0x6a, 0xb2, 0x39, 0x86, 0x22, 0x9e, 0xe3, 0x72, 0x51, 0x10, 0xe1, 0x4e, 0x48, 0xb4, 0xad,
	0x50, 0xaf, 0xae, 0xe3, 0x1b, 0x4b, 0xc7, 0xeb, 0x34, 0x59, 0x8f, 0x8a, 0xf9, 0x95, 0xca, 0x99,
	0x60, 0x98, 0x32, 0x0f, 0x21, 0x8f, 0x69, 0x9e, 0xfa, 0xfb, 0x19, 0x58, 0xbf, 0xfa, 0xfa, 0xe0,
	0xcf, 0xa8, 0xeb, 0xd6, 0x43, 0x72, 0xf7, 0x43, 0x72, 0x27, 0x8d, 0x60, 0x7c, 0xba, 0xab, 0x6b,
	0x5b, 0xa1, 0x6e, 0xfd, 0xdb, 0x57, 0x0b, 0xe7, 0xe9, 0x0d, 0xe9, 0x5e, 0xe7, 0x00, 0x0e, 0x2f,
	0xa7, 0xcb, 0xb5, 0xad, 0xaf, 0xd6, 0xb6, 0xfe, 0xbe, 0xb6, 0xf5, 0xc7, 0x8d, 0xad, 0xad, 0x36,
	0xb6, 0xf6, 0xb2, 0xb1, 0xb5, 0xdb, 0xd1, 0x51, 0x8f, 0xcd, 0x1e, 0x87, 0x59, 0x10, 0x8a, 0x7d,
	0x80, 0xef, 0x8f, 0xf7, 0xae, 0x7a, 0x0e, 0xdb, 0xaa, 0xb4, 0xf3, 0x8f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x84, 0x07, 0xd8, 0x1a, 0x02, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.UnbondingDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnbondingDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinimumRiskFactor.Size()
		i -= size
		if _, err := m.MinimumRiskFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RefreshEpochIdentifier) > 0 {
		i -= len(m.RefreshEpochIdentifier)
		copy(dAtA[i:], m.RefreshEpochIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.RefreshEpochIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RefreshEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinimumRiskFactor.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnbondingDuration)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefreshEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumRiskFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumRiskFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.UnbondingDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
