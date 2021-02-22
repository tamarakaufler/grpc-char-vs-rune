module github.com/tamarakaufler/grpc-char-vs-rune

go 1.16

replace github.com/tamarakaufler/grpc-char-vs-rune/client => ./client

require (
	github.com/caarlos0/env/v6 v6.5.0
	github.com/go-redis/redis/v8 v8.6.0
	github.com/sirupsen/logrus v1.8.0
)
