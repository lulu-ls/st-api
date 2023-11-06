// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.0
// source: game/v1/game.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGameAnnouncementList = "/api.game.v1.Game/AnnouncementList"
const OperationGameTaskList = "/api.game.v1.Game/TaskList"
const OperationGameTaskReward = "/api.game.v1.Game/TaskReward"

type GameHTTPServer interface {
	// AnnouncementList 公告列表
	AnnouncementList(context.Context, *AnnouncementListRequest) (*AnnouncementListReply, error)
	// TaskList 任务列表
	TaskList(context.Context, *TaskListRequest) (*TaskListReply, error)
	// TaskReward 兑换任务奖励
	TaskReward(context.Context, *TaskRewardRequest) (*TaskRewardReply, error)
}

func RegisterGameHTTPServer(s *http.Server, srv GameHTTPServer) {
	r := s.Route("/")
	r.POST("/st-games/v1/game/announcement/list", _Game_AnnouncementList0_HTTP_Handler(srv))
	r.POST("/st-games/v1/game/task/list", _Game_TaskList0_HTTP_Handler(srv))
	r.POST("/st-games/v1/game/task/reward", _Game_TaskReward0_HTTP_Handler(srv))
}

func _Game_AnnouncementList0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AnnouncementListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameAnnouncementList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AnnouncementList(ctx, req.(*AnnouncementListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AnnouncementListReply)
		return ctx.Result(200, reply)
	}
}

func _Game_TaskList0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameTaskList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TaskList(ctx, req.(*TaskListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskListReply)
		return ctx.Result(200, reply)
	}
}

func _Game_TaskReward0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TaskRewardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameTaskReward)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TaskReward(ctx, req.(*TaskRewardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TaskRewardReply)
		return ctx.Result(200, reply)
	}
}

type GameHTTPClient interface {
	AnnouncementList(ctx context.Context, req *AnnouncementListRequest, opts ...http.CallOption) (rsp *AnnouncementListReply, err error)
	TaskList(ctx context.Context, req *TaskListRequest, opts ...http.CallOption) (rsp *TaskListReply, err error)
	TaskReward(ctx context.Context, req *TaskRewardRequest, opts ...http.CallOption) (rsp *TaskRewardReply, err error)
}

type GameHTTPClientImpl struct {
	cc *http.Client
}

func NewGameHTTPClient(client *http.Client) GameHTTPClient {
	return &GameHTTPClientImpl{client}
}

func (c *GameHTTPClientImpl) AnnouncementList(ctx context.Context, in *AnnouncementListRequest, opts ...http.CallOption) (*AnnouncementListReply, error) {
	var out AnnouncementListReply
	pattern := "/st-games/v1/game/announcement/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameAnnouncementList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) TaskList(ctx context.Context, in *TaskListRequest, opts ...http.CallOption) (*TaskListReply, error) {
	var out TaskListReply
	pattern := "/st-games/v1/game/task/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameTaskList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) TaskReward(ctx context.Context, in *TaskRewardRequest, opts ...http.CallOption) (*TaskRewardReply, error) {
	var out TaskRewardReply
	pattern := "/st-games/v1/game/task/reward"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameTaskReward))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
