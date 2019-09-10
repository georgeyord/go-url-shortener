# Used in both stages
ARG OS_VERSION=buster

FROM golang:1.13-${OS_VERSION} as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY Makefile ./
RUN make test_deps

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY go.mod go.sum ./
RUN make deps

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make test
RUN make build-helloworld-cmd

FROM debian:${OS_VERSION}

WORKDIR /app/bin
COPY --from=builder /app/bin/* .

# Command to run the executable
CMD ["/app/bin/helloworld"]
