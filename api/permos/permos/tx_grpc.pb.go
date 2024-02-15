// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: permos/permos/tx.proto

package permos

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

const (
	Msg_UpdateParams_FullMethodName = "/permos.permos.Msg/UpdateParams"
	Msg_CreatePermos_FullMethodName = "/permos.permos.Msg/CreatePermos"
	Msg_UpdatePermos_FullMethodName = "/permos.permos.Msg/UpdatePermos"
	Msg_DeletePermos_FullMethodName = "/permos.permos.Msg/DeletePermos"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreatePermos(ctx context.Context, in *MsgCreatePermos, opts ...grpc.CallOption) (*MsgCreatePermosResponse, error)
	UpdatePermos(ctx context.Context, in *MsgUpdatePermos, opts ...grpc.CallOption) (*MsgUpdatePermosResponse, error)
	DeletePermos(ctx context.Context, in *MsgDeletePermos, opts ...grpc.CallOption) (*MsgDeletePermosResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePermos(ctx context.Context, in *MsgCreatePermos, opts ...grpc.CallOption) (*MsgCreatePermosResponse, error) {
	out := new(MsgCreatePermosResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePermos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePermos(ctx context.Context, in *MsgUpdatePermos, opts ...grpc.CallOption) (*MsgUpdatePermosResponse, error) {
	out := new(MsgUpdatePermosResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePermos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePermos(ctx context.Context, in *MsgDeletePermos, opts ...grpc.CallOption) (*MsgDeletePermosResponse, error) {
	out := new(MsgDeletePermosResponse)
	err := c.cc.Invoke(ctx, Msg_DeletePermos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreatePermos(context.Context, *MsgCreatePermos) (*MsgCreatePermosResponse, error)
	UpdatePermos(context.Context, *MsgUpdatePermos) (*MsgUpdatePermosResponse, error)
	DeletePermos(context.Context, *MsgDeletePermos) (*MsgDeletePermosResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePermos(context.Context, *MsgCreatePermos) (*MsgCreatePermosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermos not implemented")
}
func (UnimplementedMsgServer) UpdatePermos(context.Context, *MsgUpdatePermos) (*MsgUpdatePermosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermos not implemented")
}
func (UnimplementedMsgServer) DeletePermos(context.Context, *MsgDeletePermos) (*MsgDeletePermosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermos not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePermos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePermos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePermos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePermos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePermos(ctx, req.(*MsgCreatePermos))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePermos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePermos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePermos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePermos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePermos(ctx, req.(*MsgUpdatePermos))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePermos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePermos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePermos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeletePermos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePermos(ctx, req.(*MsgDeletePermos))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permos.permos.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePermos",
			Handler:    _Msg_CreatePermos_Handler,
		},
		{
			MethodName: "UpdatePermos",
			Handler:    _Msg_UpdatePermos_Handler,
		},
		{
			MethodName: "DeletePermos",
			Handler:    _Msg_DeletePermos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permos/permos/tx.proto",
}