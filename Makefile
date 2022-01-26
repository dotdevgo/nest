GOPATH := $(shell go env GOPATH)

.PHONY: init
init:
	go get -u gorm.io/gorm
	go get gorm.io/datatypes
	go get -u gorm.io/driver/sqlite
	go get github.com/google/uuid
	go get github.com/go-playground/validator/v10
	go get -u github.com/gotidy/copy
	go get -u github.com/pilagod/gorm-cursor-paginator
	go get github.com/psampaz/slice
	go get github.com/phelmkamp/metatag
	go install github.com/mitranim/gow@latest

.PHONY: start-dev
start-dev:
	 gow run cmd/gateway/main.go

.PHONY: generate
generate: gqlgen
	go generate ./...

.PHONY: gqlgen
gqlgen:
	cd graphql && go run github.com/99designs/gqlgen generate .

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: lint
lint:
	golangci-lint run ./... && \
	go vet ./...

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: godoc
godoc:
	godoc -http=:8820 -goroot "$(GOPATH)"
