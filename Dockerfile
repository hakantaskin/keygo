FROM golang:1.17.3 AS builder
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates wget
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

FROM alpine:3.13.5
RUN apk update \
 && apk add --update tzdata  \
 && apk --update add ttf-ubuntu-font-family  \
 && cp /usr/share/zoneinfo/UTC /etc/localtime  \
 && echo "UTC" > /etc/timezone
WORKDIR /app
COPY --from=builder /app/main ./main
CMD ["/app/main"]