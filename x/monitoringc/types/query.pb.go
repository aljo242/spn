// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb4a38bab58f58d, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb4a38bab58f58d, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryVerifiedClientIdsRequest struct {
	LaunchID   uint64             `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryVerifiedClientIdsRequest) Reset()         { *m = QueryVerifiedClientIdsRequest{} }
func (m *QueryVerifiedClientIdsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVerifiedClientIdsRequest) ProtoMessage()    {}
func (*QueryVerifiedClientIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb4a38bab58f58d, []int{2}
}
func (m *QueryVerifiedClientIdsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVerifiedClientIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVerifiedClientIdsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVerifiedClientIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVerifiedClientIdsRequest.Merge(m, src)
}
func (m *QueryVerifiedClientIdsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVerifiedClientIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVerifiedClientIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVerifiedClientIdsRequest proto.InternalMessageInfo

func (m *QueryVerifiedClientIdsRequest) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *QueryVerifiedClientIdsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryVerifiedClientIdsResponse struct {
	ClientIds  []string            `protobuf:"bytes,1,rep,name=clientIds,proto3" json:"clientIds,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryVerifiedClientIdsResponse) Reset()         { *m = QueryVerifiedClientIdsResponse{} }
func (m *QueryVerifiedClientIdsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVerifiedClientIdsResponse) ProtoMessage()    {}
func (*QueryVerifiedClientIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb4a38bab58f58d, []int{3}
}
func (m *QueryVerifiedClientIdsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVerifiedClientIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVerifiedClientIdsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVerifiedClientIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVerifiedClientIdsResponse.Merge(m, src)
}
func (m *QueryVerifiedClientIdsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVerifiedClientIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVerifiedClientIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVerifiedClientIdsResponse proto.InternalMessageInfo

func (m *QueryVerifiedClientIdsResponse) GetClientIds() []string {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *QueryVerifiedClientIdsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "tendermint.spn.monitoringc.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "tendermint.spn.monitoringc.QueryParamsResponse")
	proto.RegisterType((*QueryVerifiedClientIdsRequest)(nil), "tendermint.spn.monitoringc.QueryVerifiedClientIdsRequest")
	proto.RegisterType((*QueryVerifiedClientIdsResponse)(nil), "tendermint.spn.monitoringc.QueryVerifiedClientIdsResponse")
}

func init() { proto.RegisterFile("monitoringc/query.proto", fileDescriptor_ecb4a38bab58f58d) }

var fileDescriptor_ecb4a38bab58f58d = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x37, 0x6b, 0x5d, 0x6c, 0x3c, 0x19, 0x0b, 0x0e, 0x43, 0x1d, 0xcb, 0x20, 0x5a, 0x0a,
	0x26, 0xb4, 0x9e, 0x2c, 0x1e, 0x64, 0x15, 0x4b, 0x6f, 0x75, 0x0e, 0x0a, 0x5e, 0x4a, 0x66, 0x26,
	0xa6, 0x81, 0x9d, 0x24, 0x9d, 0x64, 0x8a, 0x45, 0xbc, 0xe8, 0xc1, 0xab, 0x20, 0x7e, 0xa7, 0x9e,
	0xa4, 0xe0, 0xc5, 0x93, 0xc8, 0xae, 0x1f, 0xc0, 0x8f, 0x20, 0x9b, 0x64, 0xdd, 0xa9, 0xeb, 0x76,
	0xa5, 0xb7, 0x99, 0x37, 0xef, 0xff, 0xfe, 0xbf, 0xfc, 0x5f, 0x06, 0xde, 0xa8, 0x94, 0x14, 0x56,
	0xd5, 0x42, 0xf2, 0x82, 0x1c, 0x36, 0xac, 0x3e, 0xc6, 0xba, 0x56, 0x56, 0xa1, 0xd8, 0x32, 0x59,
	0xb2, 0xba, 0x12, 0xd2, 0x62, 0xa3, 0x25, 0x6e, 0xf5, 0xc5, 0x2b, 0x5c, 0x71, 0xe5, 0xda, 0xc8,
	0xf8, 0xc9, 0x2b, 0xe2, 0x55, 0xae, 0x14, 0x1f, 0x30, 0x42, 0xb5, 0x20, 0x54, 0x4a, 0x65, 0xa9,
	0x15, 0x4a, 0x9a, 0xf0, 0x75, 0xa3, 0x50, 0xa6, 0x52, 0x86, 0xe4, 0xd4, 0x30, 0x6f, 0x44, 0x8e,
	0x36, 0x73, 0x66, 0xe9, 0x26, 0xd1, 0x94, 0x0b, 0xe9, 0x9a, 0x43, 0x6f, 0xd4, 0x86, 0xd2, 0xb4,
	0xa6, 0x55, 0x98, 0x92, 0xae, 0x40, 0xf4, 0x6c, 0xac, 0xdd, 0x73, 0xc5, 0x8c, 0x1d, 0x36, 0xcc,
	0xd8, 0xf4, 0x05, 0xbc, 0x7e, 0xa6, 0x6a, 0xb4, 0x92, 0x86, 0xa1, 0x47, 0xb0, 0xe7, 0xc5, 0x11,
	0x58, 0x03, 0xeb, 0x57, 0xb7, 0x52, 0x3c, 0xff, 0x4c, 0xd8, 0x6b, 0xfb, 0x4b, 0x27, 0xdf, 0x6f,
	0x75, 0xb2, 0xa0, 0x4b, 0xdf, 0x03, 0x78, 0xd3, 0x4d, 0x7e, 0xce, 0x6a, 0xf1, 0x4a, 0xb0, 0xf2,
	0xf1, 0x40, 0x30, 0x69, 0x77, 0xcb, 0x89, 0x35, 0x8a, 0xe1, 0x95, 0x01, 0x6d, 0x64, 0x71, 0xb0,
	0xfb, 0xc4, 0xb9, 0x2c, 0x65, 0x7f, 0xde, 0xd1, 0x53, 0x08, 0xa7, 0x47, 0x8b, 0xba, 0x8e, 0xe1,
	0x0e, 0xf6, 0x39, 0xe0, 0x71, 0x0e, 0xd8, 0x07, 0x1e, 0x72, 0xc0, 0x7b, 0x94, 0xb3, 0x30, 0x37,
	0x6b, 0x29, 0xd3, 0x0f, 0x00, 0x26, 0xf3, 0x28, 0xc2, 0x51, 0x57, 0xe1, 0x72, 0x31, 0x29, 0x46,
	0x60, 0xed, 0xd2, 0xfa, 0x72, 0x36, 0x2d, 0xa0, 0x9d, 0x7f, 0x80, 0xdc, 0x5d, 0x08, 0xe2, 0x47,
	0xb7, 0x49, 0xb6, 0x7e, 0x75, 0xe1, 0x65, 0x47, 0x82, 0x3e, 0x03, 0xd8, 0xf3, 0x91, 0x21, 0x7c,
	0x5e, 0xac, 0xb3, 0xdb, 0x8a, 0xc9, 0x7f, 0xf7, 0x7b, 0x82, 0x74, 0xe3, 0xdd, 0xd7, 0x9f, 0x9f,
	0xba, 0xb7, 0x51, 0x4a, 0xa6, 0x42, 0x62, 0xb4, 0x24, 0xb3, 0xd7, 0x04, 0x7d, 0x01, 0xf0, 0xda,
	0x4c, 0x4c, 0xe8, 0xc1, 0x42, 0xcb, 0x79, 0x0b, 0x8e, 0xb7, 0x2f, 0x22, 0x0d, 0xe0, 0x7d, 0x07,
	0xfe, 0x10, 0x6d, 0x9f, 0x07, 0x7e, 0x14, 0xe4, 0xfb, 0x7e, 0x5f, 0xfb, 0xa2, 0x34, 0xe4, 0xcd,
	0xe4, 0x0e, 0xbd, 0xed, 0xef, 0x9c, 0x0c, 0x13, 0x70, 0x3a, 0x4c, 0xc0, 0x8f, 0x61, 0x02, 0x3e,
	0x8e, 0x92, 0xce, 0xe9, 0x28, 0xe9, 0x7c, 0x1b, 0x25, 0x9d, 0x97, 0xf7, 0xb8, 0xb0, 0x07, 0x4d,
	0x8e, 0x0b, 0x55, 0xfd, 0x3d, 0xff, 0xf5, 0x19, 0x07, 0x7b, 0xac, 0x99, 0xc9, 0x7b, 0xee, 0x0f,
	0xba, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x7c, 0x4f, 0x14, 0xf2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of VerifiedClientIds items.
	VerifiedClientIds(ctx context.Context, in *QueryVerifiedClientIdsRequest, opts ...grpc.CallOption) (*QueryVerifiedClientIdsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.monitoringc.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VerifiedClientIds(ctx context.Context, in *QueryVerifiedClientIdsRequest, opts ...grpc.CallOption) (*QueryVerifiedClientIdsResponse, error) {
	out := new(QueryVerifiedClientIdsResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.monitoringc.Query/VerifiedClientIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of VerifiedClientIds items.
	VerifiedClientIds(context.Context, *QueryVerifiedClientIdsRequest) (*QueryVerifiedClientIdsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) VerifiedClientIds(ctx context.Context, req *QueryVerifiedClientIdsRequest) (*QueryVerifiedClientIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifiedClientIds not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.monitoringc.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VerifiedClientIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVerifiedClientIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VerifiedClientIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.monitoringc.Query/VerifiedClientIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VerifiedClientIds(ctx, req.(*QueryVerifiedClientIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.monitoringc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "VerifiedClientIds",
			Handler:    _Query_VerifiedClientIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoringc/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryVerifiedClientIdsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVerifiedClientIdsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVerifiedClientIdsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVerifiedClientIdsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVerifiedClientIdsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVerifiedClientIdsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientIds) > 0 {
		for iNdEx := len(m.ClientIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClientIds[iNdEx])
			copy(dAtA[i:], m.ClientIds[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.ClientIds[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryVerifiedClientIdsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovQuery(uint64(m.LaunchID))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVerifiedClientIdsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClientIds) > 0 {
		for _, s := range m.ClientIds {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryVerifiedClientIdsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryVerifiedClientIdsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVerifiedClientIdsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryVerifiedClientIdsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryVerifiedClientIdsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVerifiedClientIdsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIds = append(m.ClientIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)