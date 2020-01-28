FROM golang:1.13.6-alpine AS builder

RUN mkdir /app
ADD . /build/
WORKDIR /build
RUN apk add git
RUN go get -v -d ./...
RUN go build -o main .

FROM alpine

LABEL maintainer="Alejandro Pino Oreamuno <apinoo@gmail.com>"

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
