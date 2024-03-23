# Start with a base image containing the Go runtime
FROM golang:1.17-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o scraper .

# Start a new stage to create a lightweight container for the application
FROM alpine:latest

# Set the working directory to the app directory
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=builder /app/scraper .

# Run the application
CMD ["./scraper"]
