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
			Addr:       "localhost:6379",
			Password:   "", // password set
			DB:         0,  // use default DB
			MaxRetries: 3,  //added after a google suggestion
		}),
	}
}

// // New constructor creates a new Redis client.
// func New(conf *conf.Configuration) *Redis {
// 	pool := &redis.Pool{
// 		MaxActive:   conf.Redis.MaxActive,
// 		MaxIdle:     conf.Redis.MaxIdle,
// 		IdleTimeout: conf.Redis.IdleTimeout,
// 		Dial: func() (redis.Conn, error) {
// 			opts := []redis.DialOption{
// 				redis.DialDatabase(0),
// 				redis.DialReadTimeout(conf.Redis.ReadTimeout),
// 				redis.DialWriteTimeout(conf.Redis.WriteTimeout),
// 				redis.DialConnectTimeout(conf.Redis.ConnectionTimeout),
// 			}
// 			if conf.Redis.Password != "" {
// 				opts = append(opts, redis.DialPassword(conf.Redis.Password))
// 			}
// 			return redis.Dial("tcp", conf.Redis.Address, opts...)
// 		},
// 	}

// 	return &Redis{
// 		connectionPool: pool,
// 	}
// }

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
