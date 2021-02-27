module github.com/tamarakaufler/grpc-char-vs-rune

go 1.16

replace github.com/tamarakaufler/grpc-char-vs-rune/client => ./client

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/caarlos0/env/v6 v6.5.0
	github.com/containerd/continuity v0.0.0-20210208174643-50096c924a4e // indirect
	github.com/go-redis/redis/v8 v8.6.0
	github.com/golang/protobuf v1.4.3
	github.com/lib/pq v1.9.0 // indirect
	github.com/ory/dockertest/v3 v3.6.3
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tamarakaufler/grpc-char-vs-rune/client v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.35.0
)
