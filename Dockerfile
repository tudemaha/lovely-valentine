FROM golang:1.25.5-alpine3.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN touch /app/.env

RUN go build -v -o app

CMD ["./app"]