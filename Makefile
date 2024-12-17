GOPATH := $(shell go env GOPATH)

init:
	go get -u gorm.io/gorm
	go get gorm.io/datatypes
	go get github.com/google/uuid
	go get github.com/go-playground/validator/v10
	go get -u github.com/gotidy/copy
	go get -u github.com/pilagod/gorm-cursor-paginator
	go get github.com/psampaz/slice
	go get github.com/Masterminds/goutils
	go get github.com/joho/godotenv
	go install github.com/phelmkamp/metatag@latest
	go install github.com/mitranim/gow@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/echo-swagger

# Development
serve:
	gow run main.go

test:
	go test -v ./... -cover

lint:
	golangci-lint run ./... && \
	go vet ./...

gofmt:
	go fmt ./...

godoc:
	godoc -http=:1333 -goroot "$(GOPATH)"

generate:
	go generate ./...

swag:
	swag init
