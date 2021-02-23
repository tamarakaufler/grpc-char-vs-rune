package storage

import (
	red "github.com/go-redis/redis/v8"

	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
)

// Redis provides redis client.
type Redis struct {
	Client *red.Client
}

var _ Storage = (*Redis)(nil)

// New provides redis client. Redis customization is done through env variables.
func New(cfg conf.Configuration) *Redis {
	return &Redis{
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
	}
}

// StoreCharToRune ...
func (r *Redis) StoreCharToRune(charToRune string) error {
	return nil
}

// RetrieveCharToRuneConversion ...
func (r *Redis) RetrieveCharToRuneConversion(charToRune string) ([]rune, error) {
	return nil, nil
}

// StoreRuneToChar ...
func (r *Redis) StoreRuneToChar(runeToChar string) error {
	return nil
}

// RetrieveRuneToCharToConversion ...
func (r *Redis) RetrieveRuneToCharToConversion(runeToChar []rune) (string, error) {
	return "", nil
}
