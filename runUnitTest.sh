#!/usr/bin/env bash
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

set -ex

GITSHA=$(git rev-parse HEAD)
TIMESTAMP=$(date +"%Y-%m-%dT%H-%M-%SZ")
IMAGE="cray/hms-bss-coverage"
# image names must be lower case
UNIQUE_TAG=$(echo ${IMAGE}_${GITSHA}_${TIMESTAMP} | tr '[:upper:]' '[:lower:]')
# export NO_CACHE=--no-cache # this will cause docker build to run with no cache; off by default for local builds, enabled in jenkinsfile

DOCKER_BUILDKIT=0 docker build $NO_CACHE -t $UNIQUE_TAG -f Dockerfile.testing .
docker image rm $UNIQUE_TAG --force

