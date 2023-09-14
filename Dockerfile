# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copy only the necessary Go files and dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Stage 2: Create a minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port your Go application listens on (change to match your app)
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
