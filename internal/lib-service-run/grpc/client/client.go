package client

import (
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client interface prescribes methods to be implemented to
// create a grpc client instance.
type Client interface {
	Options() *Options
	Conn() *grpc.ClientConn
	Close() error
}

type client struct {
	options *Options
	conn    *grpc.ClientConn
}

var _ Client = (*client)(nil)

func (c *client) Options() *Options {
	return c.options
}

func (c *client) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *client) Close() error {
	return c.conn.Close()
}

// New creates a new grpc Client.
func New(opts ...Option) (Client, error) {
	options := setOptions(opts...)
	if options.Address == "" {
		return nil, errors.New("missing address")
	}

	callOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			UnaryChain(
				options.Interceptors...,
			),
		),
		grpc.WithConnectParams(grpc.ConnectParams{
			// fail if client connection is not implemented by this time
			MinConnectTimeout: 5 * time.Second,
		}),
	}

	conn, err := grpc.DialContext(options.Context, options.Address, callOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to the target service")
	}
	client := &client{
		options: options,
		conn:    conn,
	}
	return client, nil
}
