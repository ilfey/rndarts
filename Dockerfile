FROM golang:1.23.1 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/src/rndarts

COPY go.mod go.sum ./

RUN go mod download

# Copy build dependencies.
COPY ./pkg ./pkg
COPY ./main.go ./main.go

# Build the api.
RUN go build -v -o ./bin/bot .

FROM alpine:latest

WORKDIR /var/www/rndarts

RUN apk add doas

# Copy the api.
COPY --from=builder /usr/src/rndarts/bin/bot .
