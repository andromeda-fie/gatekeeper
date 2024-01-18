VERSION 0.7
FROM golang:1.21.5-alpine3.19
WORKDIR /gatekeeper

test:
	COPY go.mod ./
	COPY . ./
	RUN go test ./...

build:
	COPY go.mod ./
	COPY . ./
	RUN go build ./...
