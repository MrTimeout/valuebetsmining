#!/bin/sh

ls -al /leagues/CSV

for f in $DEFAULT_DIR_FILES/*.csv
do
    echo $(basename -s.csv $f)
    mongoimport --username $MONGO_INITDB_ROOT_USERNAME \
                --password $MONGO_INITDB_ROOT_PASSWORD \
                --authenticationDatabase $MONGO_INITDB_DATABASE \
                --db $MONGO_DB_NAME \
                --collection $(basename -s.csv $f) \
                --type=csv \
                --headerline \
                --file $f
done

# mongo -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --authenticationDatabase $MONGO_INITDB_DATABASE

# db.getCollection("SpainSP11019").distinct( 'LocalTeam', { Date: { $regex: /^[0-9]{2}\/[0-9]{2}\/(18|19)$/ }, Index: { $gte: 1 } } )

# [A-Z]{1,2}([a-z]{1,}|[0-9]?)