
#!/usr/bin/env bash

CONTAINER=alpine_project_container
IMAGE=alpine_project_image
TAG=first
NETWORK=golang_postgreSQL

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then 
        docker stop $CONTAINER
    fi
    docker container rm $CONTAINER     
fi

docker build --tag $IMAGE:$TAG --rm .

if [ $(docker image ls --all --filter "reference=$IMAGE:$TAG" --format "{{.Repository}}:{{.Tag}}" | grep -w -c $IMAGE:$TAG) -eq 1 ]; then
    docker run -d --name $CONTAINER --env-file ./src/secrets/setEnviromentVariables --network $NETWORK -p 3000-4000:80 $IMAGE:$TAG
fi