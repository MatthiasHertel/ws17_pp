#!/bin/bash

# if [ -z "$1" ]; then
#     echo '$1 must be the local external ip of this host'
#     exit 1
# fi
LOCALHOST_IP=$(ifconfig eth0 | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
if [ -z "$LOCALHOST_IP" ]; then
  LOCALHOST_IP=$(ifconfig | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
fi
set -- $LOCALHOST_IP

while [ 1 ]
do
  # check if the whole kong stack is running
  if (docker inspect -f {{.State.Running}} kong_nginx-lb_1 >/dev/null 2>&1 ) && (curl -I -s -L http://localhost:8001 >/dev/null 2>&1); then

    printf '\e[1;31m%-6s\e[0m\n' '-> start seeding consumer'

    # add consumer to kong without acl
    printf '\n\e[1;34m%-6s\e[m\n' '-> add consumer to kong without acl'
    printf 'exec command:\t curl -X POST --url http://localhost:8001/consumers/ --data "username=john" \n'
    if hash jq 2>/dev/null; then
        curl -s -X POST --url http://localhost:8001/consumers/ --data "username=john" | jq '.'
    else
        curl -s -X POST --url http://localhost:8001/consumers/ --data "username=john"
    fi

    # add api-key to consumer
    printf '\n\e[1;34m%-6s\e[m\n' '-> add api-key to consumer'
    printf 'exec command:\t curl -X POST --url http://localhost:8001/consumers/john/key-auth/ --data "key=123" \n'
    if hash jq 2>/dev/null; then
        curl -s -X POST --url http://localhost:8001/consumers/john/key-auth/ --data 'key=123' | jq '.'
    else
        curl -s -X POST --url http://localhost:8001/consumers/john/key-auth/ --data 'key=123'
    fi

    # Test - 200 forward to hpc-rest-api
    printf '\n\e[1;34m%-6s\e[m\n' '-> Test assert - 200 forward to hpc-rest-api'
    printf 'exec command:\t curl -X GET --url http://localhost:8000 --header "Host: '${1}'" --header "apikey: 123" \n'
    curl -s -X GET --url http://localhost:8000 --header "Host: ${1}" --header "apikey: 123"

    exit 1
  fi
  # printf %s .
done
