FROM golang:alpine AS builder

WORKDIR /workspace

# Copy in just the go.mod and go.sum files, and download the dependencies. By
# doing this before copying in the other dependencies, the Docker build cache
# can skip these steps so long as neither of these two files change.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest
COPY . .

# Build the Go app with CGO_ENABLED=0 so we use the pure-Go implementations for
# things like DNS resolution (so we don't build a binary that depends on system
# libraries)
RUN CGO_ENABLED=0 go build -o /app

# Create the 'nobody' user and group files that will be used in the running container to
# run the process an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# The final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app /app

# Copy the /etc/passwd file we created in the builder stage. This creates a new
# non-root user as a security best practice.
COPY --from=builder /user/group /user/passwd /etc/

# Run as the new non-root by default
USER nobody:nobody

# Run the binary
ENTRYPOINT [ "/app" ]