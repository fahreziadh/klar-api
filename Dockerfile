# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

ENV GO111MODULE=on
WORKDIR /app/server
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build 
# Command to run the executable
CMD ["./main"]