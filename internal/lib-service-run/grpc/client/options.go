package client

import (
	"context"

	"google.golang.org/grpc"
)

// Options holds all the possible configuration for the client.
type Options struct {
	TargetName   string
	CallerName   string
	Address      string
	Context      context.Context
	Interceptors []grpc.UnaryClientInterceptor
}

// Option is a functional construct to set client properties.
type Option func(*Options)

func setOptions(options ...Option) *Options {
	o := &Options{
		TargetName: "unknown",
		CallerName: "unknown",
		Address:    "unknown",
		Context:    context.Background(),
	}
	for _, option := range options {
		option(o)
	}

	defaultInterceptors := []grpc.UnaryClientInterceptor{}
	o.Interceptors = append(defaultInterceptors, o.Interceptors...)

	return o
}

// WithTargetName sets the target name, ie the name of the service that the client is calling.
func WithTargetName(targetName string) Option {
	return func(o *Options) {
		o.TargetName = targetName
	}
}

// WithCallerName sets the caller name, ie name of the service that the client is used from.
func WithCallerName(callerName string) Option {
	return func(o *Options) {
		o.CallerName = callerName
	}
}

// WithAddress sets the grpc target address.
func WithAddress(address string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

// WithContext sets the context for the client.
func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// AddInterceptor appends the provided interceptor to the list of interceptors.
func AddInterceptor(interceptor grpc.UnaryClientInterceptor) Option {
	return func(o *Options) {
		o.Interceptors = append(o.Interceptors, interceptor)
	}
}
