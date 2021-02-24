package server

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryChain chains unary interceptors processing them right-to-left
// for left-to-right execution.
func UnaryChain(
	interceptors ...grpc.UnaryServerInterceptor,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		for i := len(interceptors) - 1; i >= 0; i-- {
			interceptor := interceptors[i]
			handler = activateInterceptor(interceptor, info, handler)
		}
		return handler(ctx, req)
	}
}

// activateInterceptor wraps a unary handler with an unary server interceptor.
func activateInterceptor(
	interceptor grpc.UnaryServerInterceptor,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptor(ctx, req, info, handler)
	}
}
