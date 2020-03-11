GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

clean:
	@rm -rf ./bin

prepare: prepare_paths deps test_deps

prepare_paths: deps test_deps
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

test:
	$(GO_EXECUTABLES)/gotest -v ./...

test_deps:
	@go get -u github.com/rakyll/gotest
	@go get -u github.com/stretchr/testify

.PHONY: clean deps build-url-shortener-web docker-build-url-shortener-web run-url-shortener-web build-url-shortener-cli docker-build-url-shortener-cli run-url-shortener-cli test test_deps prepare prepare_paths
