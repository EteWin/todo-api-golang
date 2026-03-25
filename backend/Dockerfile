FROM golang:1.25

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

COPY ./backend/. ./

CMD ["air", "-c", ".air.toml"]