FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o  auth-service cmd/main/main.go

EXPOSE 8080

CMD ["./auth-service"]

