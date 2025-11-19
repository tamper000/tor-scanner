# Stage 1
FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /src

COPY go.mod go.sum ./
COPY cmd/ cmd/
COPY internal/ internal/

RUN go mod download \
    && CGO_ENABLED=0 go build -o tor-scanner ./cmd/main.go

# Stage 2
FROM alpine:3.22

RUN apk add --no-cache \
    ca-certificates \
    tor \
    && mkdir -p /etc/tor \
       /var/lib/tor \
       /var/run/tor \
    && chown -R tor:tor /etc/tor /var/lib/tor /var/run/tor

COPY --from=builder /src/tor-scanner /usr/local/bin/tor-scanner

COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN chown tor:tor /usr/local/bin/tor-scanner /entrypoint.sh

USER tor

WORKDIR /home/tor

EXPOSE 9050

ENTRYPOINT ["/entrypoint.sh"]
CMD []
