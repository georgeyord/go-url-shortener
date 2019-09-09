PATH_SRC=./src

deps:
	@go get -v $(PATH_SRC)

run: deps
	@go run $(PATH_SRC)/*

test: deps
	@which ssgotest 2>&1 > /dev/null && \
		gotest -v $(PATH_SRC) || \
		( \
			echo "'gotest' pkg is missing, falling back to 'go test' (install using 'go get -u github.com/rakyll/gotest'go get -u github.com/rakyll/gotest command)." && \
			go test -v $(PATH_SRC) \
		)

.PHONY: deps run test
