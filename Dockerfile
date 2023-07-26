FROM golang:1.20-alpine3.15 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

# Start a new stage from scratch

FROM alpine:3.15

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app ./

CMD ["./app"] 