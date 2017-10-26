#!/bin/bash


docker build --tag restapi .

docker run -it restapi
