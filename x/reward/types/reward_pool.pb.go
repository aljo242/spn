// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reward/reward_pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type RewardPool struct {
	LaunchID            uint64                                   `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
	Provider            string                                   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Coins               github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	LastRewardHeight    uint64                                   `protobuf:"varint,4,opt,name=lastRewardHeight,proto3" json:"lastRewardHeight,omitempty"`
	CurrentRewardHeight uint64                                   `protobuf:"varint,5,opt,name=currentRewardHeight,proto3" json:"currentRewardHeight,omitempty"`
}

func (m *RewardPool) Reset()         { *m = RewardPool{} }
func (m *RewardPool) String() string { return proto.CompactTextString(m) }
func (*RewardPool) ProtoMessage()    {}
func (*RewardPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_609e0d2ccc6b594f, []int{0}
}
func (m *RewardPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardPool.Merge(m, src)
}
func (m *RewardPool) XXX_Size() int {
	return m.Size()
}
func (m *RewardPool) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardPool.DiscardUnknown(m)
}

var xxx_messageInfo_RewardPool proto.InternalMessageInfo

func (m *RewardPool) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *RewardPool) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *RewardPool) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *RewardPool) GetLastRewardHeight() uint64 {
	if m != nil {
		return m.LastRewardHeight
	}
	return 0
}

func (m *RewardPool) GetCurrentRewardHeight() uint64 {
	if m != nil {
		return m.CurrentRewardHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*RewardPool)(nil), "tendermint.spn.reward.RewardPool")
}

func init() { proto.RegisterFile("reward/reward_pool.proto", fileDescriptor_609e0d2ccc6b594f) }

var fileDescriptor_609e0d2ccc6b594f = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3d, 0x4e, 0x33, 0x31,
	0x10, 0x86, 0xd7, 0xf9, 0xf9, 0xf4, 0x61, 0x1a, 0xb4, 0x80, 0xb4, 0xa4, 0x70, 0x22, 0x1a, 0x56,
	0x48, 0xd8, 0x04, 0x6e, 0x10, 0x28, 0xa0, 0x43, 0x5b, 0x42, 0x81, 0xf6, 0xc7, 0xda, 0xac, 0xd8,
	0x78, 0x2c, 0xdb, 0x09, 0x70, 0x07, 0x0a, 0xce, 0x41, 0xc7, 0x2d, 0x52, 0xa6, 0xa4, 0x0a, 0x28,
	0xb9, 0x05, 0x15, 0x5a, 0x7b, 0x15, 0x40, 0x50, 0x50, 0x8d, 0xc7, 0xf3, 0x78, 0xde, 0x77, 0x3c,
	0x38, 0x50, 0xfc, 0x36, 0x56, 0x19, 0x73, 0xe1, 0x5a, 0x02, 0x94, 0x54, 0x2a, 0x30, 0xe0, 0x6f,
	0x1b, 0x2e, 0x32, 0xae, 0x46, 0x85, 0x30, 0x54, 0x4b, 0x41, 0x1d, 0xd1, 0xd9, 0xca, 0x21, 0x07,
	0x4b, 0xb0, 0xea, 0xe4, 0xe0, 0x0e, 0x49, 0x41, 0x8f, 0x40, 0xb3, 0x24, 0xd6, 0x9c, 0x4d, 0xfa,
	0x09, 0x37, 0x71, 0x9f, 0xa5, 0x50, 0x08, 0x57, 0xdf, 0x7d, 0x6e, 0x60, 0x1c, 0xd9, 0x06, 0x17,
	0x00, 0xa5, 0xdf, 0xc1, 0xff, 0xcb, 0x78, 0x2c, 0xd2, 0xe1, 0xf9, 0x69, 0x80, 0x7a, 0x28, 0x6c,
	0x45, 0xab, 0xbc, 0xaa, 0x49, 0x05, 0x93, 0x22, 0xe3, 0x2a, 0x68, 0xf4, 0x50, 0xb8, 0x16, 0xad,
	0x72, 0xff, 0x01, 0xe1, 0x76, 0xd5, 0x55, 0x07, 0xcd, 0x5e, 0x33, 0x5c, 0x3f, 0xda, 0xa1, 0x4e,
	0x97, 0x56, 0xba, 0xb4, 0xd6, 0xa5, 0x27, 0x50, 0x88, 0xc1, 0xd5, 0x74, 0xde, 0xf5, 0xde, 0xe7,
	0xdd, 0xbd, 0xbc, 0x30, 0xc3, 0x71, 0x42, 0x53, 0x18, 0xb1, 0xda, 0xa4, 0x0b, 0x07, 0x3a, 0xbb,
	0x61, 0xe6, 0x5e, 0x72, 0x6d, 0x1f, 0x3c, 0xbd, 0x76, 0xc3, 0x3f, 0xa2, 0x3a, 0x72, 0x26, 0xfc,
	0x7d, 0xbc, 0x51, 0xc6, 0xda, 0xb8, 0xc1, 0xce, 0x78, 0x91, 0x0f, 0x4d, 0xd0, 0xb2, 0xe3, 0xfc,
	0xb8, 0xf7, 0x0f, 0xf1, 0x66, 0x3a, 0x56, 0x8a, 0x8b, 0xef, 0x78, 0xdb, 0xe2, 0xbf, 0x95, 0x06,
	0x83, 0xe9, 0x82, 0xa0, 0xd9, 0x82, 0xa0, 0xb7, 0x05, 0x41, 0x8f, 0x4b, 0xe2, 0xcd, 0x96, 0xc4,
	0x7b, 0x59, 0x12, 0xef, 0xf2, 0xab, 0xd1, 0xcf, 0x2d, 0x31, 0x2d, 0x05, 0xbb, 0xab, 0x37, 0xe9,
	0xec, 0x26, 0xff, 0xec, 0xf7, 0x1f, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x15, 0xd6, 0x7c, 0x6e,
	0xe7, 0x01, 0x00, 0x00,
}

func (m *RewardPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentRewardHeight != 0 {
		i = encodeVarintRewardPool(dAtA, i, uint64(m.CurrentRewardHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LastRewardHeight != 0 {
		i = encodeVarintRewardPool(dAtA, i, uint64(m.LastRewardHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRewardPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintRewardPool(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchID != 0 {
		i = encodeVarintRewardPool(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRewardPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewardPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovRewardPool(uint64(m.LaunchID))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovRewardPool(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovRewardPool(uint64(l))
		}
	}
	if m.LastRewardHeight != 0 {
		n += 1 + sovRewardPool(uint64(m.LastRewardHeight))
	}
	if m.CurrentRewardHeight != 0 {
		n += 1 + sovRewardPool(uint64(m.CurrentRewardHeight))
	}
	return n
}

func sovRewardPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewardPool(x uint64) (n int) {
	return sovRewardPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewardPool
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
			return fmt.Errorf("proto: RewardPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
				return ErrInvalidLengthRewardPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
				return ErrInvalidLengthRewardPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewardPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRewardHeight", wireType)
			}
			m.LastRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastRewardHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRewardHeight", wireType)
			}
			m.CurrentRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentRewardHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRewardPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewardPool
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
func skipRewardPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewardPool
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
					return 0, ErrIntOverflowRewardPool
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
					return 0, ErrIntOverflowRewardPool
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
				return 0, ErrInvalidLengthRewardPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewardPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewardPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewardPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewardPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewardPool = fmt.Errorf("proto: unexpected end of group")
)
