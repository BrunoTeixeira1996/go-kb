FROM golang:1.18

RUN mkdir /app
RUN mkdir /app/notes

COPY . /app

WORKDIR /app

RUN go build -o server .

EXPOSE 8080

ENTRYPOINT ["/app/server"]
