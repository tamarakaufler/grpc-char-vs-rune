package storage

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	red "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	libLogger "github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/logger"
)

// charToRune is a helper type used for storing data in redis.
type charToRune struct {
	List []uint32 `json:"list,omitempty"`
}

// Redis provides redis client.
type Redis struct {
	Logger *logrus.Entry
	Client *red.Client
	TTL    time.Duration
}

var _ Storage = (*Redis)(nil)

// New provides redis client. Redis customization is done through env variables.
func New(cfg conf.Configuration) *Redis {
	l := libLogger.New().WithFields(
		logrus.Fields{
			"storage": "redis",
		},
	)

	return &Redis{
		Logger: l,
		Client: red.NewClient(&red.Options{
			Addr:         cfg.Address,
			Password:     cfg.Password,
			PoolSize:     cfg.PoolSize,
			PoolTimeout:  cfg.PoolTimeout,
			IdleTimeout:  cfg.IdleTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			MinIdleConns: cfg.MinIdleConns,
			MaxRetries:   cfg.MaxRetries,
			DB:           0,
		}),
		TTL: cfg.TTL,
	}
}

// StoreCharToRune ...
func (r *Redis) StoreCharToRune(ctx context.Context, key string, uis []uint32) error {
	chtr := charToRune{
		List: uis,
	}
	v, err := json.Marshal(chtr)
	if err != nil {
		return err
	}

	err = r.Client.Set(ctx, key, v, r.TTL).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetCharToRune ...
func (r *Redis) GetCharToRune(ctx context.Context, key string) ([]uint32, error) {
	v, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	chtr := charToRune{}
	err = json.Unmarshal([]byte(v), &chtr)
	if err != nil {
		return nil, err
	}

	return chtr.List, nil
}

// StoreRuneToChar ...
func (r *Redis) StoreRuneToChar(ctx context.Context, rs []byte, s string) error {
	key := base64.StdEncoding.EncodeToString(rs)
	err := r.Client.Set(ctx, key, s, r.TTL).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetRuneToChar ...
func (r *Redis) GetRuneToChar(ctx context.Context, rs []byte) (string, error) {
	key := base64.StdEncoding.EncodeToString(rs)
	v, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return v, nil
}
