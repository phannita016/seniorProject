# Use official Golang image as a base
FROM golang:1.23


# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and dependencies
# COPY go.mod go.sum ./
COPY . .
# COPY config.yaml /app/config.yaml

RUN go install github.com/air-verse/air@latest
# Download dependencies
RUN go mod download

# Build the Go application
# RUN go build -o ./tmp/main ./cmd

# Expose application port
EXPOSE 8081

# Start the application
# CMD ["./main"]
CMD ["air", "-c", ".air.toml"]
