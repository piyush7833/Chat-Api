# Build stage
FROM golang:1.22.2 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Final stage
FROM ubuntu AS final

WORKDIR /app

COPY --from=builder /app .

# Expose port 8080
EXPOSE 8080

# Set environment variables

ARG JWT_SECRET
ARG DB_URI
ENV JWT_SECRET=${JWT_SECRET}
ENV DB_URI=${DB_URI}
ENV ENVIRONMENT="production"


CMD ["./main"]
