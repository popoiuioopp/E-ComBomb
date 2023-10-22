FROM golang:1.21.3-bookworm

WORKDIR /app

COPY ./backend/go.mod ./backend/go.sum /app/
RUN go mod download && go mod verify

COPY ./backend/ /app/

RUN go build -v -o /app/ ./...

EXPOSE 8080

CMD ["./craft-cart"]