package storage

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Options ...
type Options struct {
	Logger *logrus.Entry
	TTL    time.Duration
}

// Option ...
type Option func(*Options)

// WithLogger sets custom logger.
func WithLogger(l *logrus.Entry) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// WithTTL sets custom TTL for key expiry.
func WithTTL(ttl time.Duration) Option {
	return func(o *Options) {
		o.TTL = ttl
	}
}
