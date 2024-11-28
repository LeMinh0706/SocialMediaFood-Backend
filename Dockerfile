# build stage
FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . . 
RUN go build -o backend cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 8070

CMD ["./backend"]