package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"valuebetsmining/src/gql"
	"valuebetsmining/src/postgres"
	"valuebetsmining/src/server"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {

	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":80", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {

	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString(os.Getenv("IP"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("DBNAME"), os.Getenv("POSTGRES_PASSWD")),
	)
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,                             // log api request calls
		middleware.DefaultCompress,                    // compress results, mostly gzipping assets and json
		middleware.StripSlashes,                       // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,                          // recover from panics without crashing server
	)

	router.Post("/users", s.GraphQL()) //curl -d '{  "query": "{users (name:\"kevin\"){id, name, age}}" }' http://localhost:3000/users

	return router, db
}
