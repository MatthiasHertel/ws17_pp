# Install Docker Container

https://github.com/Kong/docker-kong/tree/master/compose

1. start your database

```
docker run -d --name kong-database \
              -p 5432:5432 \
              -e "POSTGRES_USER=kong" \
              -e "POSTGRES_DB=kong" \
              postgres:9.4
```

2. prepare your database

```
docker run --rm \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    kong:latest kong migrations up
```

3. Start Kong

```
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
```

4. Use kong

```
curl -i http://localhost:8001/
```


# Adding your API

1. Add your API using the Admin API

```
curl -i -X POST \
  --url http://localhost:8001/apis/ \
  --data 'name=hpc-rest-api' \
  --data 'hosts=192.168.247.105' \
  --data 'upstream_url=http://192.168.247.105:8080'
```

NOTE:
upstream url ... use the not localhost cause of using localhost inside docker-container

2. Verify that your API has been added

```
HTTP/1.1 201 Created
Content-Type: application/json
Connection: keep-alive

{
  "created_at": 1488830759000,
  "hosts": [
      "192.168.247.105"
  ],
  "http_if_terminated": true,
  "https_only": false,
  "id": "6378122c-a0a1-438d-a5c6-efabae9fb969",
  "name": "hpc-rest-api",
  "preserve_host": false,
  "retries": 5,
  "strip_uri": true,
  "upstream_connect_timeout": 60000,
  "upstream_read_timeout": 60000,
  "upstream_send_timeout": 60000,
  "upstream_url": "http://mockbin.org"
}
```

3. Forward your requests through Kong

```
curl -i -X GET \
  --url http://localhost:8000/ \
  --header 'Host: 192.168.247.105'
```

4. Enabling Plugin (key-auth) and Add Consumer

```
 curl -i -X POST --url http://localhost:8001/apis/hpc-rest-api/plugins/ --data 'name=key-auth'

 curl -i -X GET --url http://localhost:8000/ --header 'Host: 192.168.247.105'

 curl -i -X POST --url http://localhost:8001/consumers/ --data "username=john"

 curl -i -X POST --url http://localhost:8001/consumers/Jason/key-auth/ --data 'key=1900'

 curl -i -X GET --url http://localhost:8000 --header "Host: 192.168.247.105" --header "apikey: 1900"
```
