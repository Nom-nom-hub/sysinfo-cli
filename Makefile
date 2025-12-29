.PHONY: build test lint clean fmt

build:
	go build -o bin/sysinfo ./cmd/sysinfo

test:
	go test -v -cover ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

fmt:
	gofmt -s -w .
	go mod tidy

clean:
	rm -rf bin/ dist/ coverage.out coverage.html
	go clean

run: build
	./bin/sysinfo os

run-cpu: build
	./bin/sysinfo cpu

run-memory: build
	./bin/sysinfo memory

run-disk: build
	./bin/sysinfo disk

run-network: build
	./bin/sysinfo network

run-process: build
	./bin/sysinfo process

all: fmt lint test build
