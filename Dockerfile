FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o my-app .


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/my-app .
CMD ["./my-app"]