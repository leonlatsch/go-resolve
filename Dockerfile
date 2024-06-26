FROM golang:latest as builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

# Add a work directory
WORKDIR /build

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Build app
RUN go build -o app

FROM alpine:latest as production

# Add certificates
RUN apk add --no-cache ca-certificates

WORKDIR /go-resolve
RUN mkdir config

# Copy built binary from builder
COPY --from=builder build/app .
ENV GIN_MODE release

# Expose API port
EXPOSE 9991

# Exec built binary
CMD ./app

