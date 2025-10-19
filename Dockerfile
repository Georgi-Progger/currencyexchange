FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOBIN=/usr/local/bin go install github.com/pressly/goose/v3

RUN chmod +x run.sh

RUN go build -o main ./cmd

EXPOSE 8080

CMD ./run.sh && ./main
