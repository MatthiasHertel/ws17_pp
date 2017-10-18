#! /bin/sh

curl -v -H "Content-Type: application/json" -d '{"name":"New Todo "' http://localhost:8080/todos
