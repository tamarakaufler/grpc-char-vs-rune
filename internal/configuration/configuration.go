package configuration

import (
	env "github.com/caarlos0/env/v6"
)

// Configuration holds all the required setup for the service.
type Configuration struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	Redis
	Metrics
}

// Redis holds redis setup.
type Redis struct {
	Address           string `env:"REDIS_ADDRESS,required"`
	Password          string `env:"REDIS_PASSWORD" envDefault:""`
	MaxActive         string `env:"REDIS_MAX_ACTIVE" envDefault:"500"`
	MaxIdle           string `env:"REDIS_MAX_IDLE" envDefault:"3"`
	IdleTimeout       string `env:"REDIS_IDLE_TIMEOUT" envDefault:"5s"` // should be lower than the server timeout
	ReadTimeout       string `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout      string `env:"WRITE_TIMEOUT" envDefault:"5s"`
	ConnectionTimeout string `env:"CONNECTION_TIMEOUT" envDefault:"5s"`
	CacheTTL          string `env:"REDIS_CACHE_TTL" envDefault:"3600s"`

	// MaxIdle     int           `env:"REDIS_MAX_IDLE" envDefault:"3"`
	// IdleTimeout time.Duration `env:"REDIS_IDLE_TIMEOUT" envDefault:"5s"`
	// ReadTimeout       time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	// WriteTimeout      time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`
	// ConnectionTimeout time.Duration `env:"CONNECTION_TIMEOUT" envDefault:"5s"`
	// CacheTTL    time.Duration `env:"REDIS_CACHE_TTL" envDefault:"3600s"`
}

// Metrics represents the configuration for metrics.
type Metrics struct {
	StatsdAddress string `env:"STATSD_ADDRESS" envDefault:"localhost:8125"` // telegraph
}

// New provides new configuration, using env variable overrides if they are set up.
func New() (*Configuration, error) {
	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
