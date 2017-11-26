#!/bin/bash

cd "$(dirname "$0")"
cd ..

# TODO stop all started services and running docker containers here

# stop kong container (kong , kong-database, kong-dashboard)
./scripts/services/kong_compose_down.sh
./scripts/services/minio_compose_down.sh
# ./scripts/services/consul_down.sh
# ./scripts/services/nomad_down.sh

sleep 2;

tmux kill-session -t hpc-rest-dev-env
