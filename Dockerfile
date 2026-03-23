FROM golang:1.26.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . .

CMD ["go","run","main.go"]