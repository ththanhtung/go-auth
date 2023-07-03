# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the Go app
RUN go build -o app .

# Expose port 8080 for the container
EXPOSE 8080

# Define the command to run the executable
CMD ["./app"]