FROM golang:1.24.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o server cmd/server/main.go

EXPOSE 8080

CMD ["./server"]
