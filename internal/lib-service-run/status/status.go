package status

import (
	"net/http"
)

type Status int

const (
	UNAVAILABLE Status = http.StatusServiceUnavailable
	RUNNING     Status = http.StatusOK
	SHUTDOWN    Status = http.StatusGone
)
