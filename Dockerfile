# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12

LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o main .
EXPOSE 3000
# Command to run the executable
CMD ["./main"]