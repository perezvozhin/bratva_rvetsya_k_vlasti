FROM golang:1.25-alpine
# Для CGO
RUN apk add --no-cache gcc musl-dev
ENV CGO_ENABLED=1
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
WORKDIR /migrations
ENTRYPOINT ["goose"]