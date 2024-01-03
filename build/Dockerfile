FROM golang:1.21-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Grigorev Mikhail <d0rm4nt.4rt@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main internal/cmd/server/main.go
RUN set -a && . ./build/.env && set +a

EXPOSE 3001

CMD ["./main"]