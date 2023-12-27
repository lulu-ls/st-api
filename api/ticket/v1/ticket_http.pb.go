// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: ticket/v1/ticket.proto

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

const OperationTicketLogin = "/api.ticket.v1.Ticket/Login"
const OperationTicketLoginTest = "/api.ticket.v1.Ticket/LoginTest"
const OperationTicketStatistics = "/api.ticket.v1.Ticket/Statistics"

type TicketHTTPServer interface {
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// LoginTest 登录
	LoginTest(context.Context, *LoginTestRequest) (*LoginReply, error)
	// Statistics 首页分析数据
	Statistics(context.Context, *StatisticsRequest) (*StatisticsReply, error)
}

func RegisterTicketHTTPServer(s *http.Server, srv TicketHTTPServer) {
	r := s.Route("/")
	r.POST("/st-game/v1/ticket/login", _Ticket_Login0_HTTP_Handler(srv))
	r.POST("/st-game/v1/ticket/login_test", _Ticket_LoginTest0_HTTP_Handler(srv))
	r.POST("/st-game/v1/ticket/statistics", _Ticket_Statistics0_HTTP_Handler(srv))
}

func _Ticket_Login0_HTTP_Handler(srv TicketHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Ticket_LoginTest0_HTTP_Handler(srv TicketHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginTestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketLoginTest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginTest(ctx, req.(*LoginTestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Ticket_Statistics0_HTTP_Handler(srv TicketHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatisticsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketStatistics)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Statistics(ctx, req.(*StatisticsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StatisticsReply)
		return ctx.Result(200, reply)
	}
}

type TicketHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginTest(ctx context.Context, req *LoginTestRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Statistics(ctx context.Context, req *StatisticsRequest, opts ...http.CallOption) (rsp *StatisticsReply, err error)
}

type TicketHTTPClientImpl struct {
	cc *http.Client
}

func NewTicketHTTPClient(client *http.Client) TicketHTTPClient {
	return &TicketHTTPClientImpl{client}
}

func (c *TicketHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-game/v1/ticket/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTicketLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketHTTPClientImpl) LoginTest(ctx context.Context, in *LoginTestRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-game/v1/ticket/login_test"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTicketLoginTest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketHTTPClientImpl) Statistics(ctx context.Context, in *StatisticsRequest, opts ...http.CallOption) (*StatisticsReply, error) {
	var out StatisticsReply
	pattern := "/st-game/v1/ticket/statistics"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTicketStatistics))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
