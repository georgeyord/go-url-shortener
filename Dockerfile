FROM golang:1.13.0

# Set the Current Working Directory inside the container
WORKDIR /app

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

# Command to run the executable
CMD ["./bin/helloworld"]
