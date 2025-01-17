// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringc/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the monitoringc module's genesis state.
type GenesisState struct {
	Params                           Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId                           string                         `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	VerifiedClientIDList             []VerifiedClientID             `protobuf:"bytes,3,rep,name=verifiedClientIDList,proto3" json:"verifiedClientIDList"`
	ProviderClientIDList             []ProviderClientID             `protobuf:"bytes,4,rep,name=providerClientIDList,proto3" json:"providerClientIDList"`
	LaunchIDFromVerifiedClientIDList []LaunchIDFromVerifiedClientID `protobuf:"bytes,5,rep,name=launchIDFromVerifiedClientIDList,proto3" json:"launchIDFromVerifiedClientIDList"`
	LaunchIDFromChannelIDList        []LaunchIDFromChannelID        `protobuf:"bytes,6,rep,name=launchIDFromChannelIDList,proto3" json:"launchIDFromChannelIDList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c269cb89bd5f03, []int{0}
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

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetVerifiedClientIDList() []VerifiedClientID {
	if m != nil {
		return m.VerifiedClientIDList
	}
	return nil
}

func (m *GenesisState) GetProviderClientIDList() []ProviderClientID {
	if m != nil {
		return m.ProviderClientIDList
	}
	return nil
}

func (m *GenesisState) GetLaunchIDFromVerifiedClientIDList() []LaunchIDFromVerifiedClientID {
	if m != nil {
		return m.LaunchIDFromVerifiedClientIDList
	}
	return nil
}

func (m *GenesisState) GetLaunchIDFromChannelIDList() []LaunchIDFromChannelID {
	if m != nil {
		return m.LaunchIDFromChannelIDList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.monitoringc.GenesisState")
}

func init() { proto.RegisterFile("monitoringc/genesis.proto", fileDescriptor_10c269cb89bd5f03) }

var fileDescriptor_10c269cb89bd5f03 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4f, 0xc2, 0x40,
	0x00, 0xc5, 0x5b, 0xc1, 0x1a, 0x0f, 0xa7, 0x86, 0xc4, 0xd2, 0xa1, 0x36, 0xc4, 0x81, 0x44, 0x6d,
	0x23, 0x2c, 0x8e, 0x06, 0x88, 0x84, 0x84, 0x81, 0x60, 0xc2, 0xe0, 0xd2, 0x94, 0xf6, 0x28, 0x97,
	0xd0, 0xbb, 0xcb, 0xf5, 0x20, 0xfa, 0x15, 0x9c, 0xfc, 0x42, 0xee, 0x8c, 0x8c, 0x4e, 0xc6, 0xc0,
	0x17, 0x31, 0xfd, 0x03, 0x56, 0x2c, 0xc5, 0xad, 0x4d, 0xdf, 0x7b, 0xbf, 0x77, 0xaf, 0x07, 0x2a,
	0x3e, 0xc1, 0x88, 0x13, 0x86, 0xb0, 0xe7, 0x98, 0x1e, 0xc4, 0x30, 0x40, 0x81, 0x41, 0x19, 0xe1,
	0x44, 0x56, 0x39, 0xc4, 0x2e, 0x64, 0x3e, 0xc2, 0xdc, 0x08, 0x28, 0x36, 0x52, 0x4a, 0xb5, 0xec,
	0x11, 0x8f, 0x44, 0x32, 0x33, 0x7c, 0x8a, 0x1d, 0xaa, 0x92, 0x0e, 0xa3, 0x36, 0xb3, 0xfd, 0x24,
	0x4b, 0xbd, 0x4c, 0x7f, 0x99, 0x43, 0x86, 0xc6, 0x08, 0xba, 0x96, 0x33, 0x45, 0x10, 0x73, 0x0b,
	0xb9, 0x59, 0x2a, 0xca, 0xc8, 0x1c, 0xb9, 0x90, 0xfd, 0x51, 0x35, 0xd2, 0xaa, 0xa9, 0x3d, 0xc3,
	0xce, 0xc4, 0x42, 0xae, 0x35, 0x66, 0xc4, 0xb7, 0xf6, 0x46, 0x5f, 0xe5, 0x98, 0x9c, 0x89, 0x8d,
	0x31, 0x9c, 0x6e, 0xc5, 0xd5, 0xf7, 0x22, 0x38, 0xeb, 0xc4, 0x5b, 0x3c, 0x72, 0x9b, 0x43, 0xf9,
	0x1e, 0x48, 0xf1, 0x71, 0x14, 0x51, 0x17, 0x6b, 0xa5, 0x7a, 0xd5, 0xd8, 0xbf, 0x8d, 0xd1, 0x8f,
	0x94, 0xcd, 0xe2, 0xe2, 0xf3, 0x42, 0x18, 0x24, 0x3e, 0xf9, 0x1c, 0x9c, 0x50, 0xc2, 0xc2, 0x42,
	0xca, 0x91, 0x2e, 0xd6, 0x4e, 0x07, 0x52, 0xf8, 0xda, 0x75, 0xe5, 0x31, 0x28, 0x6f, 0x4a, 0xb7,
	0xa2, 0xce, 0xdd, 0x76, 0x0f, 0x05, 0x5c, 0x29, 0xe8, 0x85, 0x5a, 0xa9, 0x7e, 0x9d, 0x07, 0x1a,
	0xee, 0xf8, 0x12, 0x64, 0x66, 0x5e, 0xc8, 0xd9, 0x2c, 0xfa, 0x8b, 0x53, 0x3c, 0xcc, 0xe9, 0xef,
	0xf8, 0x36, 0x9c, 0xac, 0x3c, 0xf9, 0x55, 0x04, 0x7a, 0xbc, 0x6f, 0xb7, 0xfd, 0xc0, 0x88, 0x3f,
	0xcc, 0x3a, 0xdc, 0x71, 0x04, 0xbd, 0xcb, 0x83, 0xf6, 0x72, 0x32, 0x92, 0x02, 0x07, 0x39, 0xf2,
	0x0c, 0x54, 0xd2, 0x9a, 0x56, 0xfc, 0xa3, 0x93, 0x12, 0x52, 0x54, 0xe2, 0xf6, 0xbf, 0x25, 0xb6,
	0xe6, 0x84, 0xbe, 0x3f, 0xb9, 0xd9, 0x59, 0xac, 0x34, 0x71, 0xb9, 0xd2, 0xc4, 0xaf, 0x95, 0x26,
	0xbe, 0xad, 0x35, 0x61, 0xb9, 0xd6, 0x84, 0x8f, 0xb5, 0x26, 0x3c, 0xdd, 0x78, 0x88, 0x4f, 0x66,
	0x23, 0xc3, 0x21, 0xbe, 0xf9, 0xc3, 0x35, 0x03, 0x8a, 0xcd, 0x67, 0x33, 0x7d, 0x45, 0xf9, 0x0b,
	0x85, 0xc1, 0x48, 0x8a, 0xee, 0x63, 0xe3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x2f, 0x75, 0x39,
	0xa6, 0x03, 0x00, 0x00,
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
	if len(m.LaunchIDFromChannelIDList) > 0 {
		for iNdEx := len(m.LaunchIDFromChannelIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDFromChannelIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LaunchIDFromVerifiedClientIDList) > 0 {
		for iNdEx := len(m.LaunchIDFromVerifiedClientIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDFromVerifiedClientIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ProviderClientIDList) > 0 {
		for iNdEx := len(m.ProviderClientIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderClientIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.VerifiedClientIDList) > 0 {
		for iNdEx := len(m.VerifiedClientIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerifiedClientIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.VerifiedClientIDList) > 0 {
		for _, e := range m.VerifiedClientIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProviderClientIDList) > 0 {
		for _, e := range m.ProviderClientIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDFromVerifiedClientIDList) > 0 {
		for _, e := range m.LaunchIDFromVerifiedClientIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDFromChannelIDList) > 0 {
		for _, e := range m.LaunchIDFromChannelIDList {
			l = e.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiedClientIDList", wireType)
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
			m.VerifiedClientIDList = append(m.VerifiedClientIDList, VerifiedClientID{})
			if err := m.VerifiedClientIDList[len(m.VerifiedClientIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderClientIDList", wireType)
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
			m.ProviderClientIDList = append(m.ProviderClientIDList, ProviderClientID{})
			if err := m.ProviderClientIDList[len(m.ProviderClientIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDFromVerifiedClientIDList", wireType)
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
			m.LaunchIDFromVerifiedClientIDList = append(m.LaunchIDFromVerifiedClientIDList, LaunchIDFromVerifiedClientID{})
			if err := m.LaunchIDFromVerifiedClientIDList[len(m.LaunchIDFromVerifiedClientIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDFromChannelIDList", wireType)
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
			m.LaunchIDFromChannelIDList = append(m.LaunchIDFromChannelIDList, LaunchIDFromChannelID{})
			if err := m.LaunchIDFromChannelIDList[len(m.LaunchIDFromChannelIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
