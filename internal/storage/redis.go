package storage

import (
	//"github.com/gomodule/redigo/redis"

	red "github.com/go-redis/redis/v8"

	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
)

// Redis provides redis client.
type Redis struct {
	//connectionPool *redis.Pool
	Client *red.Client
}

var _ Storage = (*Redis)(nil)

// New ...
func New(conf conf.Configuration) *Redis {
	return &Redis{
		Client: red.NewClient(&red.Options{
			Addr:         conf.Address,
			Password:     conf.Password,
			PoolSize:     conf.PoolSize,
			PoolTimeout:  conf.PoolTimeout,
			IdleTimeout:  conf.IdleTimeout,
			ReadTimeout:  conf.ReadTimeout,
			WriteTimeout: conf.WriteTimeout,
			MinIdleConns: conf.MinIdleConns,
			MaxRetries:   conf.MaxRetries,
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
