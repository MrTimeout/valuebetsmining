#!/usr/bin/env bash

CONTAINER=postgreSQL_project_container
IMAGE=postgresql_project_image
TAG=first
NETWORK=golang_postgreSQL

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then 
        docker stop $CONTAINER
    fi
    docker container rm $CONTAINER     
fi

docker build --tag $IMAGE:$TAG --no-cache --rm .

if [ $(docker image ls --all --filter "reference=$IMAGE:$TAG" --format "{{.Repository}}:{{.Tag}}" | grep -w -c $IMAGE:$TAG) -eq 1 ]; then
    docker run -d --name $CONTAINER --network $NETWORK -p 5000-6000:5432 $IMAGE:$TAG
fi