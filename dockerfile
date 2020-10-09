FROM golang:1.15.2

WORKDIR /go/src/sports-parser

COPY . .

RUN go get ./...

ENTRYPOINT go run Parser-for-sports.go