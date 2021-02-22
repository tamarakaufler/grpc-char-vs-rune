package configuration_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		envs    map[string]string
		want    *conf.Configuration
		wantErr bool
	}{
		{
			name: "only required set",
			envs: map[string]string{
				"REDIS_ADDRESS": "localhost:6379",
			},
			want: &conf.Configuration{
				LogLevel: `info`,
				Redis: conf.Redis{
					Address:           `localhost:6379`,
					Password:          ``,
					PoolSize:          1,
					PoolTimeout:       time.Duration(5 * time.Second),
					MaxActive:         500,
					MaxIdle:           3,
					IdleTimeout:       time.Duration(5 * time.Second),
					ReadTimeout:       time.Duration(5 * time.Second),
					WriteTimeout:      time.Duration(5 * time.Second),
					ConnectionTimeout: time.Duration(5 * time.Second),
					MaxRetries:        10,
					CacheTTL:          time.Duration(3600 * time.Second),
					DB:                0,
				},
				Metrics: conf.Metrics{
					StatsdAddress: "localhost:8125",
				},
			},
			wantErr: false,
		},
		{
			name: "top level config set",
			envs: map[string]string{
				"REDIS_ADDRESS": "localhost:6379",
			},
			want: &conf.Configuration{
				LogLevel: `info`,
				Redis: conf.Redis{
					Address:           `localhost:6379`,
					Password:          ``,
					PoolSize:          1,
					PoolTimeout:       time.Duration(5 * time.Second),
					MaxActive:         500,
					MaxIdle:           3,
					IdleTimeout:       time.Duration(5 * time.Second),
					ReadTimeout:       time.Duration(5 * time.Second),
					WriteTimeout:      time.Duration(5 * time.Second),
					ConnectionTimeout: time.Duration(5 * time.Second),
					MaxRetries:        10,
					CacheTTL:          time.Duration(3600 * time.Second),
					DB:                0,
				},
				Metrics: conf.Metrics{
					StatsdAddress: "localhost:8125",
				},
			},
			wantErr: false,
		},
		{
			name: "multiple level config set",
			envs: map[string]string{
				"REDIS_ADDRESS":  "localhost:6379",
				"LOG_LEVEL":      "debug",
				"STATSD_ADDRESS": "127.0.0.1:6666",
			},
			want: &conf.Configuration{
				LogLevel: `debug`,
				Redis: conf.Redis{
					Address:           `localhost:6379`,
					Password:          ``,
					PoolSize:          1,
					PoolTimeout:       time.Duration(5 * time.Second),
					MaxActive:         500,
					MaxIdle:           3,
					IdleTimeout:       time.Duration(5 * time.Second),
					ReadTimeout:       time.Duration(5 * time.Second),
					WriteTimeout:      time.Duration(5 * time.Second),
					ConnectionTimeout: time.Duration(5 * time.Second),
					MaxRetries:        10,
					CacheTTL:          time.Duration(3600 * time.Second),
					DB:                0,
				},
				Metrics: conf.Metrics{
					StatsdAddress: "127.0.0.1:6666",
				},
			},
			wantErr: false,
		},
		{
			name: "required config missing",
			envs: map[string]string{
				"LOG_LEVEL":      "debug",
				"STATSD_ADDRESS": "127.0.0.1:6666",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			setEnvs(tt.envs)
			got, err := conf.New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
		unsetEnvs(tt.envs)
	}
}

func setEnvs(es map[string]string) {
	for k, v := range es {
		os.Setenv(k, v)
	}
}

func unsetEnvs(es map[string]string) {
	for k := range es {
		os.Unsetenv(k)
	}
}
