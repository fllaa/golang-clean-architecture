FROM golang:1.24-alpine AS builder

LABEL maintainer="Fallah Andy Prakasa <me@flla.my.id> (https://flla.my.id/)"

# Define the application directory.
ARG APP

# Install necessary packages.
RUN apk update && apk add \
    # explicitly install SASL package
    cyrus-sasl-dev \
    gcc \
    musl-dev

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -ldflags='-w -s' -tags musl -o app ./cmd/${APP}/main.go
RUN ls -la

FROM alpine:3.21

# Copy binary and config files from /build to root folder of alpine container.
COPY --from=builder ["/build/app", "/build/config.json", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/app"]
