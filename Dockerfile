# syntax=docker/dockerfile:1

# Stage 1 - build the binary
FROM golang:1.24.3 as builder

WORKDIR /app

# Copy go.mod and go.sum before copying the whole project
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/api

# Stage 2 - run the binary
FROM gcr.io/distroless/base-debian11

WORKDIR /root/
COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080
ENTRYPOINT ["./server"]
