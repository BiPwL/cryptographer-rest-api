.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.DEFAULT_GOAL := build

# docker run -v /home/bipwl/code/github.com/BiPwL/cryptographer-rest-api/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://admin:root@localhost:5432/db?sslmode=disable up
# migrations

# curl -X POST -H "Content-Type: application/json" \
-d "{\"email\": \"user@example.org\", \"password\": \"password\"}" -b http://localhost:8080/users
# the curl sould look somthing like this

# for save cookie use -c and path, and for use cookie use -b and path

# for set browser(header) use -H "Origin: google.com" for example