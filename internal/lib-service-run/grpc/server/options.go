package server

import (
	"context"
	"net"

	"github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/status"
	"google.golang.org/grpc"
)

// Options provides configurable properties of a server.
type Options struct {
	Name         string
	Version      string
	Address      string
	StatusPort   int
	Context      context.Context
	Interceptors []grpc.UnaryServerInterceptor

	readinessChecks []status.Check
}

func setOptions(options ...Option) (*Options, error) {
	o := &Options{
		Name:         "unknown",
		Version:      "unknown",
		Address:      "0.0.0.0:3000",
		Context:      context.Background(),
		Interceptors: []grpc.UnaryServerInterceptor{},
	}
	for _, option := range options {
		option(o)
	}

	o.readinessChecks = []status.Check{o.defaultServerCheck}

	defaultInterceptors := []grpc.UnaryServerInterceptor{}
	o.Interceptors = append(defaultInterceptors, o.Interceptors...)

	return o, nil
}

func (o *Options) defaultServerCheck() error {
	_, port, err := net.SplitHostPort(o.Address)
	if err != nil {
		return err
	}
	return status.ServerCheck(port, "/ping/pong")()
}

// Option is a method for customizing server property.
type Option func(*Options)

// WithName ...
func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

// WithVersion ...
func WithVersion(version string) Option {
	return func(o *Options) {
		o.Version = version
	}
}

// WithAddress ...
func WithAddress(address string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

// WithStatusPort ...
func WithStatusPort(port int) Option {
	return func(o *Options) {
		o.StatusPort = port
	}
}

// WithContext ...
func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// WithInterceptors ...
func WithInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(o *Options) {
		o.Interceptors = interceptors
	}
}
