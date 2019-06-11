#!/usr/bin/env bash
# First argument must be true in case of create the image an run a container of it and other values otherwise
# Second argument must be true in case of run the database container
# Third argument must be true in case of run the service and false otherwise

# Executes from the location of this file
if [ $(ls -al | grep -w -c $(basename $0)) -eq 1 ]; then
    echo "Executing $(basename $0)"
else 
    echo "Exiting"
    exit 1
fi

CONTAINER_DATABASE=valuebetsminingmongodb
CONTAINER_GO=valuebetsminingo
CONTAINER_SERVICE=valuebetsmininghttp
IMAGE_DATABASE=mongo
IMAGE_GO=valuebetsmininggo
IMAGE_SERVICE=valuebetsmininghttp
TAG=latest
DATASTORE=/docker-entrypoint-initdb.d
NETWORK=valuebetsmining
PORT_DEFAULT_MONGO=27017
PORT_DEFAULT_GO=8000
PORT_DEFAULT_SERVICE=8080

if [ $(docker network ls --format "{{.Name}}" | grep -w -c $NETWORK) -ge 1 ]; then 
    echo "Network: ${NETWORK} is working"
else 
    echo "Creating network: ${NETWORK}"
    docker network create $NETWORK
fi
if [ "$1" = true ]; then

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
            --volume "${PWD}/data/leagues:/go/src/leagues" \
            --name ${CONTAINER_GO} \
            --publish 3000-4000:${PORT_DEFAULT_GO} \
            --network ${NETWORK} ${IMAGE_GO}:${TAG}
    else 
        echo "Problem creating the image of name ${IMAGE_GO}:${TAG}"
        exit 1
    fi

    if [ "$2" = true ]; then 
        sleep 30s 
    fi    
    
fi

if [ "$2" = true ]; then

    if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then
        if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER_DATABASE) -ge 1 ]; then 
            docker stop $CONTAINER_DATABASE
        fi
        docker container rm $CONTAINER_DATABASE     
    fi

    docker run --detach \
            --env-file secret \
            --volume "${PWD}/data/leagues:/leagues" \
            --volume "${PWD}/data/script:${DATASTORE}" \
            --name ${CONTAINER_DATABASE} \
            --publish 3000-4000:${PORT_DEFAULT_MONGO} \
            --network ${NETWORK} ${IMAGE_DATABASE}:${TAG}
fi

if [ "$3" = true ]; then

    if [ $(docker container ls --all --format "{{.Names}}" | grep -w -c $CONTAINER_SERVICE) -ge 1 ]; then
        if [ $(docker container ls --format "{{.Names}}" | grep -w -c $CONTAINER_SERVICE) -ge 1 ]; then 
            docker stop $CONTAINER_SERVICE
        fi
        docker container rm $CONTAINER_SERVICE
    fi

    if [ $(docker image ls --format "{{.Repository}}" | grep -w -c ${IMAGE_SERVICE}) -eq 1 ]; then
        docker image rm ${IMAGE_SERVICE}:${TAG}
    fi

    docker build --file ${PWD}/src/Dockerfile \
                --tag ${IMAGE_SERVICE}:${TAG} \
                --rm \
                .

    if [ $(docker image ls --format "{{.Repository}}" | grep -w -c ${IMAGE_SERVICE}) -eq 1 ]; then
        IP_ADDR=$(docker container inspect --format '{{.NetworkSettings.Networks.valuebetsmining.IPAddress}}' valuebetsminingmongodb)
        #MONGO_INITDB_PORT=$(docker container inspect valuebetsminingmongodb --format '{{ (index (index .NetworkSettings.Ports "27017/tcp") 0).HostPort }}')
        docker run --detach \
            -e MONGODB_ROOT_ADDR=${IP_ADDR} \
            -e MONGO_INITDB_PORT=${PORT_DEFAULT_MONGO} \
            -e IPADDR=127.0.0.1 \
            -e DNS=valuebetsmining \
            -e PORT=${PORT_DEFAULT_SERVICE} \
            --env-file "${PWD}/secret" \
            --name ${CONTAINER_SERVICE} \
            --volume "${PWD}/src/web:/go/src/valuebetsmining/src/web" \
            --publish 3000-4000:${PORT_DEFAULT_SERVICE} \
            --network ${NETWORK} ${IMAGE_SERVICE}:${TAG}
    else 
        echo "Problem creating the image of name ${IMAGE_SERVICE}:${TAG}"
        exit 1
    fi

fi