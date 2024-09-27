# https://dev.to/plutov/docker-and-go-modules-3kkn
FROM golang:1.21.0 as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o metrics ./cmd/go-ddd/main.go

ENTRYPOINT ["/app/metrics"]