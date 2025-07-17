FROM golang:1.24.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go","run","cmd/server/main.go"]
