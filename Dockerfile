FROM golang:1.26.0

WORKDIR /app

COPY api/ ./api/

WORKDIR /app/api

RUN go mod tidy

CMD ["go","run","main.go"]