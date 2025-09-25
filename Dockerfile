FROM golang:1.25-alpine3.22  AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY fancy-adventure/go.mod fancy-adventure/go.sum ./
RUN go mod download
COPY fancy-adventure/ .
RUN go build -o fancy-adventure main.go
FROM alpine AS release
WORKDIR /app
COPY --from=builder /app/fancy-adventure .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./fancy-adventure"]
