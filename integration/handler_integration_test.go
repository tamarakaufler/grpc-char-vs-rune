// +build integration_tests

package handler_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/handler"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/logger"
)

func TestHandler(t *testing.T) {
	h := setup(t)
	h.Logger.Info("Hello friend")

	pong, err := h.Clients.Redis.Client.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.NotNil(t, pong)

	require.NotEqual(t, "Hello", "World")
	defer t.Cleanup(cleanup)

}

func setup(t *testing.T) *handler.Handler {
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
