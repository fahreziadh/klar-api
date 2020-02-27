FROM golang:latest

LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

EXPOSE 3000

CMD ["./app"]