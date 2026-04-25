# Stage 1 (Build)
FROM golang:1.26.2-trixie AS builder

ARG VERSION
RUN apt-get update && apt-get install -y git make mailcap
WORKDIR /app/
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/
RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w -X github.com/tonimatasdev/wings/system.Version=$VERSION" \
    -v \
    -trimpath \
    -o wings \
    wings.go
RUN echo "ID=\"distroless\"" > /etc/os-release

# Stage 2 (Final)
FROM gcr.io/distroless/static:latest
COPY --from=builder /etc/os-release /etc/os-release
COPY --from=builder /etc/mime.types /etc/mime.types

COPY --from=builder /app/wings /usr/bin/

ENTRYPOINT ["/usr/bin/wings"]
CMD ["--config", "/etc/pterodactyl/config.yml"]

EXPOSE 8080 2022
