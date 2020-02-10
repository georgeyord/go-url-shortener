GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

clean:
	@rm -rf ./bin

deps:
	@go mod download

build-url-shortener-web:
	go build -o ./bin/url-shortener-web ./api

docker-build-url-shortener-web:
	docker-compose build url-shortener-web

run-url-shortener-web:
	go run ./api

build-url-shortener-cli:
	go build -o ./bin/url-shortener-cli ./cmd/url-shortener/main.go

run-url-shortener-cli:
	go run ./cmd/url-shortener/main.go

docker-build-url-shortener-cli:
	docker-compose build url-shortener-cli

test:
	$(GO_EXECUTABLES)/gotest -v ./...

test_deps:
	@go get -u github.com/rakyll/gotest
	@go get -u github.com/stretchr/testify

.PHONY: clean deps build-url-shortener-web docker-build-url-shortener-web run-url-shortener-web build-url-shortener-cli docker-build-url-shortener-cli run-url-shortener-cli test test_deps
