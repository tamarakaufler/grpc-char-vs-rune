package status

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
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

// LivenessReadinessServer creates LivenessReadiness instance
// NOTE: This is not great as the application will be serving on a different port, so
// the LivenessReadinessServer may be giving good response while the application itself
// may be unavailable.
// TODO:
// A better way is to attach (if not already implemented) a new API/route to the TEST/HRRP server.
func LivenessReadinessServer(logger *logrus.Entry, readinessCheck Check) *LivenessReadiness {
	return &LivenessReadiness{
		Port:           8000,
		logger:         logger,
		liveness:       UNAVAILABLE,
		readinessCheck: readinessCheck,
	}
}

// Start registers HTTP handlers and start server.
func (s *LivenessReadiness) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/liveness", s.LivenessProbe)
	mux.HandleFunc("/readiness", s.ReadinessProbe)

	srv := &http.Server{Addr: fmt.Sprintf(":%d", s.Port), Handler: mux}

	s.wg.Add(1)
	s.lock.Lock()
	s.server = srv
	s.liveness = RUNNING
	s.lock.Unlock()
	s.wg.Done()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Fatalf("liveness/readiness server error listening on port %d: %+v", s.Port, err)
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
