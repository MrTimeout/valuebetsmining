#!/usr/bin/env bash

CONTAINER=postgreSQL_project_container
IMAGE=postgreSQL_project_image
TAG=first
NETWORK=

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then 
        docker stop $CONTAINER
    fi
    docker container rm $CONTAINER     
fi

docker build --tag $IMAGE:$TAG --no-cache --rm .

if [ $(docker image ls --all --filter "reference=$IMAGE:$TAG" --format "{{.Repository}}:{{.Tag}}" | grep -w -c $IMAGE:$TAG) -eq 1 ]; then
    docker run -d --name $CONTAINER --network  -p 3000-4000:8080 $IMAGE:$TAG
fi