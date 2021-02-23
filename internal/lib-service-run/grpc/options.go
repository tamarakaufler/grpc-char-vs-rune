package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

// Options provides configurable properties of a server.
type Options struct {
	Name         string
	Version      string
	Address      string
	Context      context.Context
	Interceptors []grpc.UnaryServerInterceptor

	readinessChecks []Check
}

func setOptions(options ...Option) (*Options, error) {
	o := &Options{
		Name:         "unknown",
		Version:      "unknown",
		Address:      "0.0.0.0:3000",
		Context:      context.Background(),
		Interceptors: []grpc.UnaryServerInterceptor{},
	}
	o.readinessChecks = []Check{o.defaultServerCheck}

	for _, option := range options {
		option(o)
	}

	return o, nil
}

func (o *Options) defaultServerCheck() error {
	_, port, err := net.SplitHostPort(o.Address)
	if err != nil {
		return err
	}
	return serverCheck(port)()
}

// Option is a method for customizing server property.
type Option func(*Options)

func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func WithVersion(version string) Option {
	return func(o *Options) {
		o.Version = version
	}
}

func WithAddress(address string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func WithInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(o *Options) {
		o.Interceptors = interceptors
	}
}
