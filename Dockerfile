# ---- Build stage ----
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app .

# ---- Runtime stage (tiny, non-root) ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app /app
ENV GIN_MODE=release
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
