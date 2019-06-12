#!/bin/sh

for f in $DEFAULT_DIR_FILES*.csv
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

# db.getCollection("SpainSP11019").find({"LocalTeam": "Ath Madrid", "Date": { $regex: /^[0-9]{2}\/[0-9]{2}\/(18|19)$/ } }).sort({"Index":-1}).limit(1).pretty()