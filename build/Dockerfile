FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o consulta-clima cmd/main.go
RUN ["touch", ".env"]

FROM alpine:3.20.3
WORKDIR /app
COPY --from=builder /app/consulta-clima .
COPY --from=builder /app/.env .
ENTRYPOINT ["/app/consulta-clima"]