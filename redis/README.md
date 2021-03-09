# Redis Bookmark Service

## Redis Docker

`docker run --name some-redis -p 6379:6379 -d redis`

## Run service with docker-compose

Start `docker-compose up -d`

Rebuild `docker-compose build`

Stop `docker-compose down -v`

## Manual test

GET bookmarks `curl localhost:8080/bookmark`

Create bookmark 

```
curl --location --request POST 'localhost:8080/bookmark' \
--header 'Content-Type: application/json' \
--data-raw '{
"category":"general",
"name":"YouTube",
"uri":"https://youtube.com"
}'
```

## Check docker-compose logs
All `docker-compose logs`

Individual service


`docker-compose logs bookmark-service`

`docker-compose logs postgres`
