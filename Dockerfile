# Stage 1: Building the Go application 
FROM golang:latest AS builder

WORKDIR /app


# Use a lightweight debian os
# as the base image
FROM debian:stable-slim

# execute the 'echo "hello world"'
# command when the container runs
CMD ["echo", "hello world"]