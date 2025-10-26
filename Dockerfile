# Etapa 1: build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copiamos go.mod y go.sum primero (para aprovechar cache)
COPY go.mod ./
COPY . .

# Compilamos el binario
RUN go build -o api ./cmd/api

# Etapa 2: runtime (más liviano)
FROM alpine:3.20

WORKDIR /app

# Copiamos solo el binario compilado
COPY --from=builder /app/api .

# Exponemos el puerto donde corre la app
EXPOSE 8080

# Comando de ejecución
CMD ["./api"]
