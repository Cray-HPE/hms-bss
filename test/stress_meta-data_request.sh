#!/bin/bash

# MIT License
#
# (C) Copyright [2025] Hewlett Packard Enterprise Development LP
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

# Simple developer script to test high bandwidth /meta-data requests
# Is not part of any automated testing

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

BOOTS=100
NODES=500

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

         check_url "http://api-gw-service-nmn.local:8888/meta-data?key=Global" &
         pids+=($!)

         numpids=$((numpids + 2))
    done

    # Wait for all requests to complete

    echo "Watiting for $numpids responses to complete"

    for pid in "${pids[@]}"; do
         wait "$pid"

         numpids=$((numpids - 1))

         echo "$numpids responses remaining"
    done
done

echo "$BOOTS boot attempts complete"
