// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package distributedml

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

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	AssignTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	ReceiveResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) AssignTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/distributedml.MasterService/AssignTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ReceiveResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/distributedml.MasterService/ReceiveResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility
type MasterServiceServer interface {
	AssignTask(context.Context, *TaskRequest) (*TaskResponse, error)
	ReceiveResult(context.Context, *ResultRequest) (*ResultResponse, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (UnimplementedMasterServiceServer) AssignTask(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTask not implemented")
}
func (UnimplementedMasterServiceServer) ReceiveResult(context.Context, *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveResult not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_AssignTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).AssignTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributedml.MasterService/AssignTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).AssignTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ReceiveResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ReceiveResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributedml.MasterService/ReceiveResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ReceiveResult(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "distributedml.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignTask",
			Handler:    _MasterService_AssignTask_Handler,
		},
		{
			MethodName: "ReceiveResult",
			Handler:    _MasterService_ReceiveResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/distributedml.proto",
}
