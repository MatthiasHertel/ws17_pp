#!/bin/bash

cd ./utils/kong/

docker-compose up -d kong-database
docker-compose up -d kong-migration
docker-compose up -d kong
docker-compose up -d consul
docker-compose up -d nginx-lb

cd -
