include .env 
export

SHELL := /bin/zsh

build:
	go build -o bin/${BINARY_NAME} ./cmd/main.go

clean:
	go clean
	rm -rf ./bin/climabr

test:
	go test -v ./...

test_coverage:
	go test -v ./... -coverprofile=coverage.out
