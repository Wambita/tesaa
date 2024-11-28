# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Stage 2: Run the JSON server
FROM node:alpine AS json-server

# Install json-server globally
RUN npm install -g json-server

# Create a data file for JSON server
COPY api/database/users.json /api/database/users.json

# Stage 3: Define the final stage
FROM alpine:latest

# Install Go, Node.js, and npm in the final image
RUN apk add --no-cache go nodejs npm

# Install json-server globally in the final stage
RUN npm install -g json-server

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go application code (since we are using go run .)
COPY . .

# Expose the necessary ports
EXPOSE 8080

# Start the JSON server and then the Go application using go run .
CMD json-server --watch ./api/database/users.json --port 3000 & go run .
