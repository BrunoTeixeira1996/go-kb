FROM golang:1.18 AS build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server .

EXPOSE 8080


ENTRYPOINT ["/app/server"]
