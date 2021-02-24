package client

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryChain chains unary interceptors processing them right-to-left
// for left-to-right execution.
func UnaryChain(
	interceptors ...grpc.UnaryClientInterceptor,
) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		for i := len(interceptors) - 1; i >= 0; i-- {
			interceptor := interceptors[i]
			invoker = activateInterceptor(interceptor, method, invoker)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// activateInterceptor wraps a unary handler with an unary server interceptor.
func activateInterceptor(
	interceptor grpc.UnaryClientInterceptor,
	method string,
	invoker grpc.UnaryInvoker,
) grpc.UnaryInvoker {
	return func(ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		opts ...grpc.CallOption) error {
		return interceptor(ctx, method, req, reply, cc, invoker, opts...)
	}
}
