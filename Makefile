unit-test:
	go test -v -count=1 -race -gcflags=-l -timeout=30s ./test

integration-test:
	go test ./... -v -race -tags=integration

run-go:
	go run ./cmd/main.go

.PHONY: build

build:
	go build -o rotator ./cmd/main.go

generate:
	go generate cmd/main.go


go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...