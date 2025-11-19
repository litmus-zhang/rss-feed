# Build stage
# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main cmd/main.go

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/internal/db/migrations ./migrations

EXPOSE 8080
CMD ["./main"]


# keploy record -c "docker run -p <appPort>:<hostPort> --name <containerName> --network keploy-network --rm <applicationImage>" --containerName "<containerName>" --delay 10

# keploy record -c "docker run -p 8080:8080 --name rss-feed --network keploy-network --rm rss-feed" --containerName "rss-feed" --delay 10

# keploy record -c "docker compose up" --container-name <containerName> --build-delay 100
# keploy record -c "docker compose up" --container-name rss-feed --build-delay 100

# keploy test -c "docker compose up" --container-name <containerName> --build-delay 50 --delay 20
# keploy test -c "docker compose up" --container-name rss-feed --build-delay 50 --delay 50