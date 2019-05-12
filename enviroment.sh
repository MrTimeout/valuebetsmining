#!/usr/bin/env bash

mkdir gql postgres server
touch gql/gql.go gql/queries.go gql/resolvers.go gql/types.go \
      postgres/postgres.go \
      server/server.go \
      main.go

go get github.com/oxequa/realize
go get github.com/go-chi/chi
go get github.com/go-chi/render
go get github.com/graphql-go/graphql
go get github.com/lib/pq