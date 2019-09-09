FROM golang:1.13 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor ./cmd/quo
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/quo /quo
ENV PORT 8080
EXPOSE 8080
CMD ["/quo"]