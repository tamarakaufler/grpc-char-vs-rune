package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/status"
	"google.golang.org/grpc"
)

// Server dictates what methods/APIs this service must provide.
type Server interface {
	Options() *Options
	Server() *grpc.Server
	Start() error
	Stop() error
}

type service struct {
	options *Options
	server  *grpc.Server
	status  *status.Probe

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

	l := logrus.New().WithFields(logrus.Fields{
		"service_name":    s.options.Name,
		"service_version": s.options.Version,
	})
	s.status = status.LivenessReadinessServer(l, func() error {
		for _, f := range options.readinessChecks {
			if err := f(); err != nil {
				return err
			}
		}
		return nil
	})
	s.status.Port = s.options.StatusPort

	return s, nil
}

// Options is an options field getter.
func (s *service) Options() *Options {
	return s.options
}

// Server is a server field getter.
func (s *service) Server() *grpc.Server {
	return s.server
}

// Start starts the gRPC server.
func (s *service) Start() error {
	if s.status != nil {
		go s.status.Start()
	}

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
