#!/bin/sh

#mongo admin --eval "db.createUser({user: '$MONGO_INITDB_ROOT_USERNAME', pwd: '$MONGO_INITDB_ROOT_PASSWORD', roles:[{role:'readWrite',db:'$MONGO_INITDB_DATABASE'}]});"
mongoimport --username $MONGO_INITDB_ROOT_USERNAME \
            --password $MONGO_INITDB_ROOT_PASSWORD \
            --authenticationDatabase $MONGO_INITDB_DATABASE \
            --db leagues \
            --collection spain \
            --type=csv \
            --headerline \
            --file /leagues/Spain/SP1.csv