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

unit-test:
	go test -count=1 -tags unit_tests --race -covermode=atomic -coverprofile=coverage1.out ./...

redis-integration-test:
	go test -count=1 -tags integration_tests --race -covermode=atomic -coverprofile=coverage2.out ./...

handler-integration-test:
	@# change to a different directory only affects this line. Use && to string all make task commands to
	@# propagate the directory change
	cd integration && docker-compose up -d
	go test -count=1 -tags integration_tests --race -covermode=atomic -coverprofile=coverage3.out ./integration/...
	cd integration && docker-compose down --remove-orphans

# runs both unit and integration tests
test:
	go test -count=1 -tags unit_tests,integration_tests --race -covermode=atomic -coverprofile=coverage.out ./internal/...

# CGO_ENABLED=0 is crucial for the test binary to run in the container. If not provided (CGO_ENABLED=1 is the default),
# running the image results in:
#		standard_init_linux.go:219: exec user process caused: no such file or directory
# even though the test binary was copied correctly and present in the image. GO_ENABLED=0 produces a statically linked binary,
# which is needed as the alpine image does not contain all the required libraries out of the box.
# The message "no such file or directory" refers to a missing C library.
acceptance-bin:
	CGO_ENABLED=0 go test -o cmd/bin/acceptancetests -c -v -tags acceptance_tests ./acceptance

acceptance-test:
	go test -count=1 -tags acceptance_tests ./acceptance

acceptance-compose-run: acceptance-bin
	cd acceptance-ci && docker-compose up --force-recreate

acceptance-image: acceptance-bin
	docker build -t grpc-char-vs-rune-test:v1.0.0 -f acceptance-ci/Dockerfile .

acceptance-image-run:
	docker run grpc-char-vs-rune-test:v1.0.0

build: LDFLAGS += -X 'main.Timestamp=$(shell date +%s)'
build: LDFLAGS += -X 'main.ServiceVersion=${VERSION}'
build: LDFLAGS += -X 'main.GitSHA=${GIT_SHA}'
build: LDFLAGS += -X 'main.ServiceName=${FULL_NAME}'
build:
	$(info building binary cmd/bin/$(NAME) with flags $(LDFLAGS))
	go build -race -o cmd/bin/$(NAME) -ldflags "$(LDFLAGS)" cmd/char-vs-rune/main.go
	#@gCGO_ENABLED=0 go build -race -o cmd/bin/$(NAME) -ldflags "$(LDFLAGS)" cmd/char-vs-rune/main.go

redis-run:
	docker-compose up -d redis

run:
	cmd/bin/$(NAME)

docker-run:
# 	docker build -t grpc-char-vs-rune:v1.0.0 .
# 	docker run -w /basedir -v $(PWD):/basedir --env AAA=aaa --env BBB=3 grpc-char-vs-rune:v1.0.0

# starts grpc-char-vs-rune, redis, telegraph and influxdb
services-start:
	docker-compose up -d

# stops grpc-char-vs-rune, redis, telegraph and influxdb
services-stop:
	docker-compose stop
	docker-compose down --remove-orphans

cover:
	@LOG_LEVEL=debug TMP_COV=$(shell mktemp); \
	go test -failfast -coverpkg=./... -coverprofile=$$TMP_COV ./... && \
	go tool cover -html=$$TMP_COV && rm $$TMP_COV

cleanup: services-stop

all: deps protoc lint test handler-integration-test acceptance-bin build

.PHONY: deps tools lint protoc mock unit-test redis-integration-test handler-integration-test test acceptance-bin acceptance-test acceptance-run acceptance-image \
cover build run docker-run redis-run services-start services-stop cover cleanup
