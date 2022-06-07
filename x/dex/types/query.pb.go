// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/query.proto

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
	return fileDescriptor_d8e98105e6e08a59, []int{0}
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
	return fileDescriptor_d8e98105e6e08a59, []int{1}
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

type QueryGetSellOrderBookRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetSellOrderBookRequest) Reset()         { *m = QueryGetSellOrderBookRequest{} }
func (m *QueryGetSellOrderBookRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSellOrderBookRequest) ProtoMessage()    {}
func (*QueryGetSellOrderBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{2}
}
func (m *QueryGetSellOrderBookRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSellOrderBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSellOrderBookRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSellOrderBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSellOrderBookRequest.Merge(m, src)
}
func (m *QueryGetSellOrderBookRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSellOrderBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSellOrderBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSellOrderBookRequest proto.InternalMessageInfo

func (m *QueryGetSellOrderBookRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetSellOrderBookResponse struct {
	SellOrderBook SellOrderBook `protobuf:"bytes,1,opt,name=sellOrderBook,proto3" json:"sellOrderBook"`
}

func (m *QueryGetSellOrderBookResponse) Reset()         { *m = QueryGetSellOrderBookResponse{} }
func (m *QueryGetSellOrderBookResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetSellOrderBookResponse) ProtoMessage()    {}
func (*QueryGetSellOrderBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{3}
}
func (m *QueryGetSellOrderBookResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSellOrderBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSellOrderBookResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSellOrderBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSellOrderBookResponse.Merge(m, src)
}
func (m *QueryGetSellOrderBookResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSellOrderBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSellOrderBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSellOrderBookResponse proto.InternalMessageInfo

func (m *QueryGetSellOrderBookResponse) GetSellOrderBook() SellOrderBook {
	if m != nil {
		return m.SellOrderBook
	}
	return SellOrderBook{}
}

type QueryAllSellOrderBookRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSellOrderBookRequest) Reset()         { *m = QueryAllSellOrderBookRequest{} }
func (m *QueryAllSellOrderBookRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllSellOrderBookRequest) ProtoMessage()    {}
func (*QueryAllSellOrderBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{4}
}
func (m *QueryAllSellOrderBookRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSellOrderBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSellOrderBookRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSellOrderBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSellOrderBookRequest.Merge(m, src)
}
func (m *QueryAllSellOrderBookRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSellOrderBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSellOrderBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSellOrderBookRequest proto.InternalMessageInfo

func (m *QueryAllSellOrderBookRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllSellOrderBookResponse struct {
	SellOrderBook []SellOrderBook     `protobuf:"bytes,1,rep,name=sellOrderBook,proto3" json:"sellOrderBook"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSellOrderBookResponse) Reset()         { *m = QueryAllSellOrderBookResponse{} }
func (m *QueryAllSellOrderBookResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllSellOrderBookResponse) ProtoMessage()    {}
func (*QueryAllSellOrderBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{5}
}
func (m *QueryAllSellOrderBookResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSellOrderBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSellOrderBookResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSellOrderBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSellOrderBookResponse.Merge(m, src)
}
func (m *QueryAllSellOrderBookResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSellOrderBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSellOrderBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSellOrderBookResponse proto.InternalMessageInfo

func (m *QueryAllSellOrderBookResponse) GetSellOrderBook() []SellOrderBook {
	if m != nil {
		return m.SellOrderBook
	}
	return nil
}

func (m *QueryAllSellOrderBookResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "elisionio.elision.dex.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "elisionio.elision.dex.QueryParamsResponse")
	proto.RegisterType((*QueryGetSellOrderBookRequest)(nil), "elisionio.elision.dex.QueryGetSellOrderBookRequest")
	proto.RegisterType((*QueryGetSellOrderBookResponse)(nil), "elisionio.elision.dex.QueryGetSellOrderBookResponse")
	proto.RegisterType((*QueryAllSellOrderBookRequest)(nil), "elisionio.elision.dex.QueryAllSellOrderBookRequest")
	proto.RegisterType((*QueryAllSellOrderBookResponse)(nil), "elisionio.elision.dex.QueryAllSellOrderBookResponse")
}

func init() { proto.RegisterFile("dex/query.proto", fileDescriptor_d8e98105e6e08a59) }

var fileDescriptor_d8e98105e6e08a59 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0xad, 0x0d, 0x38, 0x52, 0x2c, 0x63, 0x04, 0x0d, 0xed, 0x5a, 0x16, 0x69, 0x4d,
	0xc1, 0x19, 0xd2, 0x16, 0x3c, 0x78, 0x6a, 0x04, 0x7b, 0x34, 0xae, 0x37, 0x2f, 0x65, 0xb6, 0x79,
	0xae, 0x43, 0x27, 0xfb, 0x36, 0x3b, 0x1b, 0x49, 0x11, 0x2f, 0x5e, 0xbc, 0x0a, 0x7e, 0x0f, 0x2f,
	0xde, 0xfc, 0x04, 0x3d, 0x16, 0xbc, 0x78, 0x12, 0x49, 0xfc, 0x10, 0x1e, 0x65, 0x67, 0xa6, 0x34,
	0x6b, 0xb3, 0xb1, 0xd2, 0xdb, 0x64, 0xe7, 0xfd, 0xff, 0xef, 0xf7, 0xf6, 0xfd, 0xb3, 0xe4, 0x66,
	0x0f, 0x46, 0x7c, 0x30, 0x84, 0xec, 0x98, 0xa5, 0x19, 0xe6, 0x48, 0x6f, 0x83, 0x92, 0x5a, 0x62,
	0x22, 0x91, 0xb9, 0x13, 0xeb, 0xc1, 0xa8, 0xd9, 0x88, 0x31, 0x46, 0x53, 0xc1, 0x8b, 0x93, 0x2d,
	0x6e, 0xae, 0xc6, 0x88, 0xb1, 0x02, 0x2e, 0x52, 0xc9, 0x45, 0x92, 0x60, 0x2e, 0x72, 0x89, 0x89,
	0x76, 0xb7, 0x5b, 0x87, 0xa8, 0xfb, 0xa8, 0x79, 0x24, 0x34, 0xd8, 0x1e, 0xfc, 0x4d, 0x3b, 0x82,
	0x5c, 0xb4, 0x79, 0x2a, 0x62, 0x99, 0x98, 0x62, 0x57, 0xbb, 0x52, 0x70, 0xa4, 0x22, 0x13, 0xfd,
	0x33, 0xf5, 0xdd, 0xe2, 0x89, 0x06, 0xa5, 0x0e, 0x30, 0xeb, 0x41, 0x76, 0x10, 0x21, 0x1e, 0xd9,
	0xab, 0xa0, 0x41, 0xe8, 0xf3, 0xc2, 0xae, 0x6b, 0xea, 0x43, 0x18, 0x0c, 0x41, 0xe7, 0x41, 0x48,
	0x6e, 0x95, 0x9e, 0xea, 0x14, 0x13, 0x0d, 0xf4, 0x31, 0xa9, 0x5b, 0xdf, 0x3b, 0xde, 0xba, 0xf7,
	0xe0, 0xc6, 0xf6, 0x1a, 0x9b, 0x39, 0x21, 0xb3, 0xb2, 0xce, 0xb5, 0x93, 0x1f, 0xf7, 0x6a, 0xa1,
	0x93, 0x04, 0xbb, 0x64, 0xd5, 0x78, 0xee, 0x43, 0xfe, 0x02, 0x94, 0x7a, 0x56, 0x90, 0x74, 0x10,
	0x8f, 0x5c, 0x4f, 0xda, 0x20, 0x4b, 0x32, 0xe9, 0xc1, 0xc8, 0x78, 0x5f, 0x0f, 0xed, 0x8f, 0x60,
	0x40, 0xd6, 0x2a, 0x54, 0x8e, 0xa9, 0x4b, 0x96, 0xf5, 0xf4, 0x85, 0x43, 0xbb, 0x5f, 0x81, 0x56,
	0x32, 0x71, 0x84, 0x65, 0x83, 0xe0, 0x95, 0x03, 0xdd, 0x53, 0x6a, 0x26, 0xe8, 0x53, 0x42, 0xce,
	0xdf, 0xb9, 0x6b, 0xb7, 0xc1, 0xec, 0x82, 0x58, 0xb1, 0x20, 0x66, 0x43, 0xe0, 0x16, 0xc4, 0xba,
	0x22, 0x06, 0xa7, 0x0d, 0xa7, 0x94, 0xc1, 0x57, 0xcf, 0xcd, 0x76, 0xb1, 0x51, 0xf5, 0x6c, 0x8b,
	0x57, 0x9a, 0x8d, 0xee, 0x97, 0xd8, 0x17, 0x0c, 0xfb, 0xe6, 0x3f, 0xd9, 0x2d, 0xce, 0x34, 0xfc,
	0xf6, 0xef, 0x45, 0xb2, 0x64, 0xe0, 0xe9, 0x07, 0x8f, 0xd4, 0xed, 0xc2, 0x69, 0xab, 0x02, 0xec,
	0x62, 0xc2, 0x9a, 0x5b, 0x97, 0x29, 0xb5, 0x7d, 0x83, 0x8d, 0xf7, 0xdf, 0x7e, 0x7d, 0x5a, 0x58,
	0xa7, 0x3e, 0x77, 0x95, 0x0f, 0x25, 0x9e, 0x1d, 0xf9, 0x79, 0xd8, 0xe9, 0x17, 0x8f, 0x2c, 0x97,
	0xde, 0x01, 0xdd, 0x99, 0xd7, 0xa5, 0x22, 0x88, 0xcd, 0xdd, 0xff, 0x13, 0x39, 0xc8, 0x47, 0x06,
	0xb2, 0x4d, 0x79, 0x15, 0xe4, 0x5f, 0xff, 0x3f, 0xfe, 0xd6, 0x04, 0xfc, 0x1d, 0xfd, 0xec, 0x91,
	0x95, 0x92, 0xe5, 0x9e, 0x52, 0xf3, 0xc1, 0x2b, 0x82, 0x39, 0x1f, 0xbc, 0x2a, 0x64, 0x01, 0x37,
	0xe0, 0x2d, 0xba, 0x79, 0x49, 0xf0, 0xce, 0x93, 0x93, 0xb1, 0xef, 0x9d, 0x8e, 0x7d, 0xef, 0xe7,
	0xd8, 0xf7, 0x3e, 0x4e, 0xfc, 0xda, 0xe9, 0xc4, 0xaf, 0x7d, 0x9f, 0xf8, 0xb5, 0x97, 0xad, 0x58,
	0xe6, 0xaf, 0x87, 0x11, 0x3b, 0xc4, 0xfe, 0x2c, 0xb3, 0x91, 0xb1, 0xcb, 0x8f, 0x53, 0xd0, 0x51,
	0xdd, 0x7c, 0x7e, 0x76, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xca, 0x01, 0xef, 0x35, 0x05,
	0x00, 0x00,
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
	// Queries a SellOrderBook by index.
	SellOrderBook(ctx context.Context, in *QueryGetSellOrderBookRequest, opts ...grpc.CallOption) (*QueryGetSellOrderBookResponse, error)
	// Queries a list of SellOrderBook items.
	SellOrderBookAll(ctx context.Context, in *QueryAllSellOrderBookRequest, opts ...grpc.CallOption) (*QueryAllSellOrderBookResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/elisionio.elision.dex.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrderBook(ctx context.Context, in *QueryGetSellOrderBookRequest, opts ...grpc.CallOption) (*QueryGetSellOrderBookResponse, error) {
	out := new(QueryGetSellOrderBookResponse)
	err := c.cc.Invoke(ctx, "/elisionio.elision.dex.Query/SellOrderBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrderBookAll(ctx context.Context, in *QueryAllSellOrderBookRequest, opts ...grpc.CallOption) (*QueryAllSellOrderBookResponse, error) {
	out := new(QueryAllSellOrderBookResponse)
	err := c.cc.Invoke(ctx, "/elisionio.elision.dex.Query/SellOrderBookAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a SellOrderBook by index.
	SellOrderBook(context.Context, *QueryGetSellOrderBookRequest) (*QueryGetSellOrderBookResponse, error)
	// Queries a list of SellOrderBook items.
	SellOrderBookAll(context.Context, *QueryAllSellOrderBookRequest) (*QueryAllSellOrderBookResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) SellOrderBook(ctx context.Context, req *QueryGetSellOrderBookRequest) (*QueryGetSellOrderBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrderBook not implemented")
}
func (*UnimplementedQueryServer) SellOrderBookAll(ctx context.Context, req *QueryAllSellOrderBookRequest) (*QueryAllSellOrderBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrderBookAll not implemented")
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
		FullMethod: "/elisionio.elision.dex.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSellOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elisionio.elision.dex.Query/SellOrderBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrderBook(ctx, req.(*QueryGetSellOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrderBookAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSellOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrderBookAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elisionio.elision.dex.Query/SellOrderBookAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrderBookAll(ctx, req.(*QueryAllSellOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elisionio.elision.dex.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SellOrderBook",
			Handler:    _Query_SellOrderBook_Handler,
		},
		{
			MethodName: "SellOrderBookAll",
			Handler:    _Query_SellOrderBookAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dex/query.proto",
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

func (m *QueryGetSellOrderBookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSellOrderBookRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSellOrderBookRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetSellOrderBookResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSellOrderBookResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSellOrderBookResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SellOrderBook.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllSellOrderBookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSellOrderBookRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSellOrderBookRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllSellOrderBookResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSellOrderBookResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSellOrderBookResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.SellOrderBook) > 0 {
		for iNdEx := len(m.SellOrderBook) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellOrderBook[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
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

func (m *QueryGetSellOrderBookRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetSellOrderBookResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SellOrderBook.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllSellOrderBookRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllSellOrderBookResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SellOrderBook) > 0 {
		for _, e := range m.SellOrderBook {
			l = e.Size()
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
func (m *QueryGetSellOrderBookRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetSellOrderBookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSellOrderBookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetSellOrderBookResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetSellOrderBookResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSellOrderBookResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderBook", wireType)
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
			if err := m.SellOrderBook.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllSellOrderBookRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllSellOrderBookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSellOrderBookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryAllSellOrderBookResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllSellOrderBookResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSellOrderBookResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderBook", wireType)
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
			m.SellOrderBook = append(m.SellOrderBook, SellOrderBook{})
			if err := m.SellOrderBook[len(m.SellOrderBook)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
