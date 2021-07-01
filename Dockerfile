# MIT License
#
# (C) Copyright [2018-2021] Hewlett Packard Enterprise Development LP
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

# Dockerfile for building HMS bss.

# Build base just has the packages installed we need.
FROM arti.dev.cray.com/baseos-docker-master-local/golang:1.16-alpine3.13 AS build-base

RUN set -ex \
    && apk update \
    && apk add build-base

# Base copies in the files we need to test/build.
FROM build-base AS base

RUN go env -w GO111MODULE=auto

# Copy all the necessary files to the image.
COPY cmd $GOPATH/src/stash.us.cray.com/HMS/hms-bss/cmd
COPY pkg $GOPATH/src/stash.us.cray.com/HMS/hms-bss/pkg
COPY vendor $GOPATH/src/stash.us.cray.com/HMS/hms-bss/vendor
COPY .version $GOPATH/src/stash.us.cray.com/HMS/hms-bss/.version

### UNIT TEST Stage ###
FROM base AS testing

WORKDIR /go

# Run unit tests...
CMD ["sh", "-c", "set -ex && go test -v stash.us.cray.com/HMS/hms-bss/cmd/boot-script-service"]


### COVERAGE Stage ###
FROM base AS coverage

# Run test coverage...
CMD ["sh", "-c", "set -ex && go test -cover -v stash.us.cray.com/HMS/hms-bss/cmd/boot-script-service"]


### Build Stage ###
FROM base AS builder

RUN set -ex && go build -v -i -o /usr/local/bin/boot-script-service stash.us.cray.com/HMS/hms-bss/cmd/boot-script-service

### Final Stage ###
FROM arti.dev.cray.com/baseos-docker-master-local/alpine:3.13
LABEL maintainer="Hewlett Packard Enterprise"
EXPOSE 27778
STOPSIGNAL SIGTERM

# Setup environment variables.
ENV HSM_URL=http://cray-smd
ENV NFD_URL=http://cray-hmnfd

# Set up default path to the Datastore service.
#ENV DATASTORE_URL=https://$ETCD_HOST:$ETCD_PORT
# The datastore is now etcd.  We would like to set the URL to the above, as
# the etcd operator sets up those two env variables.  Unfortunately, env
# vars do not get interpretted in a Dockerfile.  Therefore, bss is set up to
# look for those environment variables.  So we no longer set the DATASTORE_URL
# environment variable.  We still allow it, however, so this setting can be
# controlled from the outside more easily.  Note the special handling below.

# WARNING: Our containers currently do not have certificates set up correctly
#          to allow for https connections to other containers.  Therefore, we
#          will use an insecure connection.  This needs to be corrected before
#          release.  Once the certificates are properly set up, the --insecure
#          option needs to be removed.
ENV BSS_OPTS="--insecure"

ENV BSS_RETRY_DELAY=30
ENV BSS_HSM_RETRIEVAL_DELAY=10
ENV BSS_INIT=/etc/bss.init
#
# Other potentially useful env variables:
# BSS_IPXE_SERVER defaults to "api-gw-service-nmn.local"
# BSS_CHAIN_PROTO defaults to "https"
# BSS_GW_URI defaults to "/apis/bss"

# Include curl in the final image.
RUN set -ex \
    && apk update \
    && apk add --no-cache curl

# Get the boot-script-service from the builder stage.
COPY --from=builder /usr/local/bin/boot-script-service /usr/local/bin/.

COPY .version /

# Set up the command to start the service, the run the init script.
CMD (sleep 4; test -x $BSS_INIT && $BSS_INIT) & boot-script-service $BSS_OPTS --hsm=$HSM_URL ${DATASTORE_URL:+--datastore=}$DATASTORE_URL --retry-delay=$BSS_RETRY_DELAY --hsm-retrieval-delay=$BSS_HSM_RETRIEVAL_DELAY
