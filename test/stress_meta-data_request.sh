#!/bin/bash

check_url() {
    url="$1"
    # Use curl to capture the HTTP status code
    http_code=$(curl -s -o /dev/null -w "%{http_code}" "$url")
    if [ "$http_code" != "200" ]; then
        echo "Process $$: Error - $url returned HTTP code $http_code"
        return 1
    fi
    return 0
}

BOOTS=1
NODES=5000

for ((i=1; i<=BOOTS; i++)); do
    pids=()
    numpids=0

    # Background each pair of requests, mimicing compute nodes booting in parallel

    echo "Issuing $NODES /meta-data requests for boot attempt $i"

    for ((j=1; j<=NODES; j++)); do
         # Request that includes client IP

         check_url "http://10.92.100.81:8888/meta-data" &
         pids+=($!)

         # Request that includes weave IP

         check_url "http://api-gw-service-nmn.local:8888/meta-data" &
         pids+=($!)

         numpids=$((numpids + 2))
    done

    # Wait for all requests to complete

    echo "Watiting for $numpids responses to complete"
    sleep 1

    for pid in "${pids[@]}"; do
         wait "$pid"

         numpids=$((numpids - 1))

         echo "$numpids responses remaining"
    done
done

echo "$BOOTS boot attempts complete"