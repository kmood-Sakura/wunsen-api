# go-api/Dockerfile
FROM golang:1.24.4-alpine AS builder

# Set working directory
WORKDIR /app

# Copy all Go code to /app
COPY . .

# Run go mod tidy to clean up dependencies
RUN go mod tidy

# Build the main application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/

# Final stage - clean Alpine image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy ONLY the compiled binary from builder stage (all Go code deleted)
COPY --from=builder /app/main .

# Copy config directory if it exists
COPY --from=builder /app/config ./config

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port 4000
EXPOSE 4000

# Run main binary on port 4000
CMD ["./main"]