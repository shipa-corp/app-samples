FROM golang:1.15-alpine AS builder

COPY main.go /app/
WORKDIR /app
RUN go build /app/main.go 

FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/app
COPY shipa.yaml /app/
COPY Procfile /app/
