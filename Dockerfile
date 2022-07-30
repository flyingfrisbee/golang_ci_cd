FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8080 8081 8082 8083

CMD ["/app/main"]