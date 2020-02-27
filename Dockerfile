# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12

LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

ENV GO111MODULE=on
WORKDIR /
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build 
# Command to run the executable
CMD ["./main"]