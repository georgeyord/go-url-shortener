GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

clean:
	@rm -r ./bin

deps:
	@go mod download

build-scrumpoker-api:
	go build -o ./bin/scrumpoker-api ./api

run-scrumpoker-api:
	go run ./api

build-helloworld-cmd:
	go build -o ./bin/helloworld ./cmd/helloworld/main.go

run-helloworld-cmd:
	go run ./cmd/helloworld/main.go

test:
	@$(GO_EXECUTABLES)/gotest -v ./...

test_deps:
	@go get -u github.com/rakyll/gotest

.PHONY: clean deps build-scrumpoker-api run-scrumpoker-api build-helloworld-cmd run-helloworld-cmd test test_deps
