all: test
TESTS := expr unrecognised

.PHONY: test profile fmt demo clean cover

fmt:
	go fmt ./...

test:
	go get -v ./... && go test ./...

demo:
	cd demo && go get -v && go run demo.go

demo.wav:
	cd demo && go get -v && go run demo.go -out wav > output.wav

profile:
	cd demo && go get -v && go run demo.go --profile cpu

clean:
	rm *.out bind/*.out

cover: # test coverage
	go test -coverprofile cover.out && go tool cover -html=cover.out
	cd bind && go test -coverprofile cover.out && go tool cover -html=cover.out
