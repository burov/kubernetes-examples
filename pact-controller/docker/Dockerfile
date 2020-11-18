FROM golang:1.15.5 AS builder

WORKDIR /go/src/github.com/burov/kubernetes-examples/pact-controller/

# Hack to cache dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source files
COPY . .

# Copy binary
RUN  CGO_ENABLED=0 GOOS=linux go build -o bin/controller ./cmd/controller

FROM alpine:latest AS runner

WORKDIR /bin

COPY --from=builder /go/src/github.com/burov/kubernetes-examples/pact-controller/bin/controller .

ENTRYPOINT ["controller"]



