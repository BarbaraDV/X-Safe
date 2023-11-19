FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

COPY go.mod .
COPY go.sum .
COPY backend/ /go/src/backend/

# Build Go binary
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod download 
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod go build gallery/backend/cmd/serve

FROM alpine:latest

COPY --from=builder /go/src/serve /serve

RUN apk add --no-cache \
    unzip \
    ca-certificates \
    # this is needed only if you want to use scp to copy later your pb_data locally
    openssh

# uncomment to copy the local pb_migrations dir into the container
# COPY ./pb_migrations /pb/pb_migrations

# uncomment to copy the local pb_hooks dir into the container
# COPY ./pb_hooks /pb/pb_hooks

EXPOSE 8080

# start PocketBase
CMD ["/serve"]
