FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 8080

CMD ["./main"]