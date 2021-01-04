#!/bin/bash

# This script will start up mimicry to server the config.json in this file.
# Config.json has all the urls and mock responses that BSS makes to HSM.

ROOT_DIR="$(dirname $0)"
ROOT_DIR="$(pushd "$ROOT_DIR" > /dev/null && pwd && popd > /dev/null)"

docker run -p 8000:8000 -v ${ROOT_DIR}/config.json:/app/config.json  --env MIMICRY_CONFIG=/app/config.json --env MIMICRY_RELOAD_TIME=0 dtr.dev.cray.com/rbezdicek/mimicry:0.0.1-beta2