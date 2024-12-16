FROM golang:latest as builder
WORKDIR /go/src/otel-hello
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/otel-hello/main .
ENTRYPOINT ["./main"]
