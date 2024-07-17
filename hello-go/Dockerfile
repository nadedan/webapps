FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .

RUN go build -o hello-go .

EXPOSE 8080

CMD ["./hello-go"]