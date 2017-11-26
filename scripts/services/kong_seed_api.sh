#!/bin/bash

# if [ -z "$1" ]; then
#     echo '$1 must be the local external ip of this host'
#     exit 1
# fi
LOCALHOST_IP=$(ifconfig docker0 | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
set -- $LOCALHOST_IP

while [ 1 ]
do
  # check if the whole kong stack is running
  if (docker inspect -f {{.State.Running}} kong_nginx-lb_1 >/dev/null 2>&1 ) && (curl -I -s -L http://localhost:8001 >/dev/null 2>&1); then

    printf "\n\e[1;31m%-6s\e[0m\n" "-> start seeding api"

    # add api with the name "hpc-rest-api" to kong
    printf '\n\e[1;34m%-6s\e[m\n' '-> add api with the name "hpc-rest-api" to kong'
    printf 'exec command:\t curl -X POST --url http://localhost:8001/apis/ --data "name=hpc-rest-api" --data "hosts='${1}'" --data "upstream_url=http://'${1}':7777"\n'
    if hash jq 2>/dev/null; then
        curl -s -X POST \
        --url http://localhost:8001/apis/ \
        --data 'name=hpc-rest-api' \
        --data 'hosts='${1} \
        --data 'upstream_url=http://'${1}':7777' | jq
    else
        curl -s -X POST \
        --url http://localhost:8001/apis/ \
        --data 'name=hpc-rest-api' \
        --data 'hosts='${1} \
        --data 'upstream_url=http://'${1}':7777'
    fi

    # Test - 200 list the apis
    printf '\n\e[1;34m%-6s\e[m\n' '-> Test assert - 200 list api'
    printf 'exec command:\t curl http://localhost:8001/apis/hpc-rest-api \n'
    if hash jq 2>/dev/null; then
        curl -s http://localhost:8001/apis/hpc-rest-api | jq
    else
        curl -s http://localhost:8001/apis/hpc-rest-api
    fi


    exit 1
  fi
  printf %s .
done
