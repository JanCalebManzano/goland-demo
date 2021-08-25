FROM golang:1.17.0-alpine3.14

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]