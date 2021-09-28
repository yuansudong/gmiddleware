package gmiddleware

import (
	"context"

	"google.golang.org/grpc"
)

// ServiceInfo 用于描述一个服务信息
type ServiceInfo struct {
	Method  string // Method 用于描述当前的方法名
	Service string //  Service 用于描述一个服务名
	Version string // Version 版本
}

// FuncHandler 用于描述一个执行句柄
type FuncHandler func(ctx context.Context, req interface{}) (rsp interface{}, err error)

// ChanFn 用于描述一个执行链
type ChanFn func(ctx context.Context, req interface{}, info *ServiceInfo, handler FuncHandler) (rsp interface{}, err error)

// ChainUnaryServer 服务端Unary调用链
func ChainUnaryServer(interceptors ...ChanFn) ChanFn {
	n := len(interceptors)
	return func(ctx context.Context, req interface{}, info *ServiceInfo, handler FuncHandler) (interface{}, error) {
		chainer := func(currentInter ChanFn, currentHandler FuncHandler) FuncHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}
		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}
		return chainedHandler(ctx, req)
	}
}

// ChainStreamServer  服务端的流式调用链
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)

	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		chainer := func(currentInter grpc.StreamServerInterceptor, currentHandler grpc.StreamHandler) grpc.StreamHandler {
			return func(currentSrv interface{}, currentStream grpc.ServerStream) error {
				return currentInter(currentSrv, currentStream, info, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(srv, ss)
	}
}

// ChainUnaryClient 客户端的Unary调用链
func ChainUnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		chainer := func(currentInter grpc.UnaryClientInterceptor, currentInvoker grpc.UnaryInvoker) grpc.UnaryInvoker {
			return func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
				return currentInter(currentCtx, currentMethod, currentReq, currentRepl, currentConn, currentInvoker, currentOpts...)
			}
		}

		chainedInvoker := invoker
		for i := n - 1; i >= 0; i-- {
			chainedInvoker = chainer(interceptors[i], chainedInvoker)
		}

		return chainedInvoker(ctx, method, req, reply, cc, opts...)
	}
}
