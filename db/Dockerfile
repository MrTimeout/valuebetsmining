ARG CODE_VERSION=latest
FROM postgres:${CODE_VERSION} AS postgreSQL
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are running the ${CODE_VERSION} of postgreSQL"]

ENV POSTGRES_DB viewnext
ENV POSTGRES_PASSWORD_FILE= 

WORKDIR /docker-entrypoint-initdb.d/
COPY ./initDB.sql ./

EXPOSE 5432