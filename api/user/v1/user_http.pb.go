// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.4
// source: user/v1/user.proto

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

const OperationUserAssetGet = "/api.user.v1.User/AssetGet"
const OperationUserInfoEdit = "/api.user.v1.User/InfoEdit"
const OperationUserRaceRecordList = "/api.user.v1.User/RaceRecordList"

type UserHTTPServer interface {
	// AssetGet 获取用户资产
	AssetGet(context.Context, *AssetGetRequest) (*AssetGetReply, error)
	// InfoEdit 修改形象
	InfoEdit(context.Context, *InfoEditRequest) (*InfoEditReply, error)
	// RaceRecordList 战绩
	RaceRecordList(context.Context, *RaceRecordListRequest) (*RaceRecordListReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/st-games/v1/user/race/record", _User_RaceRecordList0_HTTP_Handler(srv))
	r.POST("/st-games/v1/user/info/edit", _User_InfoEdit0_HTTP_Handler(srv))
	r.POST("/st-games/v1/user/asset/get", _User_AssetGet0_HTTP_Handler(srv))
}

func _User_RaceRecordList0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RaceRecordListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRaceRecordList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RaceRecordList(ctx, req.(*RaceRecordListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RaceRecordListReply)
		return ctx.Result(200, reply)
	}
}

func _User_InfoEdit0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InfoEditRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserInfoEdit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InfoEdit(ctx, req.(*InfoEditRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InfoEditReply)
		return ctx.Result(200, reply)
	}
}

func _User_AssetGet0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetGetRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAssetGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AssetGet(ctx, req.(*AssetGetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetGetReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	AssetGet(ctx context.Context, req *AssetGetRequest, opts ...http.CallOption) (rsp *AssetGetReply, err error)
	InfoEdit(ctx context.Context, req *InfoEditRequest, opts ...http.CallOption) (rsp *InfoEditReply, err error)
	RaceRecordList(ctx context.Context, req *RaceRecordListRequest, opts ...http.CallOption) (rsp *RaceRecordListReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) AssetGet(ctx context.Context, in *AssetGetRequest, opts ...http.CallOption) (*AssetGetReply, error) {
	var out AssetGetReply
	pattern := "/st-games/v1/user/asset/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAssetGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) InfoEdit(ctx context.Context, in *InfoEditRequest, opts ...http.CallOption) (*InfoEditReply, error) {
	var out InfoEditReply
	pattern := "/st-games/v1/user/info/edit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserInfoEdit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RaceRecordList(ctx context.Context, in *RaceRecordListRequest, opts ...http.CallOption) (*RaceRecordListReply, error) {
	var out RaceRecordListReply
	pattern := "/st-games/v1/user/race/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRaceRecordList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
