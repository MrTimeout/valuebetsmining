package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/estenoesmiputonombre/valuebetsmining/tree/testingGraphQL/src/gql"
	"github.com/estenoesmiputonombre/valuebetsmining/tree/testingGraphQL/src/postgres"
	"github.com/estenoesmiputonombre/valuebetsmining/tree/testingGraphQL/src/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {

	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {

	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "root", "viewnext"),
	)
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}