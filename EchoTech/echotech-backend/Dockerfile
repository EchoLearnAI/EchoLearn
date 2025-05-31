# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 for a static binary, GOOS=linux to ensure Linux binary
# -ldflags "-s -w" to strip debug information and reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo -ldflags "-s -w" -o /echotech-backend .

# Stage 2: Create a minimal final image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /echotech-backend .

# Copy .env file - it should be created in the same directory as Dockerfile or mounted
# COPY .env . 
# For production, manage .env securely (e.g., Docker secrets, k8s configmaps/secrets)

# Copy Swagger docs
COPY --from=builder /app/docs ./docs

# Expose port 8080 (the port our Go app listens on)
EXPOSE 8080

# Command to run the executable
# The JWT_SECRET_KEY and database credentials should be passed as environment variables
# when running the container, not hardcoded or solely reliant on a .env file in the image for production.
CMD ["./echotech-backend"] 