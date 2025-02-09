FROM golang:1.23.5-alpine3.21 AS builder
WORKDIR /build
COPY ./app/. .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /build/main .
EXPOSE 8090
CMD ["./main"]

