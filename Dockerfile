FROM golang:1.17-alpine
WORKDIR /app
ADD . .
RUN go build ./cmd/dns/...
CMD ["/app/dns"]