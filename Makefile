all: test
TESTS := expr unrecognised

.PHONY: test export release fmt example

fmt:
	go fmt ./...

test:
	go get -v ./... && go test -v ./...

example:
	cd example && go get -v && go run 808.go
