// calculator/calculator.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: calculator/calculator.proto

package calculator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CalculatorService_Add_FullMethodName        = "/calculator.CalculatorService/Add"
	CalculatorService_Multiply_FullMethodName   = "/calculator.CalculatorService/Multiply"
	CalculatorService_GetHistory_FullMethodName = "/calculator.CalculatorService/GetHistory"
)

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Add(ctx context.Context, in *TwoNumbers, opts ...grpc.CallOption) (*CalculationResult, error)
	Multiply(ctx context.Context, in *TwoNumbers, opts ...grpc.CallOption) (*CalculationResult, error)
	GetHistory(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HistoryEntry], error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *TwoNumbers, opts ...grpc.CallOption) (*CalculationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResult)
	err := c.cc.Invoke(ctx, CalculatorService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *TwoNumbers, opts ...grpc.CallOption) (*CalculationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResult)
	err := c.cc.Invoke(ctx, CalculatorService_Multiply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) GetHistory(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HistoryEntry], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], CalculatorService_GetHistory_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HistoryRequest, HistoryEntry]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetHistoryClient = grpc.ServerStreamingClient[HistoryEntry]

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility.
type CalculatorServiceServer interface {
	Add(context.Context, *TwoNumbers) (*CalculationResult, error)
	Multiply(context.Context, *TwoNumbers) (*CalculationResult, error)
	GetHistory(*HistoryRequest, grpc.ServerStreamingServer[HistoryEntry]) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServiceServer struct{}

func (UnimplementedCalculatorServiceServer) Add(context.Context, *TwoNumbers) (*CalculationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServiceServer) Multiply(context.Context, *TwoNumbers) (*CalculationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculatorServiceServer) GetHistory(*HistoryRequest, grpc.ServerStreamingServer[HistoryEntry]) error {
	return status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}
func (UnimplementedCalculatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*TwoNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*TwoNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_GetHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).GetHistory(m, &grpc.GenericServerStream[HistoryRequest, HistoryEntry]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetHistoryServer = grpc.ServerStreamingServer[HistoryEntry]

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetHistory",
			Handler:       _CalculatorService_GetHistory_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculator.proto",
}
