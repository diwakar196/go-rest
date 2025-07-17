# --- Build Stage ---
    FROM golang:1.24.5 AS builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/user_server/main.go
    
    # Debug: confirm binary is created
    RUN ls -la /app
    
    # --- Final Stage ---
    FROM alpine:3.20
    
    WORKDIR /app
    
    COPY --from=builder /app/server .
    
    # Debug: confirm binary is copied
    RUN ls -la /app
    
    RUN chmod +x server
    
    EXPOSE 8080
    
    ENTRYPOINT ["./server"]
    