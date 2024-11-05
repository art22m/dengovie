FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd/
COPY internal ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

CMD ["./main"]


FROM alpine

EXPOSE 8080

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --chown=app:app --from=builder /app/main /app/main

USER app

CMD ["./main"]