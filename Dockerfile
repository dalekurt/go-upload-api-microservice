# Build
FROM golang:1.20.0-alpine3.17 AS builder

# Install dependencies to be able to build our code
RUN apk add --no-cache --update git && apk add build-base

ENV GO111MODULE=on

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

# Download all required packages
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN CGO_ENABLED=0 go build main.go

# Generate clean, finale image
FROM alpine:latest as server

WORKDIR /app

COPY --from=builder /app/main ./

RUN chmod +x ./main

CMD [ "./main" ]
