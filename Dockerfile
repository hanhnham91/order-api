# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

RUN mkdir publish

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all necessary dependencies
RUN go mod download -x

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o my-order-service cmd/main.go

# Create a minimal runtime image based on Alpine Linux
FROM alpine:latest
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/my-order-service .

# Expose the port your application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./my-order-service"]