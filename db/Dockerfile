ARG CODE_VERSION=latest
FROM postgres:${CODE_VERSION} AS postgreSQL
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are running the ${CODE_VERSION} of postgreSQL"]

WORKDIR /docker-entrypoint-initdb.d/
VOLUME ["./initDB.sql", "$PWD"]
VOLUME ["$HOME/valuebetsmining/secrets/db/postgres-passwd", "$PWD"]
VOLUME ["$HOME/valuebetsmining/db/entrypoint.sh", "$PWD"]

ENV POSTGRES_DB viewnext
ENV POSTGRES_PASSWORD_FILE=$PWD/postgres-passwd

EXPOSE 5432