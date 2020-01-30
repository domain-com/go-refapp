FROM golang:1.13.7-alpine3.11 AS builder

RUN mkdir /app
COPY . /build/
WORKDIR /build
RUN apk add --no-cache git=2.24.1-r0
RUN go get -v -d ./...
RUN go build -o main .

FROM alpine:3.11

LABEL maintainer="Alejandro Pino Oreamuno <apinoo@gmail.com>"

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
