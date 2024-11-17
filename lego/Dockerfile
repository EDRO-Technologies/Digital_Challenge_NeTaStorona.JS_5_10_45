FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lego-server .

FROM alpine:latest

WORKDIR /app

RUN apk add ca-certificates

COPY --from=builder ./app/lego-server .

COPY --from=builder ./app/.env ./.env


CMD ["./lego-server"]
