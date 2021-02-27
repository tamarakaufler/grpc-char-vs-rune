package storage

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	red "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	libLogger "github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/logger"
)

// const (
// 	tableToRune = "charToRune"
// 	tableToChar = "runeToChar"
// )

var ttl = 3600 * time.Second

// type Uint32 []uint32

// type CharToRune struct {
// 	Mapping map[string]string `json:"mapping,omitempty"`
// }

type CharToRune struct {
	Mapping []uint32 `json:"mapping,omitempty"`
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
		TTL: ttl,
	}
}

// StoreCharToRune ...
func (r *Redis) StoreCharToRune(ctx context.Context, s string, uis []uint32) error {
	ss := []string{}
	for _, el := range uis {
		ssel := strconv.FormatUint(uint64(el), 10)
		ss = append(ss, ssel)
	}

	err := r.Client.RPush(ctx, s, ss).Err()
	if err != nil {
		return err
	}
	_, err = r.Client.Expire(ctx, s, r.TTL).Result()
	if err != nil {
		r.Logger.Warnf("cannot set expiry on key %s", s)
	}

	return nil
}

// GetCharToRune ...
func (r *Redis) GetCharToRune(ctx context.Context, key string) ([]uint32, error) {
	uis := []uint32{}

	for range key {
		el, err := r.Client.LPop(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		v, err := strconv.Atoi(el)
		if err != nil {
			return nil, err
		}
		uis = append(uis, uint32(v))
	}

	// err := r.StoreCharToRune(ctx, key, uis)
	// if err != nil {
	// 	r.Logger.Warnf("cound not store key %s", key)
	// }

	return uis, nil
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
