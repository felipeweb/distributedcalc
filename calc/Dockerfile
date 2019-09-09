FROM golang:1.13 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor ./cmd/calc
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/calc /calc
ENV PORT 8080
EXPOSE 8080
CMD ["/calc"]