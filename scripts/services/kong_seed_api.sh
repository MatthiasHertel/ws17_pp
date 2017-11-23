#!/bin/bash

# if [ -z "$1" ]; then
#     echo '$1 must be the local external ip of this host'
#     exit 1
# fi
LOCALHOST_IP=$(ifconfig | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p')
set -- $LOCALHOST_IP

while [ 1 ]
do
  if (docker inspect -f {{.State.Running}} kong_nginx-lb_1 >/dev/null 2>&1 ) ; then

  echo 'start seeding'

  sleep 5;

  # curl -i http://localhost:8001/

  curl -i -X POST \
    --url http://localhost:8001/apis/ \
    --data 'name=hpc-rest-api' \
    --data 'hosts='${1} \
    --data 'upstream_url=http://'${1}':7777'

  curl -i http://localhost:8001/apis/

  # curl -i -X GET \
  #   --url http://localhost:8000/ \
  #   --header 'Host: '${1}

  exit 1
  fi
  printf %s .
done
