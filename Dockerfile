FROM golang:1.18 as module-downloader
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM module-downloader as builder
ARG BUILD_TARGET_PATH
COPY . .
RUN go build -ldflags="-w -s" -o /go/bin/app ${BUILD_TARGET_PATH}

FROM gcr.io/distroless/base@sha256:2eb6d15c45d0d35ea270a53b49e93b5f5299d2086a9bcad26da077fa9ed549c5 as application
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]

FROM alpine:3.16.0 as binary-downloader
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.11 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe

FROM application as grpc-server
COPY --from=binary-downloader /bin/grpc_health_probe /bin/grpc_health_probe
