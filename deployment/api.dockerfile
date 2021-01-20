FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG port
ARG port_service_url

ENV PORT=$port
ENV SVC=$port_service_url

ENTRYPOINT go run cmd/clientapi/main.go
