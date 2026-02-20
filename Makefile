build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -v ./...

coverage:
	go test -coverprofile=.coverage.out ./...
	go tool cover -html=.coverage.out

precommit: fmt lint test
