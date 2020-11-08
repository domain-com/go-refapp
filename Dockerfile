FROM gcr.io/domaincom-dev/golang:1.13.8-alpine3.11 AS builder

RUN apk add --no-cache git=2.24.3-r0
RUN mkdir /app
COPY . /build/
WORKDIR /build
RUN go get -v -d ./...
RUN go build -o main .

FROM gcr.io/domaincom-dev/alpine:3.11

LABEL maintainer="Alejandro Pino Oreamuno <apinoo@gmail.com>"

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
