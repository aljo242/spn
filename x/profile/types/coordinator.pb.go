// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/coordinator.proto

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

type Coordinator struct {
	CoordinatorID uint64                 `protobuf:"varint,1,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Description   CoordinatorDescription `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
}

func (m *Coordinator) Reset()         { *m = Coordinator{} }
func (m *Coordinator) String() string { return proto.CompactTextString(m) }
func (*Coordinator) ProtoMessage()    {}
func (*Coordinator) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b9115302ae8cca0, []int{0}
}
func (m *Coordinator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coordinator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coordinator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Coordinator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinator.Merge(m, src)
}
func (m *Coordinator) XXX_Size() int {
	return m.Size()
}
func (m *Coordinator) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinator.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinator proto.InternalMessageInfo

func (m *Coordinator) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func (m *Coordinator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Coordinator) GetDescription() CoordinatorDescription {
	if m != nil {
		return m.Description
	}
	return CoordinatorDescription{}
}

type CoordinatorDescription struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Website  string `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	Details  string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *CoordinatorDescription) Reset()         { *m = CoordinatorDescription{} }
func (m *CoordinatorDescription) String() string { return proto.CompactTextString(m) }
func (*CoordinatorDescription) ProtoMessage()    {}
func (*CoordinatorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b9115302ae8cca0, []int{1}
}
func (m *CoordinatorDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoordinatorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoordinatorDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoordinatorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorDescription.Merge(m, src)
}
func (m *CoordinatorDescription) XXX_Size() int {
	return m.Size()
}
func (m *CoordinatorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorDescription proto.InternalMessageInfo

func (m *CoordinatorDescription) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CoordinatorDescription) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *CoordinatorDescription) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type CoordinatorByAddress struct {
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CoordinatorID uint64 `protobuf:"varint,2,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
}

func (m *CoordinatorByAddress) Reset()         { *m = CoordinatorByAddress{} }
func (m *CoordinatorByAddress) String() string { return proto.CompactTextString(m) }
func (*CoordinatorByAddress) ProtoMessage()    {}
func (*CoordinatorByAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b9115302ae8cca0, []int{2}
}
func (m *CoordinatorByAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoordinatorByAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoordinatorByAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoordinatorByAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorByAddress.Merge(m, src)
}
func (m *CoordinatorByAddress) XXX_Size() int {
	return m.Size()
}
func (m *CoordinatorByAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorByAddress.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorByAddress proto.InternalMessageInfo

func (m *CoordinatorByAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CoordinatorByAddress) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func init() {
	proto.RegisterType((*Coordinator)(nil), "tendermint.spn.profile.Coordinator")
	proto.RegisterType((*CoordinatorDescription)(nil), "tendermint.spn.profile.CoordinatorDescription")
	proto.RegisterType((*CoordinatorByAddress)(nil), "tendermint.spn.profile.CoordinatorByAddress")
}

func init() { proto.RegisterFile("profile/coordinator.proto", fileDescriptor_8b9115302ae8cca0) }

var fileDescriptor_8b9115302ae8cca0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xb1, 0x4e, 0xf3, 0x30,
	0x18, 0x8c, 0xfb, 0x57, 0x3f, 0xd4, 0x15, 0x4b, 0x54, 0x55, 0xa1, 0x83, 0xa9, 0x22, 0x86, 0xb0,
	0xd8, 0x12, 0x3c, 0x01, 0x69, 0x17, 0xd6, 0x0c, 0x1d, 0xd8, 0x92, 0xd8, 0xa4, 0x96, 0x5a, 0xdb,
	0xb2, 0x8d, 0x20, 0x6f, 0xc1, 0x4b, 0xf0, 0x2e, 0x1d, 0x3b, 0x32, 0x21, 0x94, 0xbc, 0x08, 0x4a,
	0x48, 0x48, 0x2a, 0xb2, 0xf9, 0xec, 0xf3, 0x7d, 0x77, 0xf7, 0xc1, 0x4b, 0xa5, 0xe5, 0x13, 0xdf,
	0x31, 0x92, 0x4a, 0xa9, 0x29, 0x17, 0xb1, 0x95, 0x1a, 0x2b, 0x2d, 0xad, 0x74, 0xe7, 0x96, 0x09,
	0xca, 0xf4, 0x9e, 0x0b, 0x8b, 0x8d, 0x12, 0xb8, 0x61, 0x2e, 0x66, 0x99, 0xcc, 0x64, 0x4d, 0x21,
	0xd5, 0xe9, 0x87, 0xed, 0xbf, 0x03, 0x38, 0x5d, 0x75, 0x1a, 0xee, 0x35, 0xbc, 0xe8, 0x49, 0x3e,
	0xac, 0x3d, 0xb0, 0x04, 0xc1, 0x38, 0x3a, 0xbd, 0x74, 0x3d, 0x78, 0x16, 0x53, 0xaa, 0x99, 0x31,
	0xde, 0x68, 0x09, 0x82, 0x49, 0xd4, 0x42, 0x77, 0x03, 0xa7, 0x94, 0x99, 0x54, 0x73, 0x65, 0xb9,
	0x14, 0xde, 0xbf, 0x25, 0x08, 0xa6, 0xb7, 0x18, 0x0f, 0x7b, 0xc2, 0xbd, 0xc9, 0xeb, 0xee, 0x57,
	0x38, 0x3e, 0x7c, 0x5e, 0x39, 0x51, 0x5f, 0xc8, 0xdf, 0xc2, 0xf9, 0x30, 0xd9, 0x5d, 0xc0, 0x73,
	0x4e, 0x99, 0xb0, 0xdc, 0xe6, 0xb5, 0xd9, 0x49, 0xf4, 0x8b, 0x2b, 0x9f, 0x2f, 0x2c, 0x31, 0xdc,
	0xb2, 0xd6, 0x67, 0x03, 0xab, 0x17, 0xca, 0x6c, 0xcc, 0x77, 0xa6, 0xf6, 0x38, 0x89, 0x5a, 0xe8,
	0x6f, 0xe0, 0xac, 0x37, 0x29, 0xcc, 0xef, 0x9b, 0x64, 0xbd, 0xcc, 0xe0, 0x34, 0xf3, 0x9f, 0xce,
	0x46, 0x03, 0x9d, 0x85, 0xab, 0x43, 0x81, 0xc0, 0xb1, 0x40, 0xe0, 0xab, 0x40, 0xe0, 0xad, 0x44,
	0xce, 0xb1, 0x44, 0xce, 0x47, 0x89, 0x9c, 0xc7, 0x9b, 0x8c, 0xdb, 0xed, 0x73, 0x82, 0x53, 0xb9,
	0x27, 0x5d, 0x51, 0xc4, 0x28, 0x41, 0x5e, 0x49, 0xbb, 0x68, 0x9b, 0x2b, 0x66, 0x92, 0xff, 0xf5,
	0xd6, 0xee, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x05, 0x1a, 0xbe, 0x07, 0x00, 0x02, 0x00, 0x00,
}

func (m *Coordinator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coordinator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Coordinator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoordinator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.CoordinatorID != 0 {
		i = encodeVarintCoordinator(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CoordinatorDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoordinatorDescription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoordinatorDescription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CoordinatorByAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoordinatorByAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoordinatorByAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		i = encodeVarintCoordinator(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoordinator(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoordinator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Coordinator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		n += 1 + sovCoordinator(uint64(m.CoordinatorID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovCoordinator(uint64(l))
	return n
}

func (m *CoordinatorDescription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	return n
}

func (m *CoordinatorByAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	if m.CoordinatorID != 0 {
		n += 1 + sovCoordinator(uint64(m.CoordinatorID))
	}
	return n
}

func sovCoordinator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoordinator(x uint64) (n int) {
	return sovCoordinator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Coordinator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: Coordinator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coordinator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *CoordinatorDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: CoordinatorDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoordinatorDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *CoordinatorByAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: CoordinatorByAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoordinatorByAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func skipCoordinator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoordinator
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
					return 0, ErrIntOverflowCoordinator
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
					return 0, ErrIntOverflowCoordinator
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
				return 0, ErrInvalidLengthCoordinator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoordinator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoordinator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoordinator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoordinator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoordinator = fmt.Errorf("proto: unexpected end of group")
)
