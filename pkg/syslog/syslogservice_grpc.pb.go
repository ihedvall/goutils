// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: syslogservice.proto

package syslog

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SyslogService_GetLastEvent_FullMethodName       = "/syslog.SyslogService/GetLastEvent"
	SyslogService_GetCount_FullMethodName           = "/syslog.SyslogService/GetCount"
	SyslogService_GetEvent_FullMethodName           = "/syslog.SyslogService/GetEvent"
	SyslogService_GetSyslog_FullMethodName          = "/syslog.SyslogService/GetSyslog"
	SyslogService_GetDataDefinitions_FullMethodName = "/syslog.SyslogService/GetDataDefinitions"
)

// SyslogServiceClient is the client API for SyslogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyslogServiceClient interface {
	GetLastEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventMessage, error)
	GetCount(ctx context.Context, in *SyslogFilter, opts ...grpc.CallOption) (*SyslogCount, error)
	GetEvent(ctx context.Context, in *SyslogFilter, opts ...grpc.CallOption) (SyslogService_GetEventClient, error)
	GetSyslog(ctx context.Context, in *SyslogFilter, opts ...grpc.CallOption) (SyslogService_GetSyslogClient, error)
	GetDataDefinitions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (SyslogService_GetDataDefinitionsClient, error)
}

type syslogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyslogServiceClient(cc grpc.ClientConnInterface) SyslogServiceClient {
	return &syslogServiceClient{cc}
}

func (c *syslogServiceClient) GetLastEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, SyslogService_GetLastEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogServiceClient) GetCount(ctx context.Context, in *SyslogFilter, opts ...grpc.CallOption) (*SyslogCount, error) {
	out := new(SyslogCount)
	err := c.cc.Invoke(ctx, SyslogService_GetCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogServiceClient) GetEvent(ctx context.Context, in *SyslogFilter, opts ...grpc.CallOption) (SyslogService_GetEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyslogService_ServiceDesc.Streams[0], SyslogService_GetEvent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &syslogServiceGetEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SyslogService_GetEventClient interface {
	Recv() (*EventMessage, error)
	grpc.ClientStream
}

type syslogServiceGetEventClient struct {
	grpc.ClientStream
}

func (x *syslogServiceGetEventClient) Recv() (*EventMessage, error) {
	m := new(EventMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syslogServiceClient) GetSyslog(ctx context.Context, in *SyslogFilter, opts ...grpc.CallOption) (SyslogService_GetSyslogClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyslogService_ServiceDesc.Streams[1], SyslogService_GetSyslog_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &syslogServiceGetSyslogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SyslogService_GetSyslogClient interface {
	Recv() (*SyslogMessage, error)
	grpc.ClientStream
}

type syslogServiceGetSyslogClient struct {
	grpc.ClientStream
}

func (x *syslogServiceGetSyslogClient) Recv() (*SyslogMessage, error) {
	m := new(SyslogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syslogServiceClient) GetDataDefinitions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (SyslogService_GetDataDefinitionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SyslogService_ServiceDesc.Streams[2], SyslogService_GetDataDefinitions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &syslogServiceGetDataDefinitionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SyslogService_GetDataDefinitionsClient interface {
	Recv() (*SyslogDataDefinition, error)
	grpc.ClientStream
}

type syslogServiceGetDataDefinitionsClient struct {
	grpc.ClientStream
}

func (x *syslogServiceGetDataDefinitionsClient) Recv() (*SyslogDataDefinition, error) {
	m := new(SyslogDataDefinition)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyslogServiceServer is the server API for SyslogService service.
// All implementations must embed UnimplementedSyslogServiceServer
// for forward compatibility
type SyslogServiceServer interface {
	GetLastEvent(context.Context, *emptypb.Empty) (*EventMessage, error)
	GetCount(context.Context, *SyslogFilter) (*SyslogCount, error)
	GetEvent(*SyslogFilter, SyslogService_GetEventServer) error
	GetSyslog(*SyslogFilter, SyslogService_GetSyslogServer) error
	GetDataDefinitions(*emptypb.Empty, SyslogService_GetDataDefinitionsServer) error
	mustEmbedUnimplementedSyslogServiceServer()
}

// UnimplementedSyslogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyslogServiceServer struct {
}

func (UnimplementedSyslogServiceServer) GetLastEvent(context.Context, *emptypb.Empty) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastEvent not implemented")
}
func (UnimplementedSyslogServiceServer) GetCount(context.Context, *SyslogFilter) (*SyslogCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}
func (UnimplementedSyslogServiceServer) GetEvent(*SyslogFilter, SyslogService_GetEventServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedSyslogServiceServer) GetSyslog(*SyslogFilter, SyslogService_GetSyslogServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSyslog not implemented")
}
func (UnimplementedSyslogServiceServer) GetDataDefinitions(*emptypb.Empty, SyslogService_GetDataDefinitionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDataDefinitions not implemented")
}
func (UnimplementedSyslogServiceServer) mustEmbedUnimplementedSyslogServiceServer() {}

// UnsafeSyslogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyslogServiceServer will
// result in compilation errors.
type UnsafeSyslogServiceServer interface {
	mustEmbedUnimplementedSyslogServiceServer()
}

func RegisterSyslogServiceServer(s grpc.ServiceRegistrar, srv SyslogServiceServer) {
	s.RegisterService(&SyslogService_ServiceDesc, srv)
}

func _SyslogService_GetLastEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogServiceServer).GetLastEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogService_GetLastEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogServiceServer).GetLastEvent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogService_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyslogFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogServiceServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogService_GetCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogServiceServer).GetCount(ctx, req.(*SyslogFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogService_GetEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyslogFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyslogServiceServer).GetEvent(m, &syslogServiceGetEventServer{stream})
}

type SyslogService_GetEventServer interface {
	Send(*EventMessage) error
	grpc.ServerStream
}

type syslogServiceGetEventServer struct {
	grpc.ServerStream
}

func (x *syslogServiceGetEventServer) Send(m *EventMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _SyslogService_GetSyslog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyslogFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyslogServiceServer).GetSyslog(m, &syslogServiceGetSyslogServer{stream})
}

type SyslogService_GetSyslogServer interface {
	Send(*SyslogMessage) error
	grpc.ServerStream
}

type syslogServiceGetSyslogServer struct {
	grpc.ServerStream
}

func (x *syslogServiceGetSyslogServer) Send(m *SyslogMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _SyslogService_GetDataDefinitions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyslogServiceServer).GetDataDefinitions(m, &syslogServiceGetDataDefinitionsServer{stream})
}

type SyslogService_GetDataDefinitionsServer interface {
	Send(*SyslogDataDefinition) error
	grpc.ServerStream
}

type syslogServiceGetDataDefinitionsServer struct {
	grpc.ServerStream
}

func (x *syslogServiceGetDataDefinitionsServer) Send(m *SyslogDataDefinition) error {
	return x.ServerStream.SendMsg(m)
}

// SyslogService_ServiceDesc is the grpc.ServiceDesc for SyslogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyslogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "syslog.SyslogService",
	HandlerType: (*SyslogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastEvent",
			Handler:    _SyslogService_GetLastEvent_Handler,
		},
		{
			MethodName: "GetCount",
			Handler:    _SyslogService_GetCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvent",
			Handler:       _SyslogService_GetEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSyslog",
			Handler:       _SyslogService_GetSyslog_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDataDefinitions",
			Handler:       _SyslogService_GetDataDefinitions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "syslogservice.proto",
}
