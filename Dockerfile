FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git mysql-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Set execute permission for the wait-for-db.sh script
RUN chmod +x wait-for-db.sh

RUN go build -o ./bin/api ./cmd/api

CMD ["sh", "-c", "./wait-for-db.sh db 3306 && ./bin/api"]
