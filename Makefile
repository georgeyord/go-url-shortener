GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

deps:
	@go get -v $(PATH_SRC)

install-cli: deps
	@go install $(PATH_SRC)/helloworld-cli.go

run-cli: deps
	@go run $(PATH_SRC)/helloworld-cli.go

test: deps
	@$(GO_EXECUTABLES)/gotest -v $(PATH_SRC)

test_deps:
	@go get -u github.com/rakyll/gotest

.PHONY: deps install-cli run-cli test test_deps
