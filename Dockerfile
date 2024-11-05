FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd/
COPY internal ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go
RUN addgroup -S app && adduser -S app -G app

CMD ["./main"]

FROM scratch

EXPOSE 8080

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER app

USER app

CMD ["./main"]