GO_EXECUTABLES=$(shell go env GOPATH)/bin
TEST_ARGS ?= -v

clean:
	@rm -rf ./bin

prepare: prepare-paths deps test-deps

prepare-paths: deps test-deps
	mkdir -p ./data
	mkdir -p ./log

deps:
	@go mod download

# API
build-url-shortener-api:
	@cd ./api && go build -o ../bin/url-shortener-api .

docker-build-url-shortener-api:
	docker-compose build url-shortener-api

run-url-shortener-api:
	@cd ./api && go run .

test-api: test-pkg
	@cd ./api && $(GO_EXECUTABLES)/gotest $(TEST_ARGS) ./...

# CLI
build-url-shortener-cli:
	@cd ./cli && go build -o ../bin/url-shortener-cli .

run-url-shortener-cli:
	@cd ./cli && go run .

docker-build-url-shortener-cli:
	docker-compose build url-shortener-cli

test-cli: test-pkg
	@cd ./cli && $(GO_EXECUTABLES)/gotest $(TEST_ARGS) ./...

# Worker
build-url-shortener-worker:
	@cd ./worker && go build -o ../bin/url-shortener-worker .

docker-build-url-shortener-worker:
	docker-compose build url-shortener-worker

run-url-shortener-worker:
	@cd ./worker && go run .

test-worker: test-pkg
	@cd ./worker && $(GO_EXECUTABLES)/gotest $(TEST_ARGS) ./...

# To get the test coverage run: make test TEST_ARGS="-cover"
test:
	@$(GO_EXECUTABLES)/gotest $(TEST_ARGS) ./...

test-pkg:
	@cd ./pkg && $(GO_EXECUTABLES)/gotest $(TEST_ARGS) ./...

test-deps:
	@go get -u github.com/rakyll/gotest
	@go get -u github.com/stretchr/testify

.PHONY: clean deps build-url-shortener-api docker-build-url-shortener-api run-url-shortener-api test-api build-url-shortener-worker docker-build-url-shortener-worker run-url-shortener-worker test-worker build-url-shortener-cli docker-build-url-shortener-cli run-url-shortener-cli test-cli test test-pkg test-deps prepare prepare-paths
