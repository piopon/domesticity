FROM golang:1.15.5-alpine3.12 AS build
WORKDIR /app
COPY . .
RUN go build -o /app/bin/text-event-service ./src

FROM scratch AS bin
COPY --from=build /app/bin/text-event-service /
ENTRYPOINT ["/text-event-service"]