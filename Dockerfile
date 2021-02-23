FROM golang:1.15.8-alpine AS builder

ENV VERSION=unknown
ENV GIT_SHA=unknown
ENV NAME=char-vs-rune
ENV TIMESTAMP=unknown
ENV LD_FLAGS="-w -s -X main.Timestamp=${TIMESTAMP} -X main.ServiceVersion=${VERSION}  -X main.GitSHA=${GIT_SHA} -X main.ServiceName=${NAME}"
ENV GOLANGCI_VERSION=v1.36.0

RUN echo ${LD_FLAGS}

# a) for golangci-lint to work either libc-dev needs to be installed or
# CGO_ENABLED=0 must be set
# b) libc-dev installed are prerequisites for running
# go test with the -race flag
RUN apk add --no-cache curl tzdata gcc libc-dev

WORKDIR /app

ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go mod download
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin ${GOLANGCI_VERSION}

RUN golangci-lint run --out-format=line-number
RUN go test -count=1 --race -covermode=atomic -coverprofile=coverage.out ./...

RUN go build -o ./bin/char-vs-rune ./cmd/char-vs-rune/main.go


FROM alpine:3.13
COPY --from=builder /app/bin/char-vs-rune /bin/char-vs-rune
ENTRYPOINT [ "/bin/char-vs-rune" ]
