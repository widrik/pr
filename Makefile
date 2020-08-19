unit-test:
	go test ./... -v -race -count 100 -tags=unit

integration-test:
	go test ./... -v -race -tags=integration

test:
	go test ./... -v -race -tags=unit

run:
	go run cmd/main.go

run-go:
	go run cmd/main.go

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...

.PHONY: build test

build:
	CGO_ENABLED=0 go build -o bin/pr cmd/main.go

build-dev:
	go build -o bin/pr cmd/main.go

generate:
	go generate cmd/main.go

install:
	go install

run-prod:
	docker-compose -f docker-compose.yml up