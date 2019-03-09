FROM golang as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/steren/memegen

# Download dependencies (for faster builds when deps do not change)
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy local code
COPY . .

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o memegen

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Install fonts
RUN apk --no-cache add msttcorefonts-installer fontconfig && \
    update-ms-fonts && \
    fc-cache -f

# Install certificates for HTTPS
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/steren/memegen/memegen /memegen

# Service must listen to $PORT environment variable.
# This default value facilitates local development.
ENV PORT 8080

# Run the web service on container startup.
CMD ["/memegen"]