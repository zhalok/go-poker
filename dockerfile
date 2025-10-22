# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o main ./app

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

CMD ["tail", "-f", "/dev/null"]
