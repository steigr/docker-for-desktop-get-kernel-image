FROM golang:1.14.2 AS builder
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN  go mod download
COPY main.go main.go
RUN  CGO_ENABLED=0 go build -o /bin/docker-for-desktop-get-kernel-image main.go

FROM alpine:3.11
COPY --from=builder /bin/docker-for-desktop-get-kernel-image /bin/docker-for-desktop-get-kernel-image
ENTRYPOINT ["docker-for-desktop-get-kernel-image"]
