# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files including .env
COPY go.mod go.sum .env ./

# Display the list of files in the working directory
RUN ls -la

# Download dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp ./server/cmd

# Stage 2: Set up the final image
FROM alpine:latest
WORKDIR /root/
# Copy the .env file to the correct location
COPY --from=builder /app/.env /root/.env
COPY --from=builder /app/myapp .
CMD ["./myapp"]
