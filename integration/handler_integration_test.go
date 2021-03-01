// +build integration_tests

package handler_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	proto "github.com/tamarakaufler/grpc-char-vs-rune/client/char-vs-rune"
	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/handler"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/logger"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/storage"
)

func TestToRune(t *testing.T) {
	h := setup(t)
	defer t.Cleanup(cleanup)

	fmt.Println("Hello friend")

	t.Run("Test_parallel_1", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		s := "Hello 日本"
		expected := storage.RuneToUint32(s)

		// err = h.Clients.Redis.StoreCharToRune(ctx, s, expected)
		// require.NoError(t, err)

		// v, err := h.Clients.Redis.GetCharToRune(ctx, s)
		// require.NoError(t, err)
		// require.EqualValues(t, expected, v)

		// string not cached
		req := &proto.ToRuneRequest{From: s}
		res, err := h.ToRune(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetRunes())

		// string cached
		res, err = h.ToRune(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetRunes())
	})

	t.Run("Test_parallel_2", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		s := "Hello 日本"
		expected := storage.RuneToUint32(s)

		// string not cached
		req := &proto.ToRuneRequest{From: s}
		res, err := h.ToRune(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetRunes())

		// string cached
		res, err = h.ToRune(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetRunes())
	})
}

func TestToChar(t *testing.T) {
	h := setup(t)
	defer t.Cleanup(cleanup)

	fmt.Println("Hello foe")

	t.Run("Test_parallel_1", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		rs := []uint32{72, 101, 108, 108, 111, 32, 26085, 26412}
		expected := "Hello 日本"

		// string not cached
		req := &proto.ToCharRequest{Runes: rs}

		res, err := h.ToChar(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetTo())

		// string cached
		res, err = h.ToChar(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetTo())
	})

	t.Run("Test_parallel_2", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		rs := []uint32{72, 101, 108, 108, 111, 32, 26085, 26412}
		expected := "Hello 日本"

		// string not cached
		req := &proto.ToCharRequest{Runes: rs}

		res, err := h.ToChar(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetTo())

		// string cached
		res, err = h.ToChar(ctx, req)
		require.NoError(t, err)
		require.EqualValues(t, expected, res.GetTo())
	})
}

func setup(t *testing.T) *handler.Handler {
	t.Helper()

	l := logger.New().WithFields(logrus.Fields{
		"logger": "integration-handler-test",
	})

	os.Setenv("REDIS_ADDRESS", "localhost:9736")
	cfg, err := conf.Load()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	return handler.New(l, cfg)
}

func cleanup() {
	fmt.Println("Bye friend")
}
