#!/usr/bin/env bash

IMAGE_NAME="cray-bss"

usage() {
        echo "$FUNCNAME: $0"
        echo "  -h | prints this help message" 
        echo "  -l | hostname to push to, default localhost";
        echo "  -r | repo to push to, default cray";
        echo "  -f | forces build with --no-cache and --pull";
	echo "";
        exit 0
}


REPO="cray"
REGISTRY_HOSTNAME="localhost"
FORCE=" "

while getopts "hfl:r:" opt; do
  case $opt in
    h) 
      usage; 
      exit;;
    f) 
      FORCE=" --no-cache --pull";;
    l) 
      REGISTRY_HOSTNAME=${OPTARG};;
    r) 
      REPO=${OPTARG};;
  esac
done

shift $((OPTIND-1))

echo "Building $FORCE and pushing to $REGISTRY_HOSTNAME in repo $REPO"

set -e
docker build $FORCE -t cray/$IMAGE_NAME .
docker tag cray/$IMAGE_NAME $REGISTRY_HOSTNAME/$REPO/$IMAGE_NAME
docker push $REGISTRY_HOSTNAME/$REPO/$IMAGE_NAME
