FROM golang:1.15.5-alpine3.12 AS build
ARG LD_FLAGS
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -ldflags "$LD_FLAGS" -o /app/bin/text-event-service ./src

FROM scratch AS bin
ARG EXPOSED_PORTS
COPY --from=build /app/bin/text-event-service /
COPY --from=build /app/resources /resources
EXPOSE $EXPOSED_PORTS
ENTRYPOINT ["/text-event-service"]