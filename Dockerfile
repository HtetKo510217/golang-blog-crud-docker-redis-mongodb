# Use an official Golang runtime as a base image
FROM golang

# Set the working directory in the container
WORKDIR /app

# Copy only the necessary Go modules files
COPY go.mod go.sum ./

# Clear the Go module cache
RUN go clean -modcache

# Download Go modules
RUN go mod download

# Copy the rest of the source code to the container
COPY . .

# Build the Golang application
RUN go build -o main .

# Expose port 3000
EXPOSE 3000
# Command to run the executable
CMD ["./main"]