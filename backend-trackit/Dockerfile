
FROM golang:1.22-alpine AS builder


WORKDIR /app


RUN apk add --no-cache git gcc musl-dev


COPY go.mod go.sum ./


ENV GOTOOLCHAIN=local


RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /app


RUN apk --no-cache add ca-certificates tzdata


COPY --from=builder /app/main .
COPY --from=builder /app/.env* ./


EXPOSE 8080


CMD ["./main"]