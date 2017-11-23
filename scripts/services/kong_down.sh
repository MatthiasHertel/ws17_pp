#!/bin/bash

docker stop kong && docker rm kong
docker stop kong-database && docker rm kong-database
docker stop kong-dashboard && docker rm kong-dashboard
