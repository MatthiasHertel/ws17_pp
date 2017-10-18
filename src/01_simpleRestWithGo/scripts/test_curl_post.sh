#! /bin/sh

curl -v -u username:password -H "Content-Type: application/json" -d '{"name":"New Todo "'} http://localhost:8080/todos
