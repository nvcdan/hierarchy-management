FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Etapa de testare
FROM builder as tester

WORKDIR /app

CMD ["go", "test", "-v", "./..."]

# Etapa finală
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main . 
COPY --from=builder /app/.env .

CMD ["./main"]