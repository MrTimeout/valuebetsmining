#!/bin/sh

for f in $DEFAULT_DIR_FILES/*.csv
do
    mongoimport --username $MONGO_INITDB_ROOT_USERNAME \
                --password $MONGO_INITDB_ROOT_PASSWORD \
                --authenticationDatabase $MONGO_INITDB_DATABASE \
                --db $MONGO_DB_NAME \
                --collection $(basename -s.csv $f) \
                --type=csv \
                --headerline \
                --file $f
done
