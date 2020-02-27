# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go get "github.com/auth0/go-jwt-middleware"

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]