package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"

	proto "github.com/tamarakaufler/grpc-char-vs-rune/client/char-vs-rune"
	conf "github.com/tamarakaufler/grpc-char-vs-rune/internal/configuration"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/handler"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/grpc/server"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/lib-service-run/logger"
)

// gRPC service identifiers.
var (
	ServiceVersion string
	Timestamp      string
	ServiceName    string
	GitSHA         string
)

func main() {
	fmt.Println("*** grpc-char-vs-rune service greets the world")

	lgr := logger.New().WithFields(logrus.Fields{
		"logger":    "main",
		"gitSHA":    GitSHA,
		"timestamp": Timestamp,
	})
	lgr.Infoln("Hello world")

	os.Setenv("REDIS_ADDRESS", "localhost:6379") // for testing only

	cfg, err := conf.Load()
	if err != nil {
		logrus.Panicf("cannot start grpc-char-vs-rune: %v", err)
	}

	// Setting up the gRPC server.
	s, err := server.New(
		server.WithName(ServiceName),
		server.WithVersion(ServiceVersion),
	)
	if err != nil {
		lgr.Panicf("Cannot create grpc-char-vs-rune service")
	}

	// Setting up the handler providing access to the service APIs.
	// Implements the CharVsRuneServer interface.
	h := handler.New(lgr, cfg)

	// Testing access to redis. Remove after dev done.
	pong, err := h.Clients.Redis.Client.Ping(context.Background()).Result()
	if err != nil {
		lgr.Errorf("Cannot Ping: %s", err.Error())
	} else {
		lgr.Infof("Second Pong: %s", pong)
	}

	proto.RegisterCharVsRuneServer(s.Server(), h)

	if err := s.Start(); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
