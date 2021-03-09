# MongoDB Bookmark Service

## Mongodb Docker

`docker run --name mongo -p 27017:27017 -d mongo:4`

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
