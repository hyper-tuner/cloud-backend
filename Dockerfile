FROM golang:1.19 AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o cloud-backend

FROM alpine:latest AS production

COPY --from=builder /app .

EXPOSE 8090

CMD ["./cloud-backend", "serve", "--http=0.0.0.0:8090"]