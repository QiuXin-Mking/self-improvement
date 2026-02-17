# Multi-stage build for the spaced repetition app

# Backend build stage
FROM golang:1.21-alpine AS backend-builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o web_server web_server.go

# Frontend build stage
FROM node:18-alpine AS frontend-builder

WORKDIR /app

# Copy frontend files
COPY ./frontend/package*.json ./
COPY ./frontend/ ./

# Install dependencies
RUN npm install

# Build the frontend
RUN npm run build

# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Install necessary packages for SQLite support
RUN apk add --no-cache sqlite

# Copy the backend binary
COPY --from=backend-builder /app/web_server .
COPY --from=backend-builder /app/templates ./templates
COPY --from=backend-builder /app/static ./static
COPY --from=backend-builder /app/questions ./questions

# Copy the built frontend
COPY --from=frontend-builder /app/dist ./frontend/dist

# Copy other necessary files
COPY migrations ./migrations
COPY question_input* ./

# Expose port
EXPOSE 5000

# Command to run the application
CMD ["./web_server"]