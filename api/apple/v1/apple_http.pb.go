// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: apple/v1/apple.proto

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

const OperationApplePayNotice = "/api.apple.v1.Apple/PayNotice"
const OperationAppleVerifyReceipt = "/api.apple.v1.Apple/VerifyReceipt"

type AppleHTTPServer interface {
	// PayNotice 支付通知
	PayNotice(context.Context, *PayNoticeRequest) (*PayNoticeReply, error)
	VerifyReceipt(context.Context, *VerifyReceiptRequest) (*VerifyReceiptReply, error)
}

func RegisterAppleHTTPServer(s *http.Server, srv AppleHTTPServer) {
	r := s.Route("/")
	r.POST("/st-games/v1/apple/verify_receipt", _Apple_VerifyReceipt0_HTTP_Handler(srv))
	r.POST("/st-games/v1/apple/pay/notice", _Apple_PayNotice0_HTTP_Handler(srv))
}

func _Apple_VerifyReceipt0_HTTP_Handler(srv AppleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyReceiptRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppleVerifyReceipt)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyReceipt(ctx, req.(*VerifyReceiptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyReceiptReply)
		return ctx.Result(200, reply)
	}
}

func _Apple_PayNotice0_HTTP_Handler(srv AppleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PayNoticeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApplePayNotice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PayNotice(ctx, req.(*PayNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PayNoticeReply)
		return ctx.Result(200, reply)
	}
}

type AppleHTTPClient interface {
	PayNotice(ctx context.Context, req *PayNoticeRequest, opts ...http.CallOption) (rsp *PayNoticeReply, err error)
	VerifyReceipt(ctx context.Context, req *VerifyReceiptRequest, opts ...http.CallOption) (rsp *VerifyReceiptReply, err error)
}

type AppleHTTPClientImpl struct {
	cc *http.Client
}

func NewAppleHTTPClient(client *http.Client) AppleHTTPClient {
	return &AppleHTTPClientImpl{client}
}

func (c *AppleHTTPClientImpl) PayNotice(ctx context.Context, in *PayNoticeRequest, opts ...http.CallOption) (*PayNoticeReply, error) {
	var out PayNoticeReply
	pattern := "/st-games/v1/apple/pay/notice"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApplePayNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppleHTTPClientImpl) VerifyReceipt(ctx context.Context, in *VerifyReceiptRequest, opts ...http.CallOption) (*VerifyReceiptReply, error) {
	var out VerifyReceiptReply
	pattern := "/st-games/v1/apple/verify_receipt"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppleVerifyReceipt))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
