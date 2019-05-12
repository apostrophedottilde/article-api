## This will not build a production-ready image. 
## This is only for local devloment purposes.
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache && apk add gcc && apk add musl-dev && apk add git && mkdir app
COPY . /app
WORKDIR /app
RUN go get -d -v
RUN go mod download
RUN go test ./...
RUN go build
ENTRYPOINT ["go", "run", "main.go"]