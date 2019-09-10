GO_EXECUTABLES=$(shell go env GOPATH)/bin
PATH_SRC=./src

deps:
	@go get -v $(PATH_SRC)

install-cli: deps
	@go install $(PATH_SRC)/helloworld-cli.go

run-cli: deps
	@go run $(PATH_SRC)/helloworld-cli.go

test: deps
	@which "$(GO_EXECUTABLES)/gotest" 2>&1 > /dev/null && \
		"$(GO_EXECUTABLES)/gotest" -v $(PATH_SRC) || \
		( \
			echo "'$(GO_EXECUTABLES)/gotest' binary is missing, falling back to 'go test' (install using 'make test_deps')." && \
			go test -v $(PATH_SRC) \
		)

test_deps:
	@go get -u github.com/rakyll/gotest

.PHONY: deps install-cli run-cli test test_deps
