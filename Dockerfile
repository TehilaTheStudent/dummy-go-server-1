# Use Go base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Build the Go app
RUN go build -o server main.go

# Expose port
EXPOSE 8080

# Run the server
CMD ["./server"]
