# Use the official Golang image to create a build artifact.
# This is the "builder" stage.
FROM golang:1.25-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# CGO_ENABLED=0 builds a statically linked binary
# -o /rate-limiter places the output binary in the root of the builder image
RUN CGO_ENABLED=0 GOOS=linux go build -o /rate-limiter main.go

# Start a new, scratch image to reduce the final image size.
# This is the "runtime" stage.
FROM alpine:3.14

# Copy the built binary from the "builder" stage
COPY --from=builder /rate-limiter /rate-limiter

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/rate-limiter"]