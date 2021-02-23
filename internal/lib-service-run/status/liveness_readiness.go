package status

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LivenessReadiness representing Status server object.
type LivenessReadiness struct {
	Port           int
	logger         *logrus.Entry
	liveness       Status
	readinessCheck Check

	lock   sync.RWMutex
	server *http.Server

	wg sync.WaitGroup
}

// LivenessReadinessServer Create a new Status Server, but do not start it.
// Port is overriden by server.New but left 8000 here to avoid a breaking change.
func LivenessReadinessServer(logger *logrus.Entry, readinessCheck Check) *LivenessReadiness {
	return &LivenessReadiness{
		Port:           3000,
		logger:         logger,
		liveness:       UNAVAILABLE,
		readinessCheck: readinessCheck,
	}
}

// Start registers HTTP handlers and start server.
func (s *LivenessReadiness) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status/liveness", s.LivenessProbe)
	mux.HandleFunc("/status/readiness", s.ReadinessProbe)

	srv := &http.Server{Addr: fmt.Sprintf(":%d", s.Port), Handler: mux}

	s.wg.Add(1)
	s.lock.Lock()
	s.server = srv
	s.liveness = RUNNING
	s.lock.Unlock()
	s.wg.Done()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Fatalf("error while serving http on port %d: %+v", s.Port, err)
	}
}

// Stop closes the listener used for running the server.
func (s *LivenessReadiness) Stop() error {
	s.wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// SetLiveness sets liveness of application to be reported on probes.
func (s *LivenessReadiness) SetLiveness(status Status) {
	s.lock.Lock()
	s.liveness = status
	s.lock.Unlock()
}

// LivenessProbe Handler for Kubernetes readiness check and checking service status.
func (s *LivenessReadiness) LivenessProbe(w http.ResponseWriter, r *http.Request) {
	s.lock.RLock()
	code := int(s.liveness)
	s.lock.RUnlock()

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"status": "%s"}`, http.StatusText(code))
}

// ReadinessProbe Handler for Kubernetes readiness check and checking service status.
func (s *LivenessReadiness) ReadinessProbe(w http.ResponseWriter, r *http.Request) {
	var code = int(UNAVAILABLE)
	err := s.readinessCheck()
	if err == nil {
		code = int(RUNNING)
	} else {
		s.logger.WithError(err).Error("failed readiness check")
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"status": "%s"}`, http.StatusText(code))
}

// Check is a verifying function to be run, returning error .
type Check func() error

// ServerCheck verifies the server is running.
// It invokes an unimplemented gRPC method and if the error is not Unimplemented, then
// the server is not running.
func ServerCheck(port string) Check {
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
