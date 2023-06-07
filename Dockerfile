FROM cgr.dev/chainguard/go:latest as builder

WORKDIR /src
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download -x

# Copy the go source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=./glab-runner-cleaner .

FROM cgr.dev/chainguard/static:latest

WORKDIR /app
COPY --from=builder /src/glab-runner-cleaner .
USER 65532:65532

CMD ["./glab-runner-cleaner"]
