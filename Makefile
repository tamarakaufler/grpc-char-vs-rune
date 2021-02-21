VERSION  ?= unknown
LDFLAGS  := -w -s
NAME     := grpc-char-vs-rune
GIT_SHA  ?= $(shell git rev-parse --short HEAD)
GOLANGCI_VERSION = v1.36.0

# GOLANGCI := $(shell which golangci-lint 2>/dev/null)
# ifeq ($(GOLANGCI),)
# 	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin ${GOLANGCI_VERSION}
# endif

deps:
	@go mod download
	@go mod tidy

lint:
	${GOLANGCI} -v run --out-format=line-number

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
build: LDFLAGS += -X 'main.Version=${VERSION}'
build: LDFLAGS += -X 'main.GitSHA=${GIT_SHA}'
build: LDFLAGS += -X 'main.ServiceName=${NAME}'
build:
	$(info building binary cmd/bin/$(NAME) with flags $(LDFLAGS))
	@go build -race -o cmd/bin/$(NAME) -ldflags "$(LDFLAGS)" cmd/watcher-daemon/main.go

run:
	cmd/bin/$(NAME)

docker-run:
	docker build -t watcher-daemon:v1.0.0 .
	docker run -w /basedir -v $(PWD):/basedir --env WATCHER_DAEMON_EXCLUDED=vendor --env WATCHER_DAEMON_FREQUENCY=3 watcher-daemon:v1.0.0

cover:
	@LOG_LEVEL=debug TMP_COV=$(shell mktemp); \
	go test -failfast -coverpkg=./... -coverprofile=$$TMP_COV ./... && \
	go tool cover -html=$$TMP_COV && rm $$TMP_COV

all: deps lint test build

.PHONY: deps lint protoc mock test acceptance-bin cover build run docker-run
