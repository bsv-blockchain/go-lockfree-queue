# Example Dockerfile for a Go application (this is just a placeholder)
FROM scratch
COPY go-lockfree-queue /
ENTRYPOINT ["/go-lockfree-queue"]
