package rpc_client

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt4 "github.com/golang-jwt/jwt/v4"
	googleRpc "google.golang.org/grpc"
	"time"
)

type Option struct {
	ServiceName string                  // 服务名称
	ServiceKey  string                  // 服务密钥
	NoAuth      bool                    // 不需要鉴权
	NoRecovery  bool                    // 不需要恢复
	Middleware  []middleware.Middleware // 中间件
	Discover    registry.Discovery      // 服务发现
	TimeOut     time.Duration           // rpc 超时时间
}

// NewClient [T any] ， 创建 rpc client 连接
// 参数：
//
//	opt ： desc
//	NewClientFun ： desc
//
// 返回值：
//
//	T ：desc
func NewClient[T any](opt *Option, NewClientFun func(googleRpc.ClientConnInterface) T) T {

	if opt.ServiceName == "" {
		panic("service name is required")
	}

	if opt.Discover == nil {
		panic("registry discovery is required")
	}

	if !opt.NoRecovery {
		opt.Middleware = append(opt.Middleware, recovery.Recovery())
	}

	if !opt.NoAuth {
		if opt.ServiceKey == "" {
			panic("service key is required")
		}

		opt.Middleware = append(opt.Middleware, jwt.Client(func(token *jwt4.Token) (interface{}, error) {
			return []byte(opt.ServiceKey), nil
		}, jwt.WithSigningMethod(jwt4.SigningMethodHS256)))
	}

	if opt.TimeOut == 0 {
		opt.TimeOut = time.Second * 10
	}

	//if opt.ServiceName == "" {
	//	fmt.Sprintf("discovery:///%s", opt.ServiceName)
	//}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(opt.ServiceName),
		grpc.WithDiscovery(opt.Discover),
		grpc.WithMiddleware(
			opt.Middleware...,
		),
		grpc.WithTimeout(opt.TimeOut),
	)

	if err != nil {
		panic(err)
	}

	return NewClientFun(conn)
}
