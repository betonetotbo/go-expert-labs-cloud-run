FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o consulta-clima cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/consulta-clima .
ENTRYPOINT ["consulta-clima"]