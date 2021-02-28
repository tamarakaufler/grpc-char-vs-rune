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

func TestHandler(t *testing.T) {
	h := setup(t)
	defer t.Cleanup(cleanup)

	fmt.Println("Hello friend")

	pong, err := h.Clients.Redis.Client.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.NotNil(t, pong)

	require.NotEqual(t, "Hello", "World")

	t.Run("ToRune", func(t *testing.T) {
		ctx := context.Background()
		s := "Hello 日本"
		expected := storage.RuneToUint32(s)

		pong, err := h.Clients.Redis.Client.Ping(ctx).Result()
		require.NoError(t, err)
		require.NotNil(t, pong)

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
