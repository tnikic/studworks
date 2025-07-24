FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o studworks .

FROM alpine:3
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/studworks .
EXPOSE 8080
ENV GODEBUG=tlsrsakex=1
ENTRYPOINT ["./studworks"]
