# MIT License
#
# (C) Copyright [2021] Hewlett Packard Enterprise Development LP
#
# Permission is hereby granted, free of charge, to any person obtaining a
# copy of this software and associated documentation files (the "Software"),
# to deal in the Software without restriction, including without limitation
# the rights to use, copy, modify, merge, publish, distribute, sublicense,
# and/or sell copies of the Software, and to permit persons to whom the
# Software is furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included
# in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
# THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
# OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
# ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.

# Dockerfile for running bss unit tests.

### build-base stage ###
# Build base just has the packages installed we need.
FROM artifactory.algol60.net/docker.io/library/golang:1.16-alpine AS build-base

RUN set -ex \
    && apk -U upgrade \
    && apk add build-base

### base stage ###
# Base copies in the files we need to test/build.
FROM build-base AS base

RUN go env -w GO111MODULE=auto

# Copy all the necessary files to the image.
COPY cmd $GOPATH/src/github.com/Cray-HPE/hms-bss/cmd
COPY pkg $GOPATH/src/github.com/Cray-HPE/hms-bss/pkg
COPY vendor $GOPATH/src/github.com/Cray-HPE/hms-bss/vendor
COPY .version $GOPATH/src/github.com/Cray-HPE/hms-bss/.version

### testing stage ###
FROM base AS testing

WORKDIR /go

# Run unit tests.
CMD ["sh", "-c", "go test -cover -v github.com/Cray-HPE/hms-bss/cmd/boot-script-service"]
