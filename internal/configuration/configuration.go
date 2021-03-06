package configuration

import (
	"time"

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
	Address     string        `env:"REDIS_ADDRESS,required"`
	Password    string        `env:"REDIS_PASSWORD" envDefault:""`
	PoolSize    int           `env:"POOL_SIZE" envDefault:"100"`
	PoolTimeout time.Duration `env:"POOL_TIMEOUT"  envDefault:"5s"`
	// MaxActive    int           `env:"REDIS_MAX_ACTIVE" envDefault:"100"`
	// MaxIdle      int           `env:"REDIS_MAX_IDLE" envDefault:"3"`
	IdleTimeout  time.Duration `env:"REDIS_IDLE_TIMEOUT" envDefault:"5s"` // should be lower than the server timeout
	ReadTimeout  time.Duration `env:"REDIS_READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout time.Duration `env:"REDIS_WRITE_TIMEOUT" envDefault:"5s"`
	MinIdleConns int           `env:"REDIS_MIN_RETRIES" envDefault:"3"`
	MaxRetries   int           `env:"REDIS_MAX_RETRIES" envDefault:"10"`
	DB           int           `env:"REDIS_DB" envDefault:"0"`
	TTL          time.Duration `env:"REDIS_TTL" envDefault:"1s"`
}

// Metrics represents the configuration for metrics.
type Metrics struct {
	StatsdAddress string `env:"STATSD_ADDRESS" envDefault:"localhost:8125"` // telegraph
}

// Load provides new configuration, using env variable overrides if they are set up.
func Load() (Configuration, error) {
	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return Configuration{}, err
	}
	return *cfg, nil
}
