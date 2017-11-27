#!/bin/bash

while [ 1 ]
do
  # check if the whole kong stack is running
  if (curl -I -s -L http://localhost:7777 >/dev/null 2>&1); then

    # add job (Job.json)
    printf '\n\e[1;34m%-6s\e[m\n' '-> add job (Job.json)'
    printf 'exec command:\t curl -s -H "Content-Type: application/json" -X POST localhost:7777 -d @Job.json\n'
    if hash jq 2>/dev/null; then
        curl -s -H "Content-Type: application/json" -X POST localhost:7777 -d @Job.json | jq '.'
    else
        curl -s -H "Content-Type: application/json" -X POST localhost:7777 -d @Job.json
    fi

    exit 1
  fi
  printf %s .
done
