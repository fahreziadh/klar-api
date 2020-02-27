FROM golang:latest

LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

WORKDIR /app

RUN go build -o app .

EXPOSE 3000

CMD ["./app"]