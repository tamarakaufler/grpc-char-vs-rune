## grpc-char-vs-rune

# Synopsis

An examle of a gRPC service, written in Go.
The service provides convertions between strings and runes.

To be used with other gRPC and REST services, running locally using docker-compose
or running in a Kubernetes cluster.

# Implementation

## Used technology

- Go
- gRPC
- protobufs generated with prototool
- mocks generated with counterfeiter
- redis for caching
- golangci-lint for code quality (https://raw.githubusercontent.com/golangci/golangci-lint/master/.golangci.example.yml)
- dockertest for integration testing
- docker, docker-compose
- telegraf and informixdb for metrics

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

### running services

run and stop using docker-compose:
  - make services-start
  - make services-stop


#### redis - in-memory store

## testing

### unit testing

### integration testing

redis methods for storing and retrieval are tested with the help of dockertest (github.com/ory/dockertest/v3). The dockertest module provides a redis docker container without the need for using docker-compose. These tests are strictly speaking unit tests, but as they require docker, they are tagged as intergration_tests.

###

## Metrics

Metrics sent using the StatsD protocol to Telegraf, the StatsD aware service that stores the data in a timne-series database influxdb database.

The required services, telegraph and influxdb, are started using _make start-services_ task, that also starts the grpc-char-to-rune service.

### TODO

Send the metrics

#### telegraph - StatsD metrics server

NOTE

While _docker pull telegraf_ pulled the image (_docker.io/library/telegraf:latest_), using _telegraph.latest_ as the telegraph image in docker-compose.yml ended with:
    ERROR: pull access denied for telegraph, repository does not exist or may require 'docker login': denied: requested access to the resource is denied

The solution was to use the full name of the image:
    _docker.io/library/telegraf:latest_

**Telegraf configuration**

Needs to contain at least one _inputs_ configurations and one _outputs_ configuration.


## Links

https://www.influxdata.com/blog/running-influxdb-2-0-and-telegraf-using-docker/
https://github.com/influxdata/telegraf/blob/master/etc/telegraf.conf