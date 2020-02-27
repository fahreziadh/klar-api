# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

ENV GO111MODULE=on
# Add Maintainer Info
LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app/server

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN ls   

COPY go.mod .
COPY go.sum .
RUN go mod download
# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]