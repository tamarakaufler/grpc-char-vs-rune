package handler

import (
	"context"

	"github.com/sirupsen/logrus"

	proto "github.com/tamarakaufler/grpc-char-vs-rune/client/char-vs-rune"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/storage"
)

// Handler represents an instance of the CharVsRuneServer.
type Handler struct {
	Logger  *logrus.Entry
	Clients *Clients
}

// Clients to provide access this service needs.
type Clients struct {
	Redis *storage.Redis
}

var _ proto.CharVsRuneServer = (*Handler)(nil)

// New contructs a new Handler instance for access to the service APIs.
func New(l *logrus.Entry, cfg configuration.Configuration) *Handler {
	clients := clientsBuilder(cfg)
	return &Handler{
		Logger:  l,
		Clients: clients,
	}
}

// ToChar ... the service must satisfy the CharVsRuneServer interface.
func (s *Handler) ToChar(ctx context.Context, req *proto.ToCharRequest) (*proto.ToCharResponse, error) {
	return &proto.ToCharResponse{}, nil
}

// ToRune ... the service must satisfy the CharVsRuneServer interface.
func (s *Handler) ToRune(ctx context.Context, req *proto.ToRuneRequest) (*proto.ToRuneResponse, error) {
	return &proto.ToRuneResponse{}, nil
}

// clientsBuilder assembles all clients this service needs to use.
func clientsBuilder(cfg configuration.Configuration) *Clients {
	redis := storage.New(cfg)
	return &Clients{
		Redis: redis,
	}
}
