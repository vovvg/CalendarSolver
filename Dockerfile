# Используем образ Golang
FROM golang:1.24 AS builder

WORKDIR /app

ENV TELEGRAM_APITOKEN=${TELEGRAM_APITOKEN}
ENV RUN_MODE=${RUN_MODE}

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
WORKDIR /app/cmd/calendar-solver
RUN go build -o /app/main .

CMD ["/app/main"]
