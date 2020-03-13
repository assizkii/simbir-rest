# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.14 as builder

# Add Maintainer Info
LABEL maintainer="Nickolay Kisluhin <n-kis@bitrix24.ru>"

ENV CGO_ENABLED=0

# Set the Current Working Directory inside the container
WORKDIR /go/bin
# Copy files
COPY ./ ./
# Build the Go app
RUN go get -d -v && go build -o rest ./...

# STEP 2 build a small image
# start from scratch
FROM alpine:latest
# Copy our static executable
COPY --from=builder /go/bin/rest /go/bin/rest
ENTRYPOINT ["/go/bin/rest"]
