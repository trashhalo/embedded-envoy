FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app

FROM envoyproxy/envoy-alpine:v1.14-latest
EXPOSE 8080
ARG S6_OVERLAY_RELEASE=https://github.com/just-containers/s6-overlay/releases/latest/download/s6-overlay-amd64.tar.gz
ENV S6_OVERLAY_RELEASE=${S6_OVERLAY_RELEASE}
ADD ${S6_OVERLAY_RELEASE} /tmp/s6overlay.tar.gz
RUN apk upgrade --update --no-cache \
    && rm -rf /var/cache/apk/* \
    && tar xzf /tmp/s6overlay.tar.gz -C / \
    && rm /tmp/s6overlay.tar.gz
COPY --from=builder /app/app /app
COPY envoy.yaml /etc/envoy/envoy.yaml
COPY s6 /etc/services.d
ENV S6_BEHAVIOUR_IF_STAGE2_FAILS=2
CMD ["/init"]