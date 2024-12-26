FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config .
COPY app .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app/main.go .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/config .

ENV PROJECT_CONFIG_FILE="/app/config/dev.conf"

EXPOSE 8083

CMD ["./main"]