package handler

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	proto "github.com/tamarakaufler/grpc-char-vs-rune/client/char-vs-rune"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/service"
)

// Handler represents an instance of the CharVsRuneServer.
type Handler struct {
	Logger  *logrus.Entry
	Clients *service.Clients
}

var _ proto.CharVsRuneServer = (*Handler)(nil)

// New contructs a new Handler instance for access to the service APIs.
func New(l *logrus.Entry, cfg configuration.Configuration) *Handler {
	clients := service.ClientsBuilder(cfg)
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
	if req == nil || req.GetFrom() == "" {
		return nil, errors.New("no string to convert to runes")
	}
	from := req.GetFrom()

	to := []uint32{}
	mapping := make(map[string]uint32)
	for _, ch := range from {
		ui := uint32(ch)
		to = append(to, ui)
		mapping[string(ch)] = ui
	}

	return &proto.ToRuneResponse{
		InRunes: to,
		Mapping: mapping,
	}, nil
}

// ConvertToRune converts provided string to runes.
func ConvertToRune(s string) ([]uint32, map[string]uint32) {
	r := []uint32{}
	m := make(map[string]uint32)

	for _, ch := range s {
		ui := uint32(ch)
		r = append(r, ui)
		m[string(ch)] = ui
	}
	return r, m
}
