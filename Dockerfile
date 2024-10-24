FROM golang:1.22.8-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN touch /app/.env

RUN go build -v -o app

CMD ["./app"]