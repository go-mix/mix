all: test
TESTS := expr unrecognised

.PHONY: test profile fmt example clean cover

fmt:
	go fmt ./...

test:
	go get -v ./... && go test ./...

example:
	cd example && go get -v && go run 808.go

profile:
	cd example && go get -v && go run 808.go --profile cpu

clean:
	rm *.out bind/*.out

cover: # test coverage
	go test -coverprofile cover.out && go tool cover -html=cover.out
	cd bind && go test -coverprofile cover.out && go tool cover -html=cover.out
