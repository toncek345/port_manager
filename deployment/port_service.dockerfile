FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG port
ARG database_connection

ENV PORT=$port
ENV DB=$database_connection

ENTRYPOINT go run cmd/portdomainsvc/main.go
