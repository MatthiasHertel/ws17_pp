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

    printf '\e[1;31m%-6s\e[0m\n' '-> start seeding plugin (key-auth)'

    # add plugin key-auth for the api 'hpc-rest-api'
    printf '\n\e[1;34m%-6s\e[m\n' '-> seed plugin key-auth for the api "hpc-rest-api"'
    printf 'exec command:\t curl -i -X POST --url http://localhost:8001/apis/hpc-rest-api/plugins/ --data "name=key-auth"'
    if hash jq 2>/dev/null; then
        curl -s -X POST --url http://localhost:8001/apis/hpc-rest-api/plugins/ --data 'name=key-auth' | jq
    else
        curl -s -X POST --url http://localhost:8001/apis/hpc-rest-api/plugins/ --data 'name=key-auth'
    fi


    # Test - 401 not authorized (no api key found)
    printf '\n\e[1;34m%-6s\e[m\n' '-> Test - 401 not authorized (no api key found)'
    printf 'exec command:\t curl -i -X GET --url http://localhost:8000/ --header "Host:'${1}' "\n'
    if hash jq 2>/dev/null; then
        curl -s -X GET --url http://localhost:8000/ --header 'Host: '${1} | jq
    else
        curl -s -X GET --url http://localhost:8000/ --header 'Host: '${1}
    fi

    exit 1
  fi
  # printf %s .
done
