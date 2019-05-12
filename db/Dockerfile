ARG CODE_VERSION=latest
FROM postgres:${CODE_VERSION} AS postgreSQL
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are running the ${CODE_VERSION} of postgreSQL"]

ENV POSTGRES_DB viewnext
ENV POSTGRES_USER root

WORKDIR /docker-entrypoint-initdb.d/
COPY ./initDB.sql /docker-entrypoint-initdb.d/
COPY ./secrets/db/postgres-passwd /docker-entrypoint-initdb.d/secrets/

ENV POSTGRES_PASSWORD_FILE=/docker-entrypoint-initdb.d/secrets/postgres-passwd

EXPOSE 5432