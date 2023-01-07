FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o server server.go db.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/server /build/server

CMD ["./server"]