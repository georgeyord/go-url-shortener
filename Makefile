GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

clean:
	@rm -r ./bin

deps:
	@go mod download

build-helloworld-cmd:
	@go build -o ./bin/helloworld ./cmd/helloworld/main.go

run-helloworld-cmd:
	@go run ./cmd/helloworld/main.go

test:
	@$(GO_EXECUTABLES)/gotest -v github.com/georgeyord/go-scrumpoker-api/cmd/helloworld

test_deps:
	@go get -u github.com/rakyll/gotest

.PHONY: deps install-cli run-cli test test_deps
