// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: api/v1/log.proto

package logrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoggingServiceClient is the client API for LoggingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggingServiceClient interface {
	Log(ctx context.Context, in *LogRecord, opts ...grpc.CallOption) (*LogResponse, error)
	FetchLogs(ctx context.Context, in *FetchLogRequest, opts ...grpc.CallOption) (*FetchLogResponse, error)
}

type loggingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggingServiceClient(cc grpc.ClientConnInterface) LoggingServiceClient {
	return &loggingServiceClient{cc}
}

func (c *loggingServiceClient) Log(ctx context.Context, in *LogRecord, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/logrpc.LoggingService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceClient) FetchLogs(ctx context.Context, in *FetchLogRequest, opts ...grpc.CallOption) (*FetchLogResponse, error) {
	out := new(FetchLogResponse)
	err := c.cc.Invoke(ctx, "/logrpc.LoggingService/FetchLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServiceServer is the server API for LoggingService service.
// All implementations must embed UnimplementedLoggingServiceServer
// for forward compatibility
type LoggingServiceServer interface {
	Log(context.Context, *LogRecord) (*LogResponse, error)
	FetchLogs(context.Context, *FetchLogRequest) (*FetchLogResponse, error)
	mustEmbedUnimplementedLoggingServiceServer()
}

// UnimplementedLoggingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoggingServiceServer struct {
}

func (UnimplementedLoggingServiceServer) Log(context.Context, *LogRecord) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedLoggingServiceServer) FetchLogs(context.Context, *FetchLogRequest) (*FetchLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLogs not implemented")
}
func (UnimplementedLoggingServiceServer) mustEmbedUnimplementedLoggingServiceServer() {}

// UnsafeLoggingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggingServiceServer will
// result in compilation errors.
type UnsafeLoggingServiceServer interface {
	mustEmbedUnimplementedLoggingServiceServer()
}

func RegisterLoggingServiceServer(s grpc.ServiceRegistrar, srv LoggingServiceServer) {
	s.RegisterService(&LoggingService_ServiceDesc, srv)
}

func _LoggingService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logrpc.LoggingService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).Log(ctx, req.(*LogRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingService_FetchLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).FetchLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logrpc.LoggingService/FetchLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).FetchLogs(ctx, req.(*FetchLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoggingService_ServiceDesc is the grpc.ServiceDesc for LoggingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logrpc.LoggingService",
	HandlerType: (*LoggingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LoggingService_Log_Handler,
		},
		{
			MethodName: "FetchLogs",
			Handler:    _LoggingService_FetchLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/log.proto",
}
