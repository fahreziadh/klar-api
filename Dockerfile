FROM golang:latest
RUN mkdir -p /go/src/github.com/klar
WORKDIR /go/src/github.com/klar
RUN apk add --no-cache git
COPY . /go/src/github.com/klar
RUN go-wrapper download
RUN go-wrapper install

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=cloud-native-go Version=0.0.1
EXPOSE 3000