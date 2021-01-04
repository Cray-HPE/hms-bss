#!/usr/bin/env bash

set -ex
# Setup environment variables
#export GOPATH=$(pwd)/go
RANDY=$(echo $RANDOM | md5sum | awk '{print $1}')


# Parse command line arguments
function usage() {
  echo "$FUNCNAME: $0 [-h] [-k]"
  exit 0
}

while getopts "hk" opt; do
  case $opt in
  h) usage ;;
  *) usage ;;
  esac
done

# Configure docker compose
export COMPOSE_PROJECT_NAME=$RANDY
export COMPOSE_FILE=docker-compose.test.unit.yaml

echo "RANDY: ${RANDY}"
echo "Compose project name: $COMPOSE_PROJECT_NAME"

# It's possible we don't have docker-compose, so if necessary bring our own.
docker_compose_exe=$(command -v docker-compose)
if ! [[ -x "$docker_compose_exe" ]]; then
  if ! [[ -x "./docker-compose" ]]; then
    echo "Getting docker-compose..."
    curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" \
      -o ./docker-compose

    if [[ $? -ne 0 ]]; then
      echo "Failed to fetch docker-compose!"
      exit 1
    fi

    chmod +x docker-compose
  fi
  docker_compose_exe="./docker-compose"
fi

function cleanup() {
  ${docker_compose_exe} down
  if ! [[ $? -eq 0 ]]; then
    echo "Failed to decompose environment!"
    exit 1
  fi
  exit $1
}

# Step 3) Get the base containers running
echo "Starting containers..."
${docker_compose_exe} up  -d --build
network_name=${RANDY}_hms3
docker build --rm --no-cache --network ${network_name} -f Dockerfile.unittesting.Dockerfile .
test_result=$?

# Clean up
echo "Cleaning up containers..."
if [[ $test_result -ne 0 ]]; then
  echo "Unit tests FAILED!"
  cleanup 1
fi

echo "Unit tests PASSED!"
cleanup 0