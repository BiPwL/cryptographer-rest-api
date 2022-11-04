.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.DEFAULT_GOAL := build

#docker run -v /home/bipwl/code/github.com/BiPwL/cryptographer-rest-api/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://admin:root@localhost:5432/db?sslmode=disable up
#migrations