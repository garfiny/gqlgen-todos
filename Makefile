generate:
	go run github.com/99designs/gqlgen generate

build:
	GO111MODULE=on go build -o ./cmd/server server.go

run: build
	./cmd/server