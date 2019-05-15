ARG CODE_VERSION=latest
FROM mongo:${CODE_VERSION} AS mongo
ARG CODE_VERSION

ENV MONGO_INITDB_ROOT_USERNAME something
ENV MONGO_INITDB_ROOT_PASSWORD something
ENV MONGO_INITDB_DATABASE something

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of mongo"]

WORKDIR /docker-entrypoint-initdb.d

EXPOSE 27017
