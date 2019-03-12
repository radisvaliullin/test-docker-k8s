// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/v1/tsrv/test.proto

package tsrv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_95a0419751309b28, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Response struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_95a0419751309b28, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "testsrv.v1.Request")
	proto.RegisterType((*Response)(nil), "testsrv.v1.Response")
}

func init() { proto.RegisterFile("pb/v1/tsrv/test.proto", fileDescriptor_95a0419751309b28) }

var fileDescriptor_95a0419751309b28 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0xcd, 0x14, 0xb4, 0xa6, 0x22, 0x25, 0x5a, 0x29, 0xd5, 0x45, 0x99, 0x55, 0x51, 0x3b,
	0x69, 0x75, 0x23, 0x8a, 0x22, 0xd5, 0x85, 0xeb, 0xa9, 0x2b, 0x77, 0x69, 0xe7, 0x31, 0x86, 0x4e,
	0x27, 0x93, 0x97, 0x4c, 0xc0, 0xad, 0xbf, 0xe0, 0xa7, 0xf9, 0x0b, 0x6e, 0xfc, 0x0b, 0xc9, 0xb4,
	0xa2, 0xae, 0xa4, 0xdd, 0x84, 0xc7, 0xe3, 0xde, 0x73, 0x2f, 0x49, 0x68, 0xab, 0x98, 0x70, 0x37,
	0xe4, 0xd6, 0xa0, 0xe3, 0x16, 0x8c, 0x8d, 0x0a, 0x54, 0x56, 0x31, 0xea, 0x67, 0x83, 0x2e, 0x72,
	0xc3, 0xce, 0x51, 0xaa, 0x54, 0x9a, 0x01, 0x17, 0x85, 0xe4, 0x22, 0xcf, 0x95, 0x15, 0x56, 0xaa,
	0xdc, 0x2c, 0x94, 0xe1, 0x09, 0xdd, 0x8a, 0x41, 0x97, 0x60, 0x2c, 0xdb, 0xa5, 0x81, 0x4c, 0xda,
	0xa4, 0x4b, 0x7a, 0xb5, 0x38, 0x90, 0x09, 0x6b, 0xd2, 0xda, 0xdc, 0xa4, 0xed, 0xa0, 0x4b, 0x7a,
	0xdb, 0xb1, 0x1f, 0xc3, 0x53, 0x5a, 0x8f, 0xc1, 0x14, 0x2a, 0x37, 0xf0, 0xbf, 0xfa, 0xec, 0x33,
	0xa0, 0x8d, 0x47, 0x30, 0x76, 0x0c, 0xe8, 0xe4, 0x14, 0xd8, 0x03, 0xad, 0x21, 0x68, 0xb6, 0x17,
	0xfd, 0x94, 0x8b, 0x96, 0xd9, 0x9d, 0xfd, 0xbf, 0xcb, 0x45, 0x46, 0x78, 0xf0, 0xfa, 0xfe, 0xf1,
	0x16, 0x34, 0xc3, 0x46, 0x55, 0xde, 0x0d, 0x39, 0x82, 0xbe, 0x24, 0xc7, 0x2c, 0xa6, 0x75, 0x04,
	0x7d, 0x0f, 0x99, 0x78, 0x59, 0x05, 0x77, 0x58, 0xe1, 0x5a, 0x61, 0xf3, 0x17, 0x2e, 0xf1, 0x10,
	0xcf, 0xbc, 0xa6, 0x3b, 0x08, 0x7a, 0x8c, 0x6e, 0x6c, 0x11, 0xc4, 0x7c, 0x15, 0xee, 0xc6, 0x80,
	0x2c, 0xed, 0x77, 0x59, 0xbe, 0x86, 0xbd, 0x47, 0xd8, 0x0d, 0x6d, 0x20, 0xe8, 0x91, 0x5c, 0xcb,
	0x3d, 0x20, 0xa3, 0xd1, 0xd3, 0x6d, 0x2a, 0xed, 0x73, 0x39, 0x89, 0xa6, 0x6a, 0xce, 0x51, 0x24,
	0xd2, 0x38, 0x91, 0xc9, 0x32, 0xcb, 0x64, 0x5e, 0x7d, 0x8c, 0x7e, 0xa2, 0xa6, 0x33, 0xc0, 0xfe,
	0xec, 0xc2, 0xf0, 0xea, 0xe5, 0xbf, 0xef, 0xc0, 0x03, 0xaf, 0xfc, 0x31, 0xd9, 0xac, 0xf6, 0xe7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x37, 0xda, 0x52, 0x54, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	Req(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ReqDelay(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ReqSrvStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ReqSrvStreamClient, error)
	ReqClnStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ReqClnStreamClient, error)
	ReqBiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ReqBiStreamClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Req(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/testsrv.v1.TestService/req", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ReqDelay(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/testsrv.v1.TestService/reqDelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ReqSrvStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ReqSrvStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/testsrv.v1.TestService/reqSrvStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceReqSrvStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ReqSrvStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceReqSrvStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceReqSrvStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ReqClnStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ReqClnStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/testsrv.v1.TestService/reqClnStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceReqClnStreamClient{stream}
	return x, nil
}

type TestService_ReqClnStreamClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type testServiceReqClnStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceReqClnStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceReqClnStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ReqBiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ReqBiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/testsrv.v1.TestService/reqBiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceReqBiStreamClient{stream}
	return x, nil
}

type TestService_ReqBiStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceReqBiStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceReqBiStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceReqBiStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	Req(context.Context, *Request) (*Response, error)
	ReqDelay(context.Context, *Request) (*Response, error)
	ReqSrvStream(*Request, TestService_ReqSrvStreamServer) error
	ReqClnStream(TestService_ReqClnStreamServer) error
	ReqBiStream(TestService_ReqBiStreamServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Req_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Req(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testsrv.v1.TestService/Req",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Req(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ReqDelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ReqDelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testsrv.v1.TestService/ReqDelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ReqDelay(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ReqSrvStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ReqSrvStream(m, &testServiceReqSrvStreamServer{stream})
}

type TestService_ReqSrvStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type testServiceReqSrvStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceReqSrvStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_ReqClnStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).ReqClnStream(&testServiceReqClnStreamServer{stream})
}

type TestService_ReqClnStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type testServiceReqClnStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceReqClnStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceReqClnStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_ReqBiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).ReqBiStream(&testServiceReqBiStreamServer{stream})
}

type TestService_ReqBiStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type testServiceReqBiStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceReqBiStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceReqBiStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testsrv.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req",
			Handler:    _TestService_Req_Handler,
		},
		{
			MethodName: "reqDelay",
			Handler:    _TestService_ReqDelay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "reqSrvStream",
			Handler:       _TestService_ReqSrvStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "reqClnStream",
			Handler:       _TestService_ReqClnStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "reqBiStream",
			Handler:       _TestService_ReqBiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/v1/tsrv/test.proto",
}
