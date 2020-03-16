# Used in both stages
ARG OS_VERSION=buster

# STAGE 1
FROM golang:1.13-${OS_VERSION} as builder

# Using this Dockerfile you can build both the web and the cli apps by changing
# the TARGET_APP argument
ARG TARGET_APP=web

# Set the Current Working Directory inside the container
WORKDIR /app

COPY Makefile ./
RUN make test-deps

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY go.mod go.sum ./
RUN make deps

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make prepare-paths
RUN make clean
RUN make "test-${TARGET_APP}" TEST_ARGS="-v -cover"
RUN make "build-url-shortener-${TARGET_APP}"

WORKDIR /app/bin
RUN mv "url-shortener-${TARGET_APP}" url-shortener
RUN stat -c "%n %s" ./*

# STAGE 2
FROM debian:${OS_VERSION}

ENV IS_DOCKER=1

WORKDIR /app/bin
COPY --from=builder /app/bin/url-shortener ./url-shortener
COPY --from=builder /app/config ../config
COPY ./LICENSE ../LICENSE

RUN mkdir -p ../data && \
    chown www-data:www-data ../data && \
    mkdir -p ../log && \
    chown www-data:www-data ../log

EXPOSE 8080

USER www-data

VOLUME [ "/app/data" ]

ENTRYPOINT ["/app/bin/url-shortener"]
CMD [""]
