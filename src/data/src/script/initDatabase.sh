#!/bin/sh

mongoimport --username ${MONGO_INITDB_ROOT_USERNAME} --password ${MONGO_INITDB_ROOT_PASSWORD} --db ${MONGO_INITDB_DATABASE} --collection spain --type=csv --headerline --file /leagues/Spain/SP1.csv