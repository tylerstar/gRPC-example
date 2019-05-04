// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Numbers struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	LastNumber           int32    `protobuf:"varint,2,opt,name=last_number,json=lastNumber,proto3" json:"last_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()    {}
func (*Numbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *Numbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numbers.Unmarshal(m, b)
}
func (m *Numbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numbers.Marshal(b, m, deterministic)
}
func (m *Numbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numbers.Merge(m, src)
}
func (m *Numbers) XXX_Size() int {
	return xxx_messageInfo_Numbers.Size(m)
}
func (m *Numbers) XXX_DiscardUnknown() {
	xxx_messageInfo_Numbers.DiscardUnknown(m)
}

var xxx_messageInfo_Numbers proto.InternalMessageInfo

func (m *Numbers) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Numbers) GetLastNumber() int32 {
	if m != nil {
		return m.LastNumber
	}
	return 0
}

type SumRequest struct {
	Numbers              *Numbers `protobuf:"bytes,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetNumbers() *Numbers {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type NumberDecompositionRequest struct {
	Numbers              *Numbers `protobuf:"bytes,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberDecompositionRequest) Reset()         { *m = NumberDecompositionRequest{} }
func (m *NumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*NumberDecompositionRequest) ProtoMessage()    {}
func (*NumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *NumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberDecompositionRequest.Unmarshal(m, b)
}
func (m *NumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *NumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberDecompositionRequest.Merge(m, src)
}
func (m *NumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_NumberDecompositionRequest.Size(m)
}
func (m *NumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumberDecompositionRequest proto.InternalMessageInfo

func (m *NumberDecompositionRequest) GetNumbers() *Numbers {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type NumberDecompositionResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberDecompositionResponse) Reset()         { *m = NumberDecompositionResponse{} }
func (m *NumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*NumberDecompositionResponse) ProtoMessage()    {}
func (*NumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *NumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberDecompositionResponse.Unmarshal(m, b)
}
func (m *NumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *NumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberDecompositionResponse.Merge(m, src)
}
func (m *NumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_NumberDecompositionResponse.Size(m)
}
func (m *NumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NumberDecompositionResponse proto.InternalMessageInfo

func (m *NumberDecompositionResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "calculatorpb.Numbers")
	proto.RegisterType((*SumRequest)(nil), "calculatorpb.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculatorpb.SumResponse")
	proto.RegisterType((*NumberDecompositionRequest)(nil), "calculatorpb.NumberDecompositionRequest")
	proto.RegisterType((*NumberDecompositionResponse)(nil), "calculatorpb.NumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculatorpb.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculatorpb.ComputeAverageResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0xa5, 0xc6, 0x36, 0x19, 0x9a, 0x1a, 0x37, 0x69, 0x53, 0xf1, 0xa0, 0xae, 0x9a, 0xe0,
	0xa5, 0x6d, 0x6a, 0xbc, 0xe9, 0x41, 0xeb, 0xb5, 0x1e, 0xe0, 0xe6, 0xa5, 0x01, 0x32, 0x1a, 0x0c,
	0xb0, 0xb8, 0x7f, 0xfa, 0x38, 0x3e, 0xab, 0x71, 0x61, 0x05, 0xcc, 0x46, 0x8d, 0x37, 0x66, 0xf6,
	0x9b, 0xf9, 0xcd, 0x7c, 0x0c, 0xf8, 0x49, 0x94, 0x25, 0x2a, 0x8b, 0x24, 0xe3, 0xf3, 0xe6, 0xb3,
	0x8c, 0x5b, 0xc1, 0xac, 0xe4, 0x4c, 0x32, 0x32, 0x6c, 0x3f, 0xd3, 0x35, 0x0c, 0x1e, 0x55, 0x1e,
	0x23, 0x17, 0xe4, 0x14, 0x86, 0xcf, 0x29, 0x17, 0x72, 0x53, 0xe8, 0xc4, 0xd4, 0x39, 0x71, 0xfc,
	0xbd, 0xc0, 0xd5, 0xb9, 0x4a, 0x43, 0x8e, 0xc1, 0xcd, 0xa2, 0x46, 0xd1, 0xd3, 0x0a, 0xf8, 0x4c,
	0x55, 0x02, 0x7a, 0x0b, 0x10, 0xaa, 0x3c, 0xc0, 0x37, 0x85, 0x42, 0x92, 0x39, 0x0c, 0x2a, 0xa5,
	0xd0, 0xcd, 0xdc, 0xe5, 0x78, 0xd6, 0x86, 0xcf, 0x6a, 0x72, 0x60, 0x54, 0xf4, 0x02, 0x5c, 0x5d,
	0x2e, 0x4a, 0x56, 0x08, 0x24, 0x13, 0xe8, 0x73, 0x14, 0x2a, 0x93, 0xf5, 0x2c, 0x75, 0x44, 0xd7,
	0xe0, 0x55, 0xa5, 0x0f, 0x98, 0xb0, 0xbc, 0x64, 0x22, 0x95, 0x29, 0x2b, 0xfe, 0x4d, 0xbd, 0x86,
	0x23, 0x6b, 0xbb, 0x5f, 0xa6, 0x98, 0xc3, 0x78, 0xc5, 0xf2, 0x52, 0x49, 0xbc, 0xdb, 0x22, 0x8f,
	0x5e, 0xd0, 0x0c, 0x30, 0x81, 0x7e, 0xc7, 0xc2, 0x3a, 0xa2, 0x0b, 0x98, 0x7c, 0x2f, 0xb0, 0x22,
	0x1c, 0x83, 0x58, 0xbe, 0xf7, 0xe0, 0x60, 0xf5, 0x35, 0x7b, 0x88, 0x7c, 0x9b, 0x26, 0x48, 0x6e,
	0x60, 0x37, 0x54, 0x39, 0x99, 0x76, 0xd7, 0x6a, 0x7c, 0xf7, 0x0e, 0x2d, 0x2f, 0x15, 0x89, 0xee,
	0x90, 0x57, 0xd8, 0x37, 0x7b, 0x62, 0xfd, 0x5b, 0x7d, 0x9b, 0x41, 0x36, 0x6f, 0xbd, 0xcb, 0x3f,
	0x28, 0x0d, 0x69, 0xe1, 0x90, 0x0d, 0x8c, 0xba, 0x1b, 0x93, 0xb3, 0x6e, 0x03, 0xab, 0x81, 0xde,
	0xf9, 0xcf, 0x22, 0x03, 0xf0, 0x9d, 0xfb, 0xd1, 0x53, 0xe7, 0x9c, 0xe3, 0xbe, 0xbe, 0xf1, 0xab,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x16, 0xa9, 0x9c, 0x0f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server streaming
	DecomposeNumber(ctx context.Context, in *NumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_DecomposeNumberClient, error)
	// Client Streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculatorpb.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) DecomposeNumber(ctx context.Context, in *NumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_DecomposeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculatorpb.CalculatorService/DecomposeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceDecomposeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_DecomposeNumberClient interface {
	Recv() (*NumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServiceDecomposeNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceDecomposeNumberClient) Recv() (*NumberDecompositionResponse, error) {
	m := new(NumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculatorpb.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server streaming
	DecomposeNumber(*NumberDecompositionRequest, CalculatorService_DecomposeNumberServer) error
	// Client Streaming
	ComputeAverage(CalculatorService_ComputeAverageServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) DecomposeNumber(req *NumberDecompositionRequest, srv CalculatorService_DecomposeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method DecomposeNumber not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculatorpb.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_DecomposeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).DecomposeNumber(m, &calculatorServiceDecomposeNumberServer{stream})
}

type CalculatorService_DecomposeNumberServer interface {
	Send(*NumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServiceDecomposeNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceDecomposeNumberServer) Send(m *NumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculatorpb.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DecomposeNumber",
			Handler:       _CalculatorService_DecomposeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}