BIN := "./bin"
CALENDAR_IMG="calendar:develop"
SCHEDULER_IMG="scheduler:develop"
SENDER_IMG="sender:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN)/calendar -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o $(BIN)/scheduler -ldflags "$(LDFLAGS)" ./cmd/scheduler
	go build -v -o $(BIN)/sender -ldflags "$(LDFLAGS)" ./cmd/sender

run: build
	$(BIN)/calendar -config ./configs/config_calendar.yaml &
	$(BIN)/scheduler -config ./configs/config_scheduler.yaml &
	$(BIN)/sender -config ./configs/config_sender.yaml &

build-img:
	docker build --build-arg=LDFLAGS="$(LDFLAGS)" -t $(CALENDAR_IMG) -f build/calendar/Dockerfile .
	docker build --build-arg=LDFLAGS="$(LDFLAGS)" -t $(SCHEDULER_IMG) -f build/scheduler/Dockerfile .
	docker build --build-arg=LDFLAGS="$(LDFLAGS)" -t $(SENDER_IMG) -f build/sender/Dockerfile .

run-img: build-img
	docker run $(CALENDAR_IMG)

version: build
	$(BIN)/calendar version
	$(BIN)/scheduler version
	$(BIN)/sender version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.37.0

lint: install-lint-deps
	golangci-lint run ./...

openapi_http:
	oapi-codegen -generate types -o internal/server/http/gen/openapi_types.go -package gen api/http/calendar.yml
	oapi-codegen -generate chi-server -o internal/server/http/gen/openapi_api.go -package gen api/http/calendar.yml
	oapi-codegen -generate client -o internal/server/http/gen/openapi_client.go -package gen api/http/calendar.yml

grpc_proto:
	protoc \
		--proto_path=api/grpc api/grpc/calendar.proto \
		--go_out=internal/server/grpc/gen --go_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		--go-grpc_out=internal/server/grpc/gen --go-grpc_opt=paths=source_relative \


generate: openapi_http grpc_proto

down:
	-docker-compose down

up:
	-docker-compose up -d

integration-tests: up
	sh ./sql/wait_pg.sh
	-go test ./internal/server/e2e -tags e2e
	-docker-compose down

.PHONY: build run build-img run-img version test lint up down integration-tests openapi_http grpc_proto generate