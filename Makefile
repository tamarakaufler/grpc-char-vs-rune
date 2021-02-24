VERSION  ?= unknown
LDFLAGS  := -w -s
NAME     := char-vs-rune
FULL_NAME := grpc-${NAME}
GIT_SHA  ?= $(shell git rev-parse --short HEAD)
GOLANGCI_VERSION = v1.36.0
GOLANGCI := $(shell which golangci-lint 2>/dev/null)

deps:
	@go mod download
	@go mod tidy

tools:
	$(info Install tools)
ifeq ($(GOLANGCI),)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin ${GOLANGCI_VERSION}
endif

lint: tools
	golangci-lint -v run --out-format=line-number

protoc:
	$(info Compile protocol buffers)
	@$(MAKE) -C proto protogen

mock:
	$(MAKE) -C ./client mocks
	go generate ./..

test:
	go test -count=1 --race -covermode=atomic -coverprofile=coverage.out ./...

acceptance-bin:
	CGO_ENABLED=0 go test -o bin/acceptancetests -c -v -tags acceptance_test \
	./acceptance_test ./acceptance


build: LDFLAGS += -X 'main.Timestamp=$(shell date +%s)'
build: LDFLAGS += -X 'main.ServiceVersion=${VERSION}'
build: LDFLAGS += -X 'main.GitSHA=${GIT_SHA}'
build: LDFLAGS += -X 'main.ServiceName=${FULL_NAME}'
build:
	$(info building binary cmd/bin/$(NAME) with flags $(LDFLAGS))
	@go build -race -o cmd/bin/$(NAME) -ldflags "$(LDFLAGS)" cmd/char-vs-rune/main.go

run:
	cmd/bin/$(NAME)

docker-run:
# 	docker build -t grpc-char-vs-rune:v1.0.0 .
# 	docker run -w /basedir -v $(PWD):/basedir --env AAA=aaa --env BBB=3 grpc-char-vs-rune:v1.0.0

redis-run:
	docker-compose up -d redis

# starts this service, redis, telegraph and influxdb
services-start:
	docker-compose up -d

services-stop:
	docker-compose stop
	docker-compose down --remove-orphans

cover:
	@LOG_LEVEL=debug TMP_COV=$(shell mktemp); \
	go test -failfast -coverpkg=./... -coverprofile=$$TMP_COV ./... && \
	go tool cover -html=$$TMP_COV && rm $$TMP_COV

all: deps protoc lint test build

.PHONY: deps tools lint protoc mock test acceptance-bin cover build run docker-run redis-run services-start services-stop
