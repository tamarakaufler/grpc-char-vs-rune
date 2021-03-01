package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

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

	mux *sync.Mutex
}

var _ proto.CharVsRuneServer = (*Handler)(nil)

// New contructs a new Handler instance for access to the service APIs.
func New(l *logrus.Entry, cfg configuration.Configuration) *Handler {
	clients := service.ClientsBuilder(cfg)
	return &Handler{
		Logger:  l,
		Clients: clients,

		mux: &sync.Mutex{},
	}
}

// ToChar ... the service must satisfy the CharVsRuneServer interface.
// This is a very contrived example. No need to cache.
func (h *Handler) ToChar(ctx context.Context, req *proto.ToCharRequest) (*proto.ToCharResponse, error) {
	if req == nil || req.GetRunes() == nil {
		return nil, errors.New("no runes to convert to string")
	}
	uis := req.GetRunes()

	// get cached
	key := Uint32ListToBase64(uis)

	fmt.Printf("---> ToChar: key %s", key)

	s, err := h.Clients.Redis.GetRuneToChar(ctx, key)
	if err != nil {
		if err.Error() != "redis: nil" {
			h.Logger.Warnf("Cannot get key for %v from redis", uis)
		}
	} else {
		return &proto.ToCharResponse{
			To: s,
		}, nil
	}

	s = ConvertToChar(uis)
	fmt.Printf("\t\t### %v\n", s)

	// cache the conversion
	h.mux.Lock()
	err = h.Clients.Redis.StoreRuneToChar(ctx, key, s)
	h.mux.Unlock()
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
		if err.Error() != "redis: nil" {
			h.Logger.Warnf("problem with redis: %s", err)
		}
	} else {
		rs := ConvertToRuneResponse(v)
		return rs, nil
	}

	r, m := ConvertToRune(s)

	// cache the conversion
	h.mux.Lock()
	err = h.Clients.Redis.StoreCharToRune(ctx, s, r)
	h.mux.Unlock()
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
func ConvertToChar(uis []uint32) string {
	rs := []rune{}
	for _, ui := range uis {
		rs = append(rs, rune(ui))
	}

	fmt.Printf("\t\t### %v\n", string(rs))

	return string(rs)
}

// helper functions/methods

// ConvertToRuneResponse ...
// TODO: unit test
// The index i when ranging over a string, may not be sequential as
// the looping goes through runes, which may be more than 1 byte,
// and the index refers to the byte position.
func ConvertToRuneResponse(v []uint32) *proto.ToRuneResponse {
	m := make(map[string]uint32, len(v))
	for i := range v {
		ch := fmt.Sprintf("%q", v[i])
		m[ch] = v[i]
	}

	return &proto.ToRuneResponse{
		Runes:   v,
		Mapping: m,
	}
}

// Uint32ListToBase64 converts a []uint32 into a base64 string.
func Uint32ListToBase64(uis []uint32) string {
	var sb strings.Builder
	for _, ui := range uis {
		sb.WriteString(string(rune(ui)))
	}
	return base64.StdEncoding.EncodeToString([]byte(sb.String()))
}
