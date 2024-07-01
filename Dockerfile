FROM golang:1.22-alpine

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o crup-api cmd/main.go

EXPOSE 8000

CMD ["./crup-api"]
