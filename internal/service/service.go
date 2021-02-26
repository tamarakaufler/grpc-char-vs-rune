package service

import (
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/storage"
)

// Clients to provide access this service needs.
type Clients struct {
	Redis *storage.Redis
}

// ClientsBuilder assembles all clients this service needs to use.
func ClientsBuilder(cfg configuration.Configuration) *Clients {
	redis := storage.New(cfg)
	return &Clients{
		Redis: redis,
	}
}
