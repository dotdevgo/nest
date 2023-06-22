GOPATH := $(shell go env GOPATH)

.PHONY: init \
		start-dev \
		test lint gofmt godoc generate \
		gqlgen

init:
	go get -u gorm.io/gorm
	go get gorm.io/datatypes
	go get github.com/google/uuid
	go get github.com/go-playground/validator/v10
	go get -u github.com/gotidy/copy
	go get -u github.com/pilagod/gorm-cursor-paginator
	go get github.com/psampaz/slice
	go get github.com/asaskevich/EventBus
	go get github.com/Masterminds/goutils
	go get github.com/joho/godotenv
	go install github.com/phelmkamp/metatag@latest
	go install github.com/mitranim/gow@latest

# Examples
start-dev:
	gow run cmd/api/main.go

# Graphql
gqlgen:
	cd graphql && go run github.com/99designs/gqlgen generate .

# Golang
test:
	go test -v ./... -cover

lint:
	golangci-lint run ./... && \
	go vet ./...

gofmt:
	go fmt ./...

godoc:
	godoc -http=:8820 -goroot "$(GOPATH)"

generate: gqlgen
	go generate ./...
