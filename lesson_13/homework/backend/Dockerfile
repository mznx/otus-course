# BUILD
FROM golang:alpine AS builder

ENV CGO_ENABLED=0

ENV GOOS=linux

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o /backend .

# RUN
FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

RUN apk add --no-cache tzdata=2025b-r0

ENV TZ="Europe/Moscow"

EXPOSE 3000

WORKDIR /app

COPY --from=builder /backend .

CMD ["./backend"]
