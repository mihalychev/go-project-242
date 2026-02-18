build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint:
	golangci-lint run

test:
	go test -v ./...
