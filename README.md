## grpc-char-vs-rune

# Synopsis

gRPC service, written in Go, to provide convertions between characters/strings and runes/bytes.

It will be used with other gRPC and REST services, running locally using docker-compose
or running in a Kubernetes cluster.

# Implementation

## Used technology

- Go
- gRPC
- redis for caching
- golangci-lint for code quality
- protobufs generated with prototool
- mocks generated with counterfeiter
- docker, docker-compose

## Details

### protobufs

Generated using the Uber's prototool.

a) proto description (char-vs-rune.proto)

    option go_package = "github.com/tamarakaufler/grpc-char-vs-rune/client/char_vs_rune";

    the last part _char_vs_rune_ msut not contain dashes

b) If there is an issue regarding the go_package line during the generation, try upgrading the protoc and the Go grpc related plugins.

### client submodule

Used to allow import of only what's needed by the client of this service,ie the protobufs and the related mocks.

NOTE

The protobuf package is defined as char_vs_rune in ./client/char-vs-rune/char-vs-rune.pb.go (see the protobufs section a) above),
but is imported as:

    "github.com/tamarakaufler/grpc-char-vs-rune/client/char-vs-rune"
