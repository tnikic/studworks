FROM golang:1.24 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o studworks .

FROM alpine:3
WORKDIR /app
COPY --from=builder /app/studworks .
EXPOSE 8080
ENTRYPOINT ["./studworks"]
