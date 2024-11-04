FROM golang:1.22.0-alpine AS base

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY internal internal

FROM base AS builder

WORKDIR cmd/grocery-management
RUN go build -o /build/grocery-management .

FROM base AS debug
WORKDIR cmd/grocery-management
RUN go build -gcflags="all=-N -l" -o /build/grocery-management .

ENV DEBUG_PORT=2345
EXPOSE $DEBUG_PORT
ENTRYPOINT /go/bin/dlv --listen=:$DEBUG_PORT --headless=true --api-version=2 --accept-multiclient exec /build/grocery-management

FROM alpine AS run
COPY --from=builder /build/grocery-management /
USER 1000
CMD ["/grocery-management"]

