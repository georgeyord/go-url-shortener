GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

clean:
	@rm -rf ./bin

prepare: prepare_paths deps test-deps

prepare-paths: deps test-deps
	mkdir -p ./data
	mkdir -p ./log

deps:
	@go mod download

build-url-shortener-web:
	@cd ./api
	go build -o ../bin/url-shortener-web .

docker-build-url-shortener-web:
	docker-compose build url-shortener-web

run-url-shortener-web:
	@cd ./api
	go run .

build-url-shortener-cli:
	@cd ./cmd
	go build -o ./bin/url-shortener-cli .

run-url-shortener-cli:
	@cd ./cmd
	go run .

docker-build-url-shortener-cli:
	docker-compose build url-shortener-cli

test: test-api test-cli
	@cd ./pkg
	$(GO_EXECUTABLES)/gotest -v ./...

test-pkg:
	@cd ./pkg
	$(GO_EXECUTABLES)/gotest -v ./...

test-api: test-pkg
	@cd ./api
	$(GO_EXECUTABLES)/gotest -v ./...

test-cli: test-pkg
	@cd ./cmd
	$(GO_EXECUTABLES)/gotest -v ./...

test-deps:
	@go get -u github.com/rakyll/gotest
	@go get -u github.com/stretchr/testify

.PHONY: clean deps build-url-shortener-web docker-build-url-shortener-web run-url-shortener-web build-url-shortener-cli docker-build-url-shortener-cli run-url-shortener-cli test test-pkg test-api test-cli test-deps prepare prepare-paths
