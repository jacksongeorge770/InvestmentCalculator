# ---------- Stage 1: Build the Go binaries ----------
    FROM golang:1.22 AS builder

    # Set working directory inside container
    WORKDIR /app
    
    # Copy go.mod and go.sum first to leverage Docker layer caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the entire source code
    COPY . .
    
    # Build the backend main API binary
    RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/cal ./cmd/main.go
    
    # Build the migration tool binary
    RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/migrate ./cmd/migrate/main.go
    
    # ---------- Stage 2: Create a lightweight final image ----------
    FROM alpine:latest
    
    # Install certificates for HTTPS (e.g., calling secure APIs or DBs)
    RUN apk --no-cache add ca-certificates
    
    # Set working directory in final container
    WORKDIR /app
    
    # Copy built Go binaries from builder stage
    COPY --from=builder /app/bin/cal /app/bin/migrate ./
    
    # Copy environment config
    COPY .env ./
    
    # Copy SQL migrations folder
    COPY cmd/migrate/migration ./migrations

    
    # Copy frontend static files
    COPY frontend ./frontend
    
    # Expose backend server port
    EXPOSE 8080
    
    # Start the backend server
    CMD ["/app/cal"]
    