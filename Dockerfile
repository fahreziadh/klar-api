FROM golang:latest

LABEL maintainer="Fahrezi <fahreziadh@gmail.com>"

WORKDIR /home/ubuntu/go/src/github.com/klar/

RUN go build -o app .

EXPOSE 3000

CMD ["./app"]