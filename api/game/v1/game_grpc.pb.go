// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: game/v1/game.proto

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
	Game_AnnouncementList_FullMethodName = "/api.game.v1.Game/AnnouncementList"
	Game_ActivityList_FullMethodName     = "/api.game.v1.Game/ActivityList"
	Game_TaskDetail_FullMethodName       = "/api.game.v1.Game/TaskDetail"
	Game_TaskReward_FullMethodName       = "/api.game.v1.Game/TaskReward"
	Game_GameInfo_FullMethodName         = "/api.game.v1.Game/GameInfo"
	Game_GetGameNotify_FullMethodName    = "/api.game.v1.Game/GetGameNotify"
	Game_ReadGameNotify_FullMethodName   = "/api.game.v1.Game/ReadGameNotify"
	Game_GetAppConfig_FullMethodName     = "/api.game.v1.Game/GetAppConfig"
	Game_CheckSignup_FullMethodName      = "/api.game.v1.Game/CheckSignup"
	Game_SignList_FullMethodName         = "/api.game.v1.Game/SignList"
	Game_SignIn_FullMethodName           = "/api.game.v1.Game/SignIn"
	Game_RedList_FullMethodName          = "/api.game.v1.Game/RedList"
)

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	// 公告列表
	AnnouncementList(ctx context.Context, in *AnnouncementListRequest, opts ...grpc.CallOption) (*AnnouncementListReply, error)
	// 任务列表
	ActivityList(ctx context.Context, in *ActivityListRequest, opts ...grpc.CallOption) (*ActivityListReply, error)
	// 任务详情
	TaskDetail(ctx context.Context, in *TaskDetailRequest, opts ...grpc.CallOption) (*TaskDetailReply, error)
	// 兑换任务奖励
	TaskReward(ctx context.Context, in *TaskRewardRequest, opts ...grpc.CallOption) (*TaskRewardReply, error)
	// 获取房间信息
	GameInfo(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoReply, error)
	// 获取提示
	GetGameNotify(ctx context.Context, in *GetGameNotifyRequest, opts ...grpc.CallOption) (*GetGameNotifyReply, error)
	// 阅读提示
	ReadGameNotify(ctx context.Context, in *ReadGameNotifyRequest, opts ...grpc.CallOption) (*ReadGameNotifyReply, error)
	// 获取 app 功能配置
	GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...grpc.CallOption) (*GetAppConfigReply, error)
	// 检查报名
	CheckSignup(ctx context.Context, in *CheckSignupRequest, opts ...grpc.CallOption) (*CheckSignupReply, error)
	// 签到列表
	SignList(ctx context.Context, in *SignListRequest, opts ...grpc.CallOption) (*SignListReply, error)
	// 签到
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error)
	// 红点处理
	RedList(ctx context.Context, in *RedListRequest, opts ...grpc.CallOption) (*RedListReply, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) AnnouncementList(ctx context.Context, in *AnnouncementListRequest, opts ...grpc.CallOption) (*AnnouncementListReply, error) {
	out := new(AnnouncementListReply)
	err := c.cc.Invoke(ctx, Game_AnnouncementList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) ActivityList(ctx context.Context, in *ActivityListRequest, opts ...grpc.CallOption) (*ActivityListReply, error) {
	out := new(ActivityListReply)
	err := c.cc.Invoke(ctx, Game_ActivityList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) TaskDetail(ctx context.Context, in *TaskDetailRequest, opts ...grpc.CallOption) (*TaskDetailReply, error) {
	out := new(TaskDetailReply)
	err := c.cc.Invoke(ctx, Game_TaskDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) TaskReward(ctx context.Context, in *TaskRewardRequest, opts ...grpc.CallOption) (*TaskRewardReply, error) {
	out := new(TaskRewardReply)
	err := c.cc.Invoke(ctx, Game_TaskReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GameInfo(ctx context.Context, in *GameInfoRequest, opts ...grpc.CallOption) (*GameInfoReply, error) {
	out := new(GameInfoReply)
	err := c.cc.Invoke(ctx, Game_GameInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetGameNotify(ctx context.Context, in *GetGameNotifyRequest, opts ...grpc.CallOption) (*GetGameNotifyReply, error) {
	out := new(GetGameNotifyReply)
	err := c.cc.Invoke(ctx, Game_GetGameNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) ReadGameNotify(ctx context.Context, in *ReadGameNotifyRequest, opts ...grpc.CallOption) (*ReadGameNotifyReply, error) {
	out := new(ReadGameNotifyReply)
	err := c.cc.Invoke(ctx, Game_ReadGameNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...grpc.CallOption) (*GetAppConfigReply, error) {
	out := new(GetAppConfigReply)
	err := c.cc.Invoke(ctx, Game_GetAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) CheckSignup(ctx context.Context, in *CheckSignupRequest, opts ...grpc.CallOption) (*CheckSignupReply, error) {
	out := new(CheckSignupReply)
	err := c.cc.Invoke(ctx, Game_CheckSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SignList(ctx context.Context, in *SignListRequest, opts ...grpc.CallOption) (*SignListReply, error) {
	out := new(SignListReply)
	err := c.cc.Invoke(ctx, Game_SignList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error) {
	out := new(SignInReply)
	err := c.cc.Invoke(ctx, Game_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) RedList(ctx context.Context, in *RedListRequest, opts ...grpc.CallOption) (*RedListReply, error) {
	out := new(RedListReply)
	err := c.cc.Invoke(ctx, Game_RedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	// 公告列表
	AnnouncementList(context.Context, *AnnouncementListRequest) (*AnnouncementListReply, error)
	// 任务列表
	ActivityList(context.Context, *ActivityListRequest) (*ActivityListReply, error)
	// 任务详情
	TaskDetail(context.Context, *TaskDetailRequest) (*TaskDetailReply, error)
	// 兑换任务奖励
	TaskReward(context.Context, *TaskRewardRequest) (*TaskRewardReply, error)
	// 获取房间信息
	GameInfo(context.Context, *GameInfoRequest) (*GameInfoReply, error)
	// 获取提示
	GetGameNotify(context.Context, *GetGameNotifyRequest) (*GetGameNotifyReply, error)
	// 阅读提示
	ReadGameNotify(context.Context, *ReadGameNotifyRequest) (*ReadGameNotifyReply, error)
	// 获取 app 功能配置
	GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigReply, error)
	// 检查报名
	CheckSignup(context.Context, *CheckSignupRequest) (*CheckSignupReply, error)
	// 签到列表
	SignList(context.Context, *SignListRequest) (*SignListReply, error)
	// 签到
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	// 红点处理
	RedList(context.Context, *RedListRequest) (*RedListReply, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) AnnouncementList(context.Context, *AnnouncementListRequest) (*AnnouncementListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnouncementList not implemented")
}
func (UnimplementedGameServer) ActivityList(context.Context, *ActivityListRequest) (*ActivityListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityList not implemented")
}
func (UnimplementedGameServer) TaskDetail(context.Context, *TaskDetailRequest) (*TaskDetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskDetail not implemented")
}
func (UnimplementedGameServer) TaskReward(context.Context, *TaskRewardRequest) (*TaskRewardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskReward not implemented")
}
func (UnimplementedGameServer) GameInfo(context.Context, *GameInfoRequest) (*GameInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GameInfo not implemented")
}
func (UnimplementedGameServer) GetGameNotify(context.Context, *GetGameNotifyRequest) (*GetGameNotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameNotify not implemented")
}
func (UnimplementedGameServer) ReadGameNotify(context.Context, *ReadGameNotifyRequest) (*ReadGameNotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadGameNotify not implemented")
}
func (UnimplementedGameServer) GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}
func (UnimplementedGameServer) CheckSignup(context.Context, *CheckSignupRequest) (*CheckSignupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSignup not implemented")
}
func (UnimplementedGameServer) SignList(context.Context, *SignListRequest) (*SignListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignList not implemented")
}
func (UnimplementedGameServer) SignIn(context.Context, *SignInRequest) (*SignInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedGameServer) RedList(context.Context, *RedListRequest) (*RedListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedList not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_AnnouncementList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnouncementListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).AnnouncementList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_AnnouncementList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).AnnouncementList(ctx, req.(*AnnouncementListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_ActivityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).ActivityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_ActivityList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).ActivityList(ctx, req.(*ActivityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_TaskDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).TaskDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_TaskDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).TaskDetail(ctx, req.(*TaskDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_TaskReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).TaskReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_TaskReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).TaskReward(ctx, req.(*TaskRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GameInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GameInfo(ctx, req.(*GameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetGameNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetGameNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetGameNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetGameNotify(ctx, req.(*GetGameNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_ReadGameNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadGameNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).ReadGameNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_ReadGameNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).ReadGameNotify(ctx, req.(*ReadGameNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetAppConfig(ctx, req.(*GetAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_CheckSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).CheckSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_CheckSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).CheckSignup(ctx, req.(*CheckSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SignList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SignList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SignList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SignList(ctx, req.(*SignListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_RedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).RedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_RedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).RedList(ctx, req.(*RedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.game.v1.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnnouncementList",
			Handler:    _Game_AnnouncementList_Handler,
		},
		{
			MethodName: "ActivityList",
			Handler:    _Game_ActivityList_Handler,
		},
		{
			MethodName: "TaskDetail",
			Handler:    _Game_TaskDetail_Handler,
		},
		{
			MethodName: "TaskReward",
			Handler:    _Game_TaskReward_Handler,
		},
		{
			MethodName: "GameInfo",
			Handler:    _Game_GameInfo_Handler,
		},
		{
			MethodName: "GetGameNotify",
			Handler:    _Game_GetGameNotify_Handler,
		},
		{
			MethodName: "ReadGameNotify",
			Handler:    _Game_ReadGameNotify_Handler,
		},
		{
			MethodName: "GetAppConfig",
			Handler:    _Game_GetAppConfig_Handler,
		},
		{
			MethodName: "CheckSignup",
			Handler:    _Game_CheckSignup_Handler,
		},
		{
			MethodName: "SignList",
			Handler:    _Game_SignList_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Game_SignIn_Handler,
		},
		{
			MethodName: "RedList",
			Handler:    _Game_RedList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game/v1/game.proto",
}
