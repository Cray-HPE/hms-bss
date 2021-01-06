#!/bin/bash

NODE_IP="0.0.0.0"

TAG="v3.4.10"

docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  quay.io/coreos/etcd:${TAG} \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE_IP}:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${NODE_IP}:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${NODE_IP}:2380