package status

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Probe represents liveneass and readiness probe object.
type Probe struct {
	Port           int
	logger         *logrus.Entry
	liveness       Status
	readinessCheck Check

	lock   sync.RWMutex
	server *http.Server

	wg sync.WaitGroup
}

// LivenessReadinessServer creates LivenessReadiness instance.
// NOTE: This is not great as the application will be serving on a different port, so
// the LivenessReadinessServer may be giving good response while the application itself
// may be unavailable.
// TODO:
// rest service to expose an endpoint for doing the checks.
// gRPC service to expose an endpoint for doing liveness check (status.Check).
func LivenessReadinessServer(logger *logrus.Entry, readinessCheck Check) *Probe {
	return &Probe{
		Port:           8888,
		logger:         logger,
		liveness:       UNAVAILABLE,
		readinessCheck: readinessCheck,
	}
}

// Start registers HTTP handlers and start server.
func (s *Probe) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/liveness", s.Liveness)
	mux.HandleFunc("/readiness", s.Readiness)

	addr := fmt.Sprintf(":%d", 8888)
	if s.Port != 0 {
		addr = fmt.Sprintf(":%d", s.Port)
	}
	srv := &http.Server{Addr: addr, Handler: mux}
	s.logger.Infof("Starting liveness/readiness probe on %s", addr)

	s.wg.Add(1)
	s.lock.Lock()
	s.server = srv
	s.liveness = RUNNING
	s.lock.Unlock()
	s.wg.Done()

	s.logger.Infof("Liveness status: %d", s.liveness)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.liveness = UNAVAILABLE
		s.logger.Fatalf("liveness/readiness server error listening on port %d: %+v", s.Port, err)
	}
}

// Stop closes the listener used for running the server.
func (s *Probe) Stop() error {
	s.wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// SetLiveness sets liveness of application to be reported on Probes.
func (s *Probe) SetLiveness(status Status) {
	s.lock.Lock()
	s.liveness = status
	s.lock.Unlock()
}

// Liveness endpoint for Kubernetes readiness check and checking service status.
func (s *Probe) Liveness(w http.ResponseWriter, r *http.Request) {
	s.lock.RLock()
	code := int(s.liveness)
	s.lock.RUnlock()

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"status": "%s"}`, http.StatusText(code))
}

// Readiness endpoint for Kubernetes readiness check and checking service status.
func (s *Probe) Readiness(w http.ResponseWriter, r *http.Request) {
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
