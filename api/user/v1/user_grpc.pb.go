// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: user/v1/user.proto

package v1

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
	User_RaceRecordList_FullMethodName = "/api.user.v1.User/RaceRecordList"
	User_InfoEdit_FullMethodName       = "/api.user.v1.User/InfoEdit"
	User_AssetGet_FullMethodName       = "/api.user.v1.User/AssetGet"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 战绩
	RaceRecordList(ctx context.Context, in *RaceRecordListRequest, opts ...grpc.CallOption) (*RaceRecordListReply, error)
	// 修改形象
	InfoEdit(ctx context.Context, in *InfoEditRequest, opts ...grpc.CallOption) (*InfoEditReply, error)
	// 获取用户资产
	AssetGet(ctx context.Context, in *AssetGetRequest, opts ...grpc.CallOption) (*AssetGetReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) RaceRecordList(ctx context.Context, in *RaceRecordListRequest, opts ...grpc.CallOption) (*RaceRecordListReply, error) {
	out := new(RaceRecordListReply)
	err := c.cc.Invoke(ctx, User_RaceRecordList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoEdit(ctx context.Context, in *InfoEditRequest, opts ...grpc.CallOption) (*InfoEditReply, error) {
	out := new(InfoEditReply)
	err := c.cc.Invoke(ctx, User_InfoEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AssetGet(ctx context.Context, in *AssetGetRequest, opts ...grpc.CallOption) (*AssetGetReply, error) {
	out := new(AssetGetReply)
	err := c.cc.Invoke(ctx, User_AssetGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 战绩
	RaceRecordList(context.Context, *RaceRecordListRequest) (*RaceRecordListReply, error)
	// 修改形象
	InfoEdit(context.Context, *InfoEditRequest) (*InfoEditReply, error)
	// 获取用户资产
	AssetGet(context.Context, *AssetGetRequest) (*AssetGetReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) RaceRecordList(context.Context, *RaceRecordListRequest) (*RaceRecordListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceRecordList not implemented")
}
func (UnimplementedUserServer) InfoEdit(context.Context, *InfoEditRequest) (*InfoEditReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoEdit not implemented")
}
func (UnimplementedUserServer) AssetGet(context.Context, *AssetGetRequest) (*AssetGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetGet not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_RaceRecordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceRecordListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RaceRecordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RaceRecordList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RaceRecordList(ctx, req.(*RaceRecordListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_InfoEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).InfoEdit(ctx, req.(*InfoEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AssetGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AssetGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AssetGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AssetGet(ctx, req.(*AssetGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RaceRecordList",
			Handler:    _User_RaceRecordList_Handler,
		},
		{
			MethodName: "InfoEdit",
			Handler:    _User_InfoEdit_Handler,
		},
		{
			MethodName: "AssetGet",
			Handler:    _User_AssetGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
