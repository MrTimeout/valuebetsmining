
#!/usr/bin/env bash

CONTAINER=golang_project_container
IMAGE=golang_project_image
TAG=first
FILE=src/run.sh

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER) -ge 1 ]; then 
        docker stop $CONTAINER
    fi
    docker container rm $CONTAINER     
fi

if [ -f "$FILE" ]; then
    chmod +x $FILE
fi


docker build --tag $IMAGE:$TAG --no-cache --rm .

if [ $(docker image ls --all --filter "reference=$IMAGE:$TAG" --format "{{.Repository}}:{{.Tag}}" | grep -w -c $IMAGE:$TAG) -eq 1 ]; then
    docker run -d --name $CONTAINER --network net -p 3000-4000:8080 $IMAGE:$TAG
fi