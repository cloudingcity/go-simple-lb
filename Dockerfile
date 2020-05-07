FROM golang:1.13-alpine as builder
COPY . /app
WORKDIR /app
RUN go build -o lb

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/lb .
ENTRYPOINT ["./lb"]
