FROM golang:1.19.3-alpine3.16

COPY src /app/src
WORKDIR /app/src

RUN go build .

CMD ["/app/src/simple_http_chatapp"]
