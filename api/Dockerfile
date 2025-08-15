# Build
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN apk add --no-cache build-base && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# Run
FROM gcr.io/distroless/base-debian12
WORKDIR /
ENV PORT=8080
COPY --from=build /app/server /server
EXPOSE 8080
USER nonroot:nonroot
CMD ["/server"]