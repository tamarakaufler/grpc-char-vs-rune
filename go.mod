module github.com/tamarakaufler/grpc-char-vs-rune

go 1.16

replace github.com/tamarakaufler/grpc-char-vs-rune/client => ./client

require (
	github.com/caarlos0/env/v6 v6.5.0
	github.com/go-redis/redis/v8 v8.6.0
	github.com/golang/protobuf v1.4.3
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.0
	github.com/tamarakaufler/grpc-char-vs-rune/client v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.35.0
)
