# Use the official Golang image to create a build artifact
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN GO111MODULE=on go build -o simple-main ./server/cmd

# Use the official Alpine image for a lean production image
FROM alpine:latest
WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/simple-main /usr/local/bin/app

# Command to run the executable
CMD ["/usr/local/bin/app"]
