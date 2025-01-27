# Build the Go application
FROM golang:1.23-alpine AS builder

# Install git (required for some Go modules)
RUN apk update && apk add --no-cache git

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Debugging: List files to confirm go.sum is present
RUN ls -la

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o backend .

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/backend .

# Expose the application's port
EXPOSE 8080

# Command to run the executable
CMD ["./backend"]
