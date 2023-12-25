// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: coupon/v1/coupon.proto

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

const OperationCouponDetail = "/api.coupon.v1.Coupon/Detail"
const OperationCouponHistoryList = "/api.coupon.v1.Coupon/HistoryList"
const OperationCouponVerification = "/api.coupon.v1.Coupon/Verification"

type CouponHTTPServer interface {
	// Detail 获取优惠券详情
	Detail(context.Context, *DetailRequest) (*DetailReply, error)
	// HistoryList 获取优惠券核销记录列表
	HistoryList(context.Context, *HistoryListRequest) (*HistoryListReply, error)
	// Verification 核销优惠券
	Verification(context.Context, *VerificationRequest) (*VerificationReply, error)
}

func RegisterCouponHTTPServer(s *http.Server, srv CouponHTTPServer) {
	r := s.Route("/")
	r.POST("/st-game/v1/coupon/verification", _Coupon_Verification0_HTTP_Handler(srv))
	r.POST("/st-game/v1/coupon/detail", _Coupon_Detail0_HTTP_Handler(srv))
	r.POST("/st-game/v1/coupon/history/list", _Coupon_HistoryList0_HTTP_Handler(srv))
}

func _Coupon_Verification0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerificationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCouponVerification)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Verification(ctx, req.(*VerificationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerificationReply)
		return ctx.Result(200, reply)
	}
}

func _Coupon_Detail0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCouponDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detail(ctx, req.(*DetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetailReply)
		return ctx.Result(200, reply)
	}
}

func _Coupon_HistoryList0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HistoryListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCouponHistoryList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HistoryList(ctx, req.(*HistoryListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HistoryListReply)
		return ctx.Result(200, reply)
	}
}

type CouponHTTPClient interface {
	Detail(ctx context.Context, req *DetailRequest, opts ...http.CallOption) (rsp *DetailReply, err error)
	HistoryList(ctx context.Context, req *HistoryListRequest, opts ...http.CallOption) (rsp *HistoryListReply, err error)
	Verification(ctx context.Context, req *VerificationRequest, opts ...http.CallOption) (rsp *VerificationReply, err error)
}

type CouponHTTPClientImpl struct {
	cc *http.Client
}

func NewCouponHTTPClient(client *http.Client) CouponHTTPClient {
	return &CouponHTTPClientImpl{client}
}

func (c *CouponHTTPClientImpl) Detail(ctx context.Context, in *DetailRequest, opts ...http.CallOption) (*DetailReply, error) {
	var out DetailReply
	pattern := "/st-game/v1/coupon/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCouponDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CouponHTTPClientImpl) HistoryList(ctx context.Context, in *HistoryListRequest, opts ...http.CallOption) (*HistoryListReply, error) {
	var out HistoryListReply
	pattern := "/st-game/v1/coupon/history/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCouponHistoryList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CouponHTTPClientImpl) Verification(ctx context.Context, in *VerificationRequest, opts ...http.CallOption) (*VerificationReply, error) {
	var out VerificationReply
	pattern := "/st-game/v1/coupon/verification"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCouponVerification))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}