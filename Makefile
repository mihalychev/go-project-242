build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./...

precommit: fmt lint test
