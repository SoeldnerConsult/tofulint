FROM --platform=$BUILDPLATFORM golang:1.25-alpine3.21 AS builder

ARG TARGETOS TARGETARCH

RUN apk add --no-cache make

WORKDIR /tofulint
COPY . /tofulint
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH make build

FROM alpine:3.22

LABEL maintainer=SoeldnerConsult

RUN apk add --no-cache ca-certificates

COPY --from=builder /tofulint/dist/tofulint /usr/local/bin

ENTRYPOINT ["tofulint"]
WORKDIR /data