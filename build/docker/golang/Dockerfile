FROM golang:alpine

RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /go/src/app

RUN go get -u github.com/mitranim/gow

COPY go.mod go.mod
RUN go mod tidy

#COPY . .
#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN rm -f go.mod go.sum