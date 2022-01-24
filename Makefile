GOPATH := $(shell go env GOPATH)

.PHONY: init
init:
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get github.com/google/uuid
	go get github.com/go-playground/validator/v10
	go get -u github.com/gotidy/copy
	go get -u github.com/pilagod/gorm-cursor-paginator
	go get github.com/psampaz/slice
	go install github.com/mitranim/gow@latest

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: generate
generate:
	go generate ./...
