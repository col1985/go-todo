# build
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o ./go-todo

# Deploy
FROM golang:1.23.1-alpine

WORKDIR /

COPY --from=builder /app/go-todo /go-todo

COPY --from=builder /app/.env /

EXPOSE 8080

CMD ["/go-todo"]