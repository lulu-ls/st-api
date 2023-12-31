// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: race/v1/race.proto

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
	Race_RaceType_FullMethodName              = "/api.race.v1.Race/RaceType"
	Race_RaceList_FullMethodName              = "/api.race.v1.Race/RaceList"
	Race_RaceDetail_FullMethodName            = "/api.race.v1.Race/RaceDetail"
	Race_RaceReward_FullMethodName            = "/api.race.v1.Race/RaceReward"
	Race_RaceRule_FullMethodName              = "/api.race.v1.Race/RaceRule"
	Race_RaceSignup_FullMethodName            = "/api.race.v1.Race/RaceSignup"
	Race_RaceSignupCancel_FullMethodName      = "/api.race.v1.Race/RaceSignupCancel"
	Race_RaceBatchSignupCancel_FullMethodName = "/api.race.v1.Race/RaceBatchSignupCancel"
	Race_Podium_FullMethodName                = "/api.race.v1.Race/Podium"
	Race_GameSeriesBindUser_FullMethodName    = "/api.race.v1.Race/GameSeriesBindUser"
	Race_GameRaceCheck_FullMethodName         = "/api.race.v1.Race/GameRaceCheck"
)

// RaceClient is the client API for Race service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaceClient interface {
	// 比赛侧边栏分类
	RaceType(ctx context.Context, in *RaceTypeRequest, opts ...grpc.CallOption) (*RaceTypeReply, error)
	// 比赛列表
	RaceList(ctx context.Context, in *RaceListRequest, opts ...grpc.CallOption) (*RaceListReply, error)
	// 比赛详情
	RaceDetail(ctx context.Context, in *RaceDetailRequest, opts ...grpc.CallOption) (*ListRaceItem, error)
	// 比赛奖励
	RaceReward(ctx context.Context, in *RaceRewardRequest, opts ...grpc.CallOption) (*RaceRewardReply, error)
	// 赛制
	RaceRule(ctx context.Context, in *RaceRuleRequest, opts ...grpc.CallOption) (*RaceRuleReply, error)
	// 报名
	RaceSignup(ctx context.Context, in *RaceSignupRequest, opts ...grpc.CallOption) (*RaceSignupReply, error)
	// 取消报名
	RaceSignupCancel(ctx context.Context, in *RaceSignupCancelRequest, opts ...grpc.CallOption) (*RaceSignupCancelReply, error)
	// 批量取消报名（服务端接口）
	RaceBatchSignupCancel(ctx context.Context, in *RaceBatchSignupCancelRequest, opts ...grpc.CallOption) (*RaceBatchSignupCancelReply, error)
	// 领奖台
	Podium(ctx context.Context, in *PodiumRequest, opts ...grpc.CallOption) (*PodiumReply, error)
	// 用户绑定赛事系列
	GameSeriesBindUser(ctx context.Context, in *GameSeriesBindUserRequest, opts ...grpc.CallOption) (*GameSeriesBindUserReply, error)
	// 人满赛检查是否符合开赛条件
	GameRaceCheck(ctx context.Context, in *GameRaceCheckRequest, opts ...grpc.CallOption) (*GameRaceCheckReply, error)
}

type raceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaceClient(cc grpc.ClientConnInterface) RaceClient {
	return &raceClient{cc}
}

func (c *raceClient) RaceType(ctx context.Context, in *RaceTypeRequest, opts ...grpc.CallOption) (*RaceTypeReply, error) {
	out := new(RaceTypeReply)
	err := c.cc.Invoke(ctx, Race_RaceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceList(ctx context.Context, in *RaceListRequest, opts ...grpc.CallOption) (*RaceListReply, error) {
	out := new(RaceListReply)
	err := c.cc.Invoke(ctx, Race_RaceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceDetail(ctx context.Context, in *RaceDetailRequest, opts ...grpc.CallOption) (*ListRaceItem, error) {
	out := new(ListRaceItem)
	err := c.cc.Invoke(ctx, Race_RaceDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceReward(ctx context.Context, in *RaceRewardRequest, opts ...grpc.CallOption) (*RaceRewardReply, error) {
	out := new(RaceRewardReply)
	err := c.cc.Invoke(ctx, Race_RaceReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceRule(ctx context.Context, in *RaceRuleRequest, opts ...grpc.CallOption) (*RaceRuleReply, error) {
	out := new(RaceRuleReply)
	err := c.cc.Invoke(ctx, Race_RaceRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceSignup(ctx context.Context, in *RaceSignupRequest, opts ...grpc.CallOption) (*RaceSignupReply, error) {
	out := new(RaceSignupReply)
	err := c.cc.Invoke(ctx, Race_RaceSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceSignupCancel(ctx context.Context, in *RaceSignupCancelRequest, opts ...grpc.CallOption) (*RaceSignupCancelReply, error) {
	out := new(RaceSignupCancelReply)
	err := c.cc.Invoke(ctx, Race_RaceSignupCancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) RaceBatchSignupCancel(ctx context.Context, in *RaceBatchSignupCancelRequest, opts ...grpc.CallOption) (*RaceBatchSignupCancelReply, error) {
	out := new(RaceBatchSignupCancelReply)
	err := c.cc.Invoke(ctx, Race_RaceBatchSignupCancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) Podium(ctx context.Context, in *PodiumRequest, opts ...grpc.CallOption) (*PodiumReply, error) {
	out := new(PodiumReply)
	err := c.cc.Invoke(ctx, Race_Podium_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) GameSeriesBindUser(ctx context.Context, in *GameSeriesBindUserRequest, opts ...grpc.CallOption) (*GameSeriesBindUserReply, error) {
	out := new(GameSeriesBindUserReply)
	err := c.cc.Invoke(ctx, Race_GameSeriesBindUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raceClient) GameRaceCheck(ctx context.Context, in *GameRaceCheckRequest, opts ...grpc.CallOption) (*GameRaceCheckReply, error) {
	out := new(GameRaceCheckReply)
	err := c.cc.Invoke(ctx, Race_GameRaceCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaceServer is the server API for Race service.
// All implementations must embed UnimplementedRaceServer
// for forward compatibility
type RaceServer interface {
	// 比赛侧边栏分类
	RaceType(context.Context, *RaceTypeRequest) (*RaceTypeReply, error)
	// 比赛列表
	RaceList(context.Context, *RaceListRequest) (*RaceListReply, error)
	// 比赛详情
	RaceDetail(context.Context, *RaceDetailRequest) (*ListRaceItem, error)
	// 比赛奖励
	RaceReward(context.Context, *RaceRewardRequest) (*RaceRewardReply, error)
	// 赛制
	RaceRule(context.Context, *RaceRuleRequest) (*RaceRuleReply, error)
	// 报名
	RaceSignup(context.Context, *RaceSignupRequest) (*RaceSignupReply, error)
	// 取消报名
	RaceSignupCancel(context.Context, *RaceSignupCancelRequest) (*RaceSignupCancelReply, error)
	// 批量取消报名（服务端接口）
	RaceBatchSignupCancel(context.Context, *RaceBatchSignupCancelRequest) (*RaceBatchSignupCancelReply, error)
	// 领奖台
	Podium(context.Context, *PodiumRequest) (*PodiumReply, error)
	// 用户绑定赛事系列
	GameSeriesBindUser(context.Context, *GameSeriesBindUserRequest) (*GameSeriesBindUserReply, error)
	// 人满赛检查是否符合开赛条件
	GameRaceCheck(context.Context, *GameRaceCheckRequest) (*GameRaceCheckReply, error)
	mustEmbedUnimplementedRaceServer()
}

// UnimplementedRaceServer must be embedded to have forward compatible implementations.
type UnimplementedRaceServer struct {
}

func (UnimplementedRaceServer) RaceType(context.Context, *RaceTypeRequest) (*RaceTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceType not implemented")
}
func (UnimplementedRaceServer) RaceList(context.Context, *RaceListRequest) (*RaceListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceList not implemented")
}
func (UnimplementedRaceServer) RaceDetail(context.Context, *RaceDetailRequest) (*ListRaceItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceDetail not implemented")
}
func (UnimplementedRaceServer) RaceReward(context.Context, *RaceRewardRequest) (*RaceRewardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceReward not implemented")
}
func (UnimplementedRaceServer) RaceRule(context.Context, *RaceRuleRequest) (*RaceRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceRule not implemented")
}
func (UnimplementedRaceServer) RaceSignup(context.Context, *RaceSignupRequest) (*RaceSignupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceSignup not implemented")
}
func (UnimplementedRaceServer) RaceSignupCancel(context.Context, *RaceSignupCancelRequest) (*RaceSignupCancelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceSignupCancel not implemented")
}
func (UnimplementedRaceServer) RaceBatchSignupCancel(context.Context, *RaceBatchSignupCancelRequest) (*RaceBatchSignupCancelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaceBatchSignupCancel not implemented")
}
func (UnimplementedRaceServer) Podium(context.Context, *PodiumRequest) (*PodiumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Podium not implemented")
}
func (UnimplementedRaceServer) GameSeriesBindUser(context.Context, *GameSeriesBindUserRequest) (*GameSeriesBindUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GameSeriesBindUser not implemented")
}
func (UnimplementedRaceServer) GameRaceCheck(context.Context, *GameRaceCheckRequest) (*GameRaceCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GameRaceCheck not implemented")
}
func (UnimplementedRaceServer) mustEmbedUnimplementedRaceServer() {}

// UnsafeRaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaceServer will
// result in compilation errors.
type UnsafeRaceServer interface {
	mustEmbedUnimplementedRaceServer()
}

func RegisterRaceServer(s grpc.ServiceRegistrar, srv RaceServer) {
	s.RegisterService(&Race_ServiceDesc, srv)
}

func _Race_RaceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceType(ctx, req.(*RaceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceList(ctx, req.(*RaceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceDetail(ctx, req.(*RaceDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceReward(ctx, req.(*RaceRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceRule(ctx, req.(*RaceRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceSignup(ctx, req.(*RaceSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceSignupCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceSignupCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceSignupCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceSignupCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceSignupCancel(ctx, req.(*RaceSignupCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_RaceBatchSignupCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaceBatchSignupCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).RaceBatchSignupCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_RaceBatchSignupCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).RaceBatchSignupCancel(ctx, req.(*RaceBatchSignupCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_Podium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodiumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).Podium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_Podium_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).Podium(ctx, req.(*PodiumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_GameSeriesBindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameSeriesBindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).GameSeriesBindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_GameSeriesBindUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).GameSeriesBindUser(ctx, req.(*GameSeriesBindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Race_GameRaceCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRaceCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServer).GameRaceCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Race_GameRaceCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServer).GameRaceCheck(ctx, req.(*GameRaceCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Race_ServiceDesc is the grpc.ServiceDesc for Race service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Race_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.race.v1.Race",
	HandlerType: (*RaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RaceType",
			Handler:    _Race_RaceType_Handler,
		},
		{
			MethodName: "RaceList",
			Handler:    _Race_RaceList_Handler,
		},
		{
			MethodName: "RaceDetail",
			Handler:    _Race_RaceDetail_Handler,
		},
		{
			MethodName: "RaceReward",
			Handler:    _Race_RaceReward_Handler,
		},
		{
			MethodName: "RaceRule",
			Handler:    _Race_RaceRule_Handler,
		},
		{
			MethodName: "RaceSignup",
			Handler:    _Race_RaceSignup_Handler,
		},
		{
			MethodName: "RaceSignupCancel",
			Handler:    _Race_RaceSignupCancel_Handler,
		},
		{
			MethodName: "RaceBatchSignupCancel",
			Handler:    _Race_RaceBatchSignupCancel_Handler,
		},
		{
			MethodName: "Podium",
			Handler:    _Race_Podium_Handler,
		},
		{
			MethodName: "GameSeriesBindUser",
			Handler:    _Race_GameSeriesBindUser_Handler,
		},
		{
			MethodName: "GameRaceCheck",
			Handler:    _Race_GameRaceCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "race/v1/race.proto",
}
