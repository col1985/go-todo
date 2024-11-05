# build
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

ARG CGO_ENABLED
ARG GOOS
ARG GOARCH

RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./go-todo

# Deploy
FROM golang:1.23.1-alpine

WORKDIR /

COPY --from=builder /app/go-todo /go-todo

EXPOSE 8080

CMD ["/go-todo"]