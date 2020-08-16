build:
	go build -o rotator ./cmd/main.go

run:
	go run ./cmd/main.go -config ./config/main.json

test:
	go test -v -count=1 -race -gcflags=-l -timeout=30s ./test

go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...

install:
	go mod download

generate:
	go generate ./...