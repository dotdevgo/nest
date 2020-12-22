FROM app/golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o gateway ./cmd/gateway/main.go

CMD ["./gateway"]

#RUN go mod tidy
#RUN go get -d -v ./...
#RUN go install -v ./...
