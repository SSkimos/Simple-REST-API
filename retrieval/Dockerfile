FROM golang:1.22.1-alpine

RUN apk update && apk add --no-cache git bash build-base

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download

RUN go build -o /build/main ./cmd/

EXPOSE 80

CMD ["/build/main"]
