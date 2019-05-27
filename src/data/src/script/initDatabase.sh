#!/bin/sh

mongoimport --username $MONGO_INITDB_ROOT_USERNAME \
            --password $MONGO_INITDB_ROOT_PASSWORD \
            --authenticationDatabase $MONGO_INITDB_DATABASE \
            --db leagues \
            --collection spain \
            --type=csv \
            --headerline \
            --file /leagues/Spain/SP1.csv