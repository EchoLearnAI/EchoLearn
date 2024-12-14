# Stage 1: Build stage
FROM golang:1.21-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the source code
COPY . .

# Build the Go binary statically
# -ldflags "-s -w" is used to reduce the binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o server main.go

# Stage 2: Final image
# Using distroless as the base image
FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=builder /app/server /server

# Distroless images run as non-root by default
USER nonroot:nonroot

# Expose the server port
EXPOSE 8081

# Run the server
ENTRYPOINT [ "/server" ]