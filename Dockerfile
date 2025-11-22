# ---------- builder ----------
# Use an ARG so GitHub Actions can pass the version string
ARG VERSION
FROM golang:latest AS builder

# make build args visible in this stage
ARG VERSION

# Build environment
ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /build

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources
COPY . .

# Build the binary and inject VERSION into main.Version
# Note: ensure your code declares `package main` and `var Version = "dev"`
RUN go build -ldflags="-X 'main.Version=${VERSION}'" -o /build/app

# ---------- production ----------
ARG VERSION
FROM alpine:latest AS production

# Add certificates for TLS
RUN apk add --no-cache ca-certificates

WORKDIR /go-resolve
RUN mkdir -p config

# Copy built binary from builder
COPY --from=builder /build/app .

# Run the app
CMD ["./app"]
