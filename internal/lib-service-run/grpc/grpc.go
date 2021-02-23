package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server dictates what methods/APIs this service must provide.
type Server interface {
	Options() *Options
	Server() *grpc.Server
	Run() error
	Stop() error
}

type service struct {
	options *Options
	server  *grpc.Server
	//status  *status.LivenessAndReadinessStatus

	stop chan struct{}
}

// New creates a gRPC server.
func New(opts ...Option) (Server, error) {
	options, err := setOptions(opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failes to create a gRPC server")
	}

	if len(options.Interceptors) == 0 {
		options.Interceptors = []grpc.UnaryServerInterceptor{}
	}

	s := &service{
		options: options,
		server: grpc.NewServer(
			grpc.UnaryInterceptor(
				UnaryChain(
					options.Interceptors...,
				),
			),
		),

		stop: make(chan struct{}),
	}

	return s, nil
}

func (s *service) Options() *Options {
	return s.options
}

func (s *service) Server() *grpc.Server {
	return s.server
}

func (s *service) Run() error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	lis, err := net.Listen("tcp", s.options.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serveErrCh := make(chan error, 1)
	go func() {
		if err := s.server.Serve(lis); err != nil {
			serveErrCh <- errors.Wrap(err, "failed to serve grpc")
		}
	}()

	select {
	case err := <-serveErrCh:
		return err
	case <-s.options.Context.Done():
		s.server.GracefulStop()
	case sig := <-sigCh:
		log.Printf("received signal: %s ... 💥", sig.String())
		s.server.GracefulStop()
	case <-s.stop:
		s.server.GracefulStop()
	}

	return nil
}

func (s *service) Stop() error {
	s.stop <- struct{}{}

	return nil
}

// Check is a verifying function to be run, returning error .
type Check func() error

// ServerCheck verifies the server is running.
// It invokes an unimplemented gRPC method and if the error is not Unimplemented, then
// the server is not running.
func serverCheck(port string) Check {
	return func() error {
		conn, err := grpc.Dial(net.JoinHostPort("127.0.0.1", port), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		err = conn.Invoke(context.Background(), "/ping/pong", &empty.Empty{}, &empty.Empty{})
		s, ok := status.FromError(err)
		if !ok {
			return errors.New("unable to parse grpc error")
		}
		if s.Code() != codes.Unimplemented {
			return errors.Wrapf(err, "expected code %d but got %d", codes.Unimplemented, s.Code())
		}

		return nil
	}
}
