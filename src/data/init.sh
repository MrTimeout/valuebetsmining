#!/usr/bin/env bash

# Executes from the location of this file
if [ $(ls -al | grep -w -c $(basename $0)) -eq 1 ]; then
    echo "Executing $(basename $0)"
else 
    echo "Exiting"
    exit 1
fi

CONTAINER_DATABASE=valuebetsminingmongodb
CONTAINER_GO=valuebetsminingo
IMAGE_DATABASE=mongo
IMAGE_GO=valuebetsmininggo
TAG=latest
DATASTORE=/docker-entrypoint-initdb.d
NETWORK=valuebetsmining
PORT_DEFAULT_MONGO=27017

if [ $(docker network ls --format "{{.Name}}" | grep -w -c $NETWORK) -ge 1 ]; then 
    echo "Network: ${NETWORK} is working"
else 
    echo "Creating network: ${NETWORK}"
    docker network create $NETWORK
fi

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_GO) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER_GO) -ge 1 ]; then 
        docker stop $CONTAINER_GO
    fi
    docker container rm $CONTAINER_GO
fi

if [ $(docker image ls --format "{{.Repository}}" | grep -w -c ${IMAGE_GO}) -eq 1 ]; then
    docker image rm ${IMAGE_GO}:${TAG}
fi

docker build --file ${PWD}/Dockerfile \
             --tag ${IMAGE_GO}:${TAG} \
             --rm \
             .

if [ $(docker image ls --format "{{.Repository}}" | grep -w -c ${IMAGE_GO}) -eq 1 ]; then
    docker run --detach \
           --volume "${PWD}/src/leagues:/go/src/leagues" \
           --name ${CONTAINER_GO} \
           --publish 3000-4000:8080 \
           --workdir /go/src \
           --network ${NETWORK} ${IMAGE_GO}:${TAG}
else 
    echo "Problem creating the image of name ${IMAGE_GO}:${TAG}"
    exit 1
fi

if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then
    if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then 
        docker stop $CONTAINER_DATABASE
    fi
    docker container rm $CONTAINER_DATABASE     
fi

docker run --detach \
           --env-file secret \
           --volume "${PWD}/src/leagues:/leagues" \
           --volume "${PWD}/src/script:${DATASTORE}" \
           --name ${CONTAINER_DATABASE} \
           --publish 3000-4000:${PORT_DEFAULT_MONGO} \
           --network ${NETWORK} ${IMAGE_DATABASE}:${TAG}
