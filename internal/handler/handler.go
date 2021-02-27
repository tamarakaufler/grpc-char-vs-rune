package handler

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
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
func (h *Handler) ToChar(ctx context.Context, req *proto.ToCharRequest) (*proto.ToCharResponse, error) {
	if req == nil || req.GetRunes() == nil {
		return nil, errors.New("no runes to convert to string")
	}
	uis := req.GetRunes()

	// get cached
	rs := []rune{}
	for _, ui := range uis {
		rs = append(rs, rune(ui))
	}

	b := []byte(string(rs))
	s, err := h.Clients.Redis.GetRuneToChar(ctx, b)
	if err != nil {
		h.Logger.Warnf("Cannot get key for %v from redis", rs)
	} else {
		return &proto.ToCharResponse{
			To: s,
		}, nil
	}

	// cache the conversion
	err = h.Clients.Redis.StoreRuneToChar(ctx, b, s)
	if err != nil {
		h.Logger.Warnf("Cannot store key for %v in redis", uis)
	}

	return &proto.ToCharResponse{
		To: s,
	}, nil
}

// ToRune ... the service must satisfy the CharVsRuneServer interface.
func (h *Handler) ToRune(ctx context.Context, req *proto.ToRuneRequest) (*proto.ToRuneResponse, error) {
	if req == nil || req.GetFrom() == "" {
		return nil, errors.New("no string to convert to runes")
	}
	s := req.GetFrom()

	// get cached
	v, err := h.Clients.Redis.GetCharToRune(ctx, s)
	if err != nil {
		h.Logger.Warnf("Cannot get key %s from redis", s)
	} else {
		rs := ConvertToRuneResponse(s, v)
		return rs, nil
	}

	r, m := ConvertToRune(s)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot convert string %s to runes", s)
	}

	// cache the conversion
	err = h.Clients.Redis.StoreCharToRune(ctx, s, v)
	if err != nil {
		h.Logger.Warnf("Cannot store key %s in redis", s)
	}

	return &proto.ToRuneResponse{
		Runes:   r,
		Mapping: m,
	}, nil
}

// ConvertToRune converts provided string to runes.
func ConvertToRune(s string) ([]uint32, map[string]uint32) {
	m := make(map[string]uint32, len(s))

	uis := []uint32{}
	for _, r := range s {
		ui := uint32(r)
		uis = append(uis, ui)
		m[string(r)] = ui
	}
	return uis, m
}

// ConvertToChar converts provided runes to string.
func ConvertToChar(rs []rune) string {
	return string(rs)
}

// helper functions/methods

// ConvertToRuneResponse ...
func ConvertToRuneResponse(s string, v []uint32) *proto.ToRuneResponse {
	m := make(map[string]uint32, len(s))
	for i, ch := range s {
		ch := fmt.Sprint(ch)
		m[ch] = v[i]
	}

	return &proto.ToRuneResponse{
		Runes:   v,
		Mapping: m,
	}
}
