FROM golang:1.21.4-alpine3.18

WORKDIR /app

COPY . .

RUN go build -o crup-api .

EXPOSE 8000

CMD ["./crup-api"]