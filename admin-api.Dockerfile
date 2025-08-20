# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the admin-api binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/admin-api ./cmd/admin-api

# Stage 2: Create the final, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/admin-api .

# Copy the configuration file
COPY ./configs/config.yaml ./configs/

# Expose the port the API runs on
EXPOSE 8081

# Command to run the application
CMD ["./admin-api"]
