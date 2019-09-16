FROM golang:1.12.9-buster

RUN mkdir -p /var/projects/curve-tech-test
COPY . /var/projects/curve-tech-test

WORKDIR /var/projects/curve-tech-test

RUN cp .env.dist .env

RUN go test ./...

RUN apt-get update -y && apt-get install -y \
    curl

RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

RUN ./bin/golangci-lint run \
    --enable=golint \
    --enable=gofmt \
    --enable=gosec  \
    --enable=misspell \
    --enable=maligned \
    --enable=interfacer \
    --enable=stylecheck  \
    --enable=unconvert \
    --enable=maligned \
    --enable=goconst \
    --enable=nakedret

RUN go build -o serve application/serve.go

EXPOSE 8091

ARG EXCHANGE_API_URI
ENV EXCHANGE_API_URI=$EXCHANGE_API_URI

ENTRYPOINT ["/var/projects/curve-tech-test/serve"]

HEALTHCHECK \
  --interval=30s \
  --timeout=10s \
  CMD curl -f http://localhost:8091/v1/health || exit 1