FROM golang:1.21 AS builder
WORKDIR /app
COPY main.go .
RUN go build -o my-app main.go
CMD ["./my-app"]

FROM scratch
WORKDIR /app
COPY --from=builder /app/my-app .
CMD ["./my-app"]