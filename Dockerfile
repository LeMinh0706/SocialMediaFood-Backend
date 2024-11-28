# build stage
FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . . 

RUN go build -o backend cmd/main.go

EXPOSE 8070

CMD ["./backend"]