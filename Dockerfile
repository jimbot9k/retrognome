FROM node:14 AS react-builder
WORKDIR /app/react
COPY react/ .
RUN npm install
RUN npm run build

FROM golang:1.17 AS go-builder
WORKDIR /app
COPY . .
RUN go build -o retrognome cmd/main.go

CMD ["./retrognome"]