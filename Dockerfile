
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
RUN apt-get update && apt-get install -y ca-certificates

CMD ["./main"]
