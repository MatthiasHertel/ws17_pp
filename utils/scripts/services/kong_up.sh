#!/bin/bash
# cd utils/kong && docker-compose up -d
# cd -


docker run -d --name kong-database \
              -p 5432:5432 \
              -e "POSTGRES_USER=kong" \
              -e "POSTGRES_DB=kong" \
              postgres:9.4

sleep 10;

docker run --rm \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    kong:latest kong migrations up


docker run -d --name kong \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 8001:8001 \
    -p 8444:8444 \
    kong:latest

sleep 10;

curl -i http://localhost:8001/

curl -i -X POST \
  --url http://localhost:8001/apis/ \
  --data 'name=hpc-rest-api' \
  --data 'hosts=192.168.0.9' \
  --data 'upstream_url=http://192.168.0.9:7777'

curl -i http://localhost:8001/apis/

curl -i -X GET \
  --url http://localhost:8000/ \
  --header 'Host: 192.168.0.9'

docker run -d --name kong-dashboard -p 8080:8080 pgbi/kong-dashboard:v2
