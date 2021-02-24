package status

import (
	"context"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Check is a verifying function to be run, returning error .
type Check func() error

// ServerCheck verifies the gRPC server is running.
// It invokes an unimplemented gRPC method and if the error is not Unimplemented, then
// the server is not running.
// TODO:
// This needs an implementation of a helper script, that will envoke this function. The script
// can be used as a Kubernetes liveness/readiness probe.
// endpoint must start with a slash, eg /ping/pong.
func ServerCheck(port, endpoint string) Check {
	return func() error {
		conn, err := grpc.Dial(net.JoinHostPort("127.0.0.1", port), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		err = conn.Invoke(context.Background(), endpoint, &empty.Empty{}, &empty.Empty{})
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
