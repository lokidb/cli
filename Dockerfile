# Build the executables
FROM golang:1.18-bullseye as base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# Download requirements
RUN go mod download
RUN go mod verify

# Copy all files
COPY ./ ./


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /lokidb-cli cmd/cli.go

# Use executables
FROM gcr.io/distroless/base-debian11

COPY --from=base /lokidb-cli /

ENTRYPOINT [ "/lokidb-cli" ]
