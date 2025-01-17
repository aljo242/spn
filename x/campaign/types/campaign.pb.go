// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: campaign/campaign.proto

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

type Campaign struct {
	CampaignID         uint64                                   `protobuf:"varint,1,opt,name=campaignID,proto3" json:"campaignID,omitempty"`
	CampaignName       string                                   `protobuf:"bytes,2,opt,name=campaignName,proto3" json:"campaignName,omitempty"`
	CoordinatorID      uint64                                   `protobuf:"varint,3,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
	MainnetID          uint64                                   `protobuf:"varint,4,opt,name=mainnetID,proto3" json:"mainnetID,omitempty"`
	MainnetInitialized bool                                     `protobuf:"varint,5,opt,name=mainnetInitialized,proto3" json:"mainnetInitialized,omitempty"`
	TotalSupply        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=totalSupply,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"totalSupply"`
	AllocatedShares    Shares                                   `protobuf:"bytes,7,rep,name=allocatedShares,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"allocatedShares"`
	DynamicShares      bool                                     `protobuf:"varint,8,opt,name=dynamicShares,proto3" json:"dynamicShares,omitempty"`
	TotalShares        Shares                                   `protobuf:"bytes,9,rep,name=totalShares,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"totalShares"`
}

func (m *Campaign) Reset()         { *m = Campaign{} }
func (m *Campaign) String() string { return proto.CompactTextString(m) }
func (*Campaign) ProtoMessage()    {}
func (*Campaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6f0d6f3906b81bb, []int{0}
}
func (m *Campaign) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Campaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Campaign.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Campaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Campaign.Merge(m, src)
}
func (m *Campaign) XXX_Size() int {
	return m.Size()
}
func (m *Campaign) XXX_DiscardUnknown() {
	xxx_messageInfo_Campaign.DiscardUnknown(m)
}

var xxx_messageInfo_Campaign proto.InternalMessageInfo

func (m *Campaign) GetCampaignID() uint64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

func (m *Campaign) GetCampaignName() string {
	if m != nil {
		return m.CampaignName
	}
	return ""
}

func (m *Campaign) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func (m *Campaign) GetMainnetID() uint64 {
	if m != nil {
		return m.MainnetID
	}
	return 0
}

func (m *Campaign) GetMainnetInitialized() bool {
	if m != nil {
		return m.MainnetInitialized
	}
	return false
}

func (m *Campaign) GetTotalSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

func (m *Campaign) GetAllocatedShares() Shares {
	if m != nil {
		return m.AllocatedShares
	}
	return nil
}

func (m *Campaign) GetDynamicShares() bool {
	if m != nil {
		return m.DynamicShares
	}
	return false
}

func (m *Campaign) GetTotalShares() Shares {
	if m != nil {
		return m.TotalShares
	}
	return nil
}

func init() {
	proto.RegisterType((*Campaign)(nil), "tendermint.spn.campaign.Campaign")
}

func init() { proto.RegisterFile("campaign/campaign.proto", fileDescriptor_f6f0d6f3906b81bb) }

var fileDescriptor_f6f0d6f3906b81bb = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbd, 0xae, 0xd3, 0x30,
	0x18, 0x8d, 0x69, 0x29, 0xad, 0x0b, 0x42, 0xb2, 0x90, 0x1a, 0x2a, 0xe4, 0x46, 0x15, 0x12, 0x11,
	0x12, 0xb6, 0x0a, 0x13, 0x6b, 0x9b, 0xa5, 0x0b, 0x43, 0xba, 0xc1, 0xe4, 0x24, 0x56, 0x6a, 0x91,
	0xd8, 0x51, 0xec, 0x56, 0x94, 0x81, 0x67, 0x60, 0xe4, 0x19, 0x78, 0x92, 0x8e, 0x1d, 0x99, 0x0a,
	0x6a, 0xdf, 0x02, 0x16, 0xd4, 0xfc, 0xd0, 0x14, 0x5d, 0xe9, 0xde, 0x3b, 0xdc, 0x29, 0xf6, 0xf9,
	0x4e, 0xce, 0x39, 0xdf, 0x97, 0x2f, 0x70, 0x10, 0xb2, 0x34, 0x63, 0x22, 0x96, 0xb4, 0x3e, 0x90,
	0x2c, 0x57, 0x46, 0xa1, 0x81, 0xe1, 0x32, 0xe2, 0x79, 0x2a, 0xa4, 0x21, 0x3a, 0x93, 0xa4, 0x2e,
	0x0f, 0x9f, 0xc4, 0x2a, 0x56, 0x05, 0x87, 0x9e, 0x4e, 0x25, 0x7d, 0x88, 0x43, 0xa5, 0x53, 0xa5,
	0x69, 0xc0, 0x34, 0xa7, 0xeb, 0x49, 0xc0, 0x0d, 0x9b, 0xd0, 0x50, 0x89, 0x4a, 0x6e, 0xfc, 0xa7,
	0x0d, 0xbb, 0xb3, 0x4a, 0x02, 0x61, 0x08, 0x6b, 0xb9, 0xb9, 0x67, 0x03, 0x07, 0xb8, 0x6d, 0xbf,
	0x81, 0xa0, 0x31, 0x7c, 0x58, 0xdf, 0xde, 0xb1, 0x94, 0xdb, 0xf7, 0x1c, 0xe0, 0xf6, 0xfc, 0x0b,
	0x0c, 0x3d, 0x87, 0x8f, 0x42, 0xa5, 0xf2, 0x48, 0x48, 0x66, 0x54, 0x3e, 0xf7, 0xec, 0x56, 0x21,
	0x73, 0x09, 0xa2, 0x67, 0xb0, 0x97, 0x32, 0x21, 0x25, 0x37, 0x73, 0xcf, 0x6e, 0x17, 0x8c, 0x33,
	0x80, 0x08, 0x44, 0xf5, 0x45, 0x0a, 0x23, 0x58, 0x22, 0x3e, 0xf3, 0xc8, 0xbe, 0xef, 0x00, 0xb7,
	0xeb, 0x5f, 0x51, 0x41, 0xdf, 0x00, 0xec, 0x1b, 0x65, 0x58, 0xb2, 0x58, 0x65, 0x59, 0xb2, 0xb1,
	0x3b, 0x4e, 0xcb, 0xed, 0xbf, 0x7e, 0x4a, 0xca, 0xde, 0xc9, 0xa9, 0x77, 0x52, 0xf5, 0x4e, 0x66,
	0x4a, 0xc8, 0xe9, 0x87, 0xed, 0x7e, 0x64, 0xfd, 0xde, 0x8f, 0x5e, 0xc4, 0xc2, 0x2c, 0x57, 0x01,
	0x09, 0x55, 0x4a, 0xab, 0x41, 0x95, 0x8f, 0x57, 0x3a, 0xfa, 0x48, 0xcd, 0x26, 0xe3, 0xba, 0x78,
	0xe1, 0xfb, 0xcf, 0x91, 0x7b, 0x43, 0xaa, 0xf6, 0x9b, 0x51, 0xd0, 0x17, 0xf8, 0x98, 0x25, 0x89,
	0x0a, 0x99, 0xe1, 0xd1, 0x62, 0xc9, 0x72, 0xae, 0xed, 0x07, 0xd7, 0xa5, 0x7b, 0x7b, 0xfb, 0x74,
	0x9d, 0x52, 0xdb, 0xff, 0xdf, 0xec, 0xf4, 0x39, 0xa2, 0x8d, 0x64, 0xa9, 0x08, 0x2b, 0xf7, 0x6e,
	0x31, 0xc5, 0x4b, 0x10, 0xad, 0xeb, 0xf9, 0x95, 0x9c, 0xde, 0x1d, 0x26, 0x6c, 0x1a, 0x4d, 0xbd,
	0xed, 0x01, 0x83, 0xdd, 0x01, 0x83, 0x5f, 0x07, 0x0c, 0xbe, 0x1e, 0xb1, 0xb5, 0x3b, 0x62, 0xeb,
	0xc7, 0x11, 0x5b, 0xef, 0x5f, 0x36, 0x94, 0xcf, 0x1b, 0x4f, 0x75, 0x26, 0xe9, 0xa7, 0x7f, 0xbf,
	0x44, 0xe9, 0x10, 0x74, 0x8a, 0x55, 0x7e, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x67, 0x40, 0xbb,
	0x5c, 0x34, 0x03, 0x00, 0x00,
}

func (m *Campaign) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Campaign) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Campaign) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalShares) > 0 {
		for iNdEx := len(m.TotalShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.DynamicShares {
		i--
		if m.DynamicShares {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.AllocatedShares) > 0 {
		for iNdEx := len(m.AllocatedShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllocatedShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.TotalSupply) > 0 {
		for iNdEx := len(m.TotalSupply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalSupply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MainnetInitialized {
		i--
		if m.MainnetInitialized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.MainnetID != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.MainnetID))
		i--
		dAtA[i] = 0x20
	}
	if m.CoordinatorID != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CampaignName) > 0 {
		i -= len(m.CampaignName)
		copy(dAtA[i:], m.CampaignName)
		i = encodeVarintCampaign(dAtA, i, uint64(len(m.CampaignName)))
		i--
		dAtA[i] = 0x12
	}
	if m.CampaignID != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.CampaignID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCampaign(dAtA []byte, offset int, v uint64) int {
	offset -= sovCampaign(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Campaign) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CampaignID != 0 {
		n += 1 + sovCampaign(uint64(m.CampaignID))
	}
	l = len(m.CampaignName)
	if l > 0 {
		n += 1 + l + sovCampaign(uint64(l))
	}
	if m.CoordinatorID != 0 {
		n += 1 + sovCampaign(uint64(m.CoordinatorID))
	}
	if m.MainnetID != 0 {
		n += 1 + sovCampaign(uint64(m.MainnetID))
	}
	if m.MainnetInitialized {
		n += 2
	}
	if len(m.TotalSupply) > 0 {
		for _, e := range m.TotalSupply {
			l = e.Size()
			n += 1 + l + sovCampaign(uint64(l))
		}
	}
	if len(m.AllocatedShares) > 0 {
		for _, e := range m.AllocatedShares {
			l = e.Size()
			n += 1 + l + sovCampaign(uint64(l))
		}
	}
	if m.DynamicShares {
		n += 2
	}
	if len(m.TotalShares) > 0 {
		for _, e := range m.TotalShares {
			l = e.Size()
			n += 1 + l + sovCampaign(uint64(l))
		}
	}
	return n
}

func sovCampaign(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCampaign(x uint64) (n int) {
	return sovCampaign(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Campaign) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCampaign
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
			return fmt.Errorf("proto: Campaign: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Campaign: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignID", wireType)
			}
			m.CampaignID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
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
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CampaignName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainnetID", wireType)
			}
			m.MainnetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainnetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainnetInitialized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
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
			m.MainnetInitialized = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
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
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalSupply = append(m.TotalSupply, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.TotalSupply[len(m.TotalSupply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatedShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
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
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatedShares = append(m.AllocatedShares, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.AllocatedShares[len(m.AllocatedShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicShares", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
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
			m.DynamicShares = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
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
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalShares = append(m.TotalShares, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.TotalShares[len(m.TotalShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCampaign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCampaign
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
func skipCampaign(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCampaign
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
					return 0, ErrIntOverflowCampaign
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
					return 0, ErrIntOverflowCampaign
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
				return 0, ErrInvalidLengthCampaign
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCampaign
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCampaign
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCampaign        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCampaign          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCampaign = fmt.Errorf("proto: unexpected end of group")
)
