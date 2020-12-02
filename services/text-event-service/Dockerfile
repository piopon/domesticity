FROM golang:1.15.5-alpine3.12
WORKDIR /app
COPY . .
RUN go build -o /app/bin/text-event-service ./src