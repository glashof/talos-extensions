FROM golang:1.25-alpine AS builder

ARG NETBIRD_VERSION=0.66.0

RUN apk add --no-cache git

RUN git clone --branch v${NETBIRD_VERSION} --depth 1 https://github.com/netbirdio/netbird.git /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags "-s -w -X github.com/netbirdio/netbird/version.version=${NETBIRD_VERSION}" \
    -o /netbird ./client

FROM golang:1.23-alpine AS uname-builder

COPY netbird/src/ /src/
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /uname .

FROM scratch

COPY netbird/manifest.yaml /manifest.yaml
COPY --from=builder /netbird /rootfs/usr/local/lib/containers/netbird/usr/local/bin/netbird
COPY --from=uname-builder /uname /rootfs/usr/local/lib/containers/netbird/usr/local/bin/uname
COPY netbird/netbird.yaml /rootfs/usr/local/etc/containers/netbird.yaml
