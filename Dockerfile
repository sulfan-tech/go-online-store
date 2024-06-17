FROM golang:1.20-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .

RUN go build -o main ./server/cmd

EXPOSE 8080

CMD ["./main"]
