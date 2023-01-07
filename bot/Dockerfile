FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o bot bot.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/bot /build/bot

CMD ["./bot"]