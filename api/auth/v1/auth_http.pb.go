// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: auth/v1/auth.proto

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

const OperationAuthDecrypt = "/api.auth.v1.Auth/Decrypt"
const OperationAuthGetInfo = "/api.auth.v1.Auth/GetInfo"
const OperationAuthLogin = "/api.auth.v1.Auth/Login"
const OperationAuthLoginByApple = "/api.auth.v1.Auth/LoginByApple"
const OperationAuthLoginForApp = "/api.auth.v1.Auth/LoginForApp"
const OperationAuthLoginTest = "/api.auth.v1.Auth/LoginTest"
const OperationAuthSendCode = "/api.auth.v1.Auth/SendCode"
const OperationAuthVerifyBindCode = "/api.auth.v1.Auth/VerifyBindCode"
const OperationAuthVerifyLoginCode = "/api.auth.v1.Auth/VerifyLoginCode"

type AuthHTTPServer interface {
	// Decrypt 解密
	Decrypt(context.Context, *DecryptRequest) (*DecryptReply, error)
	// GetInfo 获取登录信息
	GetInfo(context.Context, *GetInfoRequest) (*LoginReply, error)
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// LoginByApple 苹果登录
	LoginByApple(context.Context, *LoginByAppleRequest) (*LoginReply, error)
	// LoginForApp app 登录
	LoginForApp(context.Context, *LoginForAppRequest) (*LoginForAppReply, error)
	// LoginTest 测试登录
	LoginTest(context.Context, *LoginTestRequest) (*LoginReply, error)
	// SendCode 注册短信
	SendCode(context.Context, *SendCodeRequest) (*SendCodeReply, error)
	// VerifyBindCode 验证绑定短信
	VerifyBindCode(context.Context, *VerifyBindCodeRequest) (*LoginReply, error)
	// VerifyLoginCode 验证登录短信
	VerifyLoginCode(context.Context, *VerifyLoginCodeRequest) (*LoginReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/st-games/v1/auth/login", _Auth_Login0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/login/app", _Auth_LoginForApp0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/login_test", _Auth_LoginTest0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/decrypt", _Auth_Decrypt0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/get_info", _Auth_GetInfo0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/code/send", _Auth_SendCode0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/code/login/verify", _Auth_VerifyLoginCode0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/code/bind/verify", _Auth_VerifyBindCode0_HTTP_Handler(srv))
	r.POST("/st-games/v1/auth/login/apple", _Auth_LoginByApple0_HTTP_Handler(srv))
}

func _Auth_Login0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthLogin)
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

func _Auth_LoginForApp0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginForAppRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthLoginForApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginForApp(ctx, req.(*LoginForAppRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginForAppReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_LoginTest0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginTestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthLoginTest)
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

func _Auth_Decrypt0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DecryptRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthDecrypt)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Decrypt(ctx, req.(*DecryptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DecryptReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetInfo0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetInfo(ctx, req.(*GetInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SendCode0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSendCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendCode(ctx, req.(*SendCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendCodeReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_VerifyLoginCode0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyLoginCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthVerifyLoginCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyLoginCode(ctx, req.(*VerifyLoginCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_VerifyBindCode0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyBindCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthVerifyBindCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyBindCode(ctx, req.(*VerifyBindCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_LoginByApple0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginByAppleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthLoginByApple)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByApple(ctx, req.(*LoginByAppleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	Decrypt(ctx context.Context, req *DecryptRequest, opts ...http.CallOption) (rsp *DecryptReply, err error)
	GetInfo(ctx context.Context, req *GetInfoRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginByApple(ctx context.Context, req *LoginByAppleRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginForApp(ctx context.Context, req *LoginForAppRequest, opts ...http.CallOption) (rsp *LoginForAppReply, err error)
	LoginTest(ctx context.Context, req *LoginTestRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	SendCode(ctx context.Context, req *SendCodeRequest, opts ...http.CallOption) (rsp *SendCodeReply, err error)
	VerifyBindCode(ctx context.Context, req *VerifyBindCodeRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	VerifyLoginCode(ctx context.Context, req *VerifyLoginCodeRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) Decrypt(ctx context.Context, in *DecryptRequest, opts ...http.CallOption) (*DecryptReply, error) {
	var out DecryptReply
	pattern := "/st-games/v1/auth/decrypt"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthDecrypt))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-games/v1/auth/get_info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthGetInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-games/v1/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) LoginByApple(ctx context.Context, in *LoginByAppleRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-games/v1/auth/login/apple"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthLoginByApple))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) LoginForApp(ctx context.Context, in *LoginForAppRequest, opts ...http.CallOption) (*LoginForAppReply, error) {
	var out LoginForAppReply
	pattern := "/st-games/v1/auth/login/app"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthLoginForApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) LoginTest(ctx context.Context, in *LoginTestRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-games/v1/auth/login_test"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthLoginTest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) SendCode(ctx context.Context, in *SendCodeRequest, opts ...http.CallOption) (*SendCodeReply, error) {
	var out SendCodeReply
	pattern := "/st-games/v1/auth/code/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSendCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) VerifyBindCode(ctx context.Context, in *VerifyBindCodeRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-games/v1/auth/code/bind/verify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthVerifyBindCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) VerifyLoginCode(ctx context.Context, in *VerifyLoginCodeRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/st-games/v1/auth/code/login/verify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthVerifyLoginCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
