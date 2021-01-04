# Copyright 2020 Hewlett Packard Enterprise Development LP

# Dockerfile for testing HMS s3 code.

FROM dtr.dev.cray.com/baseos/golang:1.14-alpine3.12 AS build-base

RUN set -ex \
    && apk update \
    && apk add build-base

# Copy the files in for the next stages to use.
FROM build-base

COPY *.go $GOPATH/src/stash.us.cray.com/HMS/hms-s3/
COPY vendor $GOPATH/src/stash.us.cray.com/HMS/hms-s3/vendor

ENV LOG_LEVEL "TRACE"

# if you do CMD, then it will run like a service; we want this to run the tests and quit
RUN set -ex \
    && go test -cover -v -o hms-s3 stash.us.cray.com/HMS/hms-s3/...
