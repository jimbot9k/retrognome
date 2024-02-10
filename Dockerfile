FROM golang:1.17 AS go-builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
RUN go build -o retrognome cmd/main.go

CMD ["./retrognome"]