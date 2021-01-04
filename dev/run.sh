#!/bin/sh

go run stash.us.cray.com/HMS/hms-bss/cmd/boot-script-service --cloud-init-address 0.0.0.0:27778 --datastore http://0.0.0.0:2379 --hsm http://0.0.0.0:8000