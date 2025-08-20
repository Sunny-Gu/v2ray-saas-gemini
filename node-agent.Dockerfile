# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the node-agent binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/node-agent ./cmd/node-agent

# Stage 2: Create the final, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/node-agent .

# Note: The node-agent might need its own config file in a real scenario,
# but for now, we assume it connects to localhost services.

# Expose the port the agent runs on
EXPOSE 8083

# Command to run the application
CMD ["./node-agent"]
