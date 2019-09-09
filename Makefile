PATH_SRC=./src

deps:
	go get -v $(PATH_SRC)

run: deps
	go run $(PATH_SRC)/*

test: deps
	go test -v $(PATH_SRC)

.PHONY: deps run test
