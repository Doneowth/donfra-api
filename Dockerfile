FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git ca-certificates
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/donfra-api ./cmd/donfra-api

FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app
RUN apk add --no-cache python3 ca-certificates
WORKDIR /home/app
COPY --from=builder /out/donfra-api /usr/local/bin/donfra-api
USER app
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/donfra-api"]
