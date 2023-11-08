// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.0
// source: wx/v1/message.proto

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

const OperationMessagePayNotify = "/api.wx.v1.Message/PayNotify"
const OperationMessageWxJSPayHtml = "/api.wx.v1.Message/WxJSPayHtml"
const OperationMessageWxService = "/api.wx.v1.Message/WxService"

type MessageHTTPServer interface {
	PayNotify(context.Context, *PayNotifyRequest) (*PayNotifyReply, error)
	WxJSPayHtml(context.Context, *WxJSPayHtmlRequest) (*WxJSPayHtmlReply, error)
	WxService(context.Context, *WxServiceRequest) (*WxServiceReply, error)
}

func RegisterMessageHTTPServer(s *http.Server, srv MessageHTTPServer) {
	r := s.Route("/")
	r.POST("/st-games/v1/wx/notify/pay", _Message_PayNotify0_HTTP_Handler(srv))
	r.GET("/st-games/v1/wx/service", _Message_WxService0_HTTP_Handler(srv))
	r.POST("/st-games/v1/wx/service", _Message_WxService1_HTTP_Handler(srv))
	r.GET("/st-games/v1/wx/html", _Message_WxJSPayHtml0_HTTP_Handler(srv))
}

func _Message_PayNotify0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PayNotifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessagePayNotify)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PayNotify(ctx, req.(*PayNotifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PayNotifyReply)
		return ctx.Result(200, reply)
	}
}

func _Message_WxService0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageWxService)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxService(ctx, req.(*WxServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxServiceReply)
		return ctx.Result(200, reply)
	}
}

func _Message_WxService1_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxServiceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageWxService)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxService(ctx, req.(*WxServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxServiceReply)
		return ctx.Result(200, reply)
	}
}

func _Message_WxJSPayHtml0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxJSPayHtmlRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessageWxJSPayHtml)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxJSPayHtml(ctx, req.(*WxJSPayHtmlRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxJSPayHtmlReply)
		return ctx.Result(200, reply)
	}
}

type MessageHTTPClient interface {
	PayNotify(ctx context.Context, req *PayNotifyRequest, opts ...http.CallOption) (rsp *PayNotifyReply, err error)
	WxJSPayHtml(ctx context.Context, req *WxJSPayHtmlRequest, opts ...http.CallOption) (rsp *WxJSPayHtmlReply, err error)
	WxService(ctx context.Context, req *WxServiceRequest, opts ...http.CallOption) (rsp *WxServiceReply, err error)
}

type MessageHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageHTTPClient(client *http.Client) MessageHTTPClient {
	return &MessageHTTPClientImpl{client}
}

func (c *MessageHTTPClientImpl) PayNotify(ctx context.Context, in *PayNotifyRequest, opts ...http.CallOption) (*PayNotifyReply, error) {
	var out PayNotifyReply
	pattern := "/st-games/v1/wx/notify/pay"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessagePayNotify))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) WxJSPayHtml(ctx context.Context, in *WxJSPayHtmlRequest, opts ...http.CallOption) (*WxJSPayHtmlReply, error) {
	var out WxJSPayHtmlReply
	pattern := "/st-games/v1/wx/html"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMessageWxJSPayHtml))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) WxService(ctx context.Context, in *WxServiceRequest, opts ...http.CallOption) (*WxServiceReply, error) {
	var out WxServiceReply
	pattern := "/st-games/v1/wx/service"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessageWxService))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
