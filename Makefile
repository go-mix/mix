all: test
TESTS := expr unrecognised

.PHONY: test export profile release fmt example

fmt:
	go fmt ./...

test:
	go get -v ./... && go test -v ./...

example:
	cd example && go get -v && go run 808.go

profile:
	cd example && go get -v && go run 808.go --profile cpu
