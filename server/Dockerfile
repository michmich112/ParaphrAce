FROM golang:1.16-alpine

WORKDIR /usr/local/server

COPY . .

RUN go get
RUN go build -o server .

EXPOSE 8080

CMD ./server

