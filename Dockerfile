FROM golang:1.19-alpine3.15 as builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o cafebackend ./cmd/cafebackend/main.go

FROM alpine:3
COPY --from=builder /build/cafebackend .
COPY --from=builder /build/config.toml .
ENTRYPOINT ["./cafebackend"]
