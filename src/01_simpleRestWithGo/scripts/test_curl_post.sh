#! /bin/sh

curl -H "Content-Type: application/json" -d '{"name":"New Todo '$1'"}' http://localhost:8080/todos
