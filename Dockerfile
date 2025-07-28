# Stage 1: Build static Go binary
FROM golang:1.24 as builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/api

# Stage 2: Run with no OS dependencies
FROM gcr.io/distroless/static

WORKDIR /root/
COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080
ENTRYPOINT ["./server"]
