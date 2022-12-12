#!/bin/bash

alias docker='sudo docker'

# Clear old images and containers if needed
# this may return an error message if the container
# or image is not present, but it still works
docker kill db
docker kill api
docker rm db
docker rm api
docker rmi muzz-db
docker rmi muzz-api

docker build -t muzz-db db/
docker run -d --name db muzz-db

export DB_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' db)

docker build -t muzz-api .
docker run -d --name api -e DB_IP=$DB_IP muzz-api

# can be used to test with curl
API_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' api)

# Run tests
docker exec -t api /bin/sh -c "cd handlers && go test"
