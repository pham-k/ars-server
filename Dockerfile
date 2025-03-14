# Build stage
FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/app -v ./cmd/api

# Final stage
FROM alpine:latest

RUN addgroup -S app && adduser -S app -G app
COPY --from=builder --chown=app /go/bin/app /app
COPY --from=builder --chown=app /go/src/app/config.toml /config.toml
USER app

EXPOSE 8080

ENTRYPOINT ["/app"]