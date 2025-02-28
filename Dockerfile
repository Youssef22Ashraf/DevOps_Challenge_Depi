# ---- Build Stage ----
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Install Git (needed for some dependencies)
RUN apk add --no-cache git

# Copy application code & dependencies
COPY app/ /app

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o myapp ./main.go

# ---- Run Stage ----
FROM alpine:latest
WORKDIR /app

# Copy only the built binary (no Golang compiler needed)
COPY --from=builder /app/myapp /app/myapp
COPY app/.env /app/.env  

# Expose application port
EXPOSE 80

# Run the application with environment variables
CMD env $(cat /app/.env | xargs) /app/myapp
