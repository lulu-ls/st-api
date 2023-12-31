// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: wx/v1/deal.proto

package v1

import (
	context "context"
	v1 "github.com/lulu-ls/st-api/api/order/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Deal_MiniPay_FullMethodName    = "/api.wx.v1.Deal/MiniPay"
	Deal_JSPay_FullMethodName      = "/api.wx.v1.Deal/JSPay"
	Deal_VirtualPay_FullMethodName = "/api.wx.v1.Deal/VirtualPay"
	Deal_AppPay_FullMethodName     = "/api.wx.v1.Deal/AppPay"
)

// DealClient is the client API for Deal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DealClient interface {
	MiniPay(ctx context.Context, in *MiniPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error)
	JSPay(ctx context.Context, in *JSPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error)
	VirtualPay(ctx context.Context, in *VirtualPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error)
	AppPay(ctx context.Context, in *AppPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error)
}

type dealClient struct {
	cc grpc.ClientConnInterface
}

func NewDealClient(cc grpc.ClientConnInterface) DealClient {
	return &dealClient{cc}
}

func (c *dealClient) MiniPay(ctx context.Context, in *MiniPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error) {
	out := new(v1.ConfirmPayReply)
	err := c.cc.Invoke(ctx, Deal_MiniPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealClient) JSPay(ctx context.Context, in *JSPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error) {
	out := new(v1.ConfirmPayReply)
	err := c.cc.Invoke(ctx, Deal_JSPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealClient) VirtualPay(ctx context.Context, in *VirtualPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error) {
	out := new(v1.ConfirmPayReply)
	err := c.cc.Invoke(ctx, Deal_VirtualPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealClient) AppPay(ctx context.Context, in *AppPayRequest, opts ...grpc.CallOption) (*v1.ConfirmPayReply, error) {
	out := new(v1.ConfirmPayReply)
	err := c.cc.Invoke(ctx, Deal_AppPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DealServer is the server API for Deal service.
// All implementations must embed UnimplementedDealServer
// for forward compatibility
type DealServer interface {
	MiniPay(context.Context, *MiniPayRequest) (*v1.ConfirmPayReply, error)
	JSPay(context.Context, *JSPayRequest) (*v1.ConfirmPayReply, error)
	VirtualPay(context.Context, *VirtualPayRequest) (*v1.ConfirmPayReply, error)
	AppPay(context.Context, *AppPayRequest) (*v1.ConfirmPayReply, error)
	mustEmbedUnimplementedDealServer()
}

// UnimplementedDealServer must be embedded to have forward compatible implementations.
type UnimplementedDealServer struct {
}

func (UnimplementedDealServer) MiniPay(context.Context, *MiniPayRequest) (*v1.ConfirmPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniPay not implemented")
}
func (UnimplementedDealServer) JSPay(context.Context, *JSPayRequest) (*v1.ConfirmPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JSPay not implemented")
}
func (UnimplementedDealServer) VirtualPay(context.Context, *VirtualPayRequest) (*v1.ConfirmPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VirtualPay not implemented")
}
func (UnimplementedDealServer) AppPay(context.Context, *AppPayRequest) (*v1.ConfirmPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppPay not implemented")
}
func (UnimplementedDealServer) mustEmbedUnimplementedDealServer() {}

// UnsafeDealServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DealServer will
// result in compilation errors.
type UnsafeDealServer interface {
	mustEmbedUnimplementedDealServer()
}

func RegisterDealServer(s grpc.ServiceRegistrar, srv DealServer) {
	s.RegisterService(&Deal_ServiceDesc, srv)
}

func _Deal_MiniPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiniPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealServer).MiniPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deal_MiniPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealServer).MiniPay(ctx, req.(*MiniPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deal_JSPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JSPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealServer).JSPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deal_JSPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealServer).JSPay(ctx, req.(*JSPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deal_VirtualPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealServer).VirtualPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deal_VirtualPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealServer).VirtualPay(ctx, req.(*VirtualPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deal_AppPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealServer).AppPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Deal_AppPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealServer).AppPay(ctx, req.(*AppPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deal_ServiceDesc is the grpc.ServiceDesc for Deal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.wx.v1.Deal",
	HandlerType: (*DealServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MiniPay",
			Handler:    _Deal_MiniPay_Handler,
		},
		{
			MethodName: "JSPay",
			Handler:    _Deal_JSPay_Handler,
		},
		{
			MethodName: "VirtualPay",
			Handler:    _Deal_VirtualPay_Handler,
		},
		{
			MethodName: "AppPay",
			Handler:    _Deal_AppPay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wx/v1/deal.proto",
}
