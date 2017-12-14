#!/bin/bash

curl -X PUT -d @fib-ex.nomad.json http://127.0.0.1:4646/v1/jobs 
