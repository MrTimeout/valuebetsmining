package main

import (
	"log"
	"valuebetsmining/src/server"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.NotFoundHandler = server.Error404()
	server.StablishHanlders(rtr)
	srv := server.NewConfigServer(rtr)
	log.Fatal(srv.ListenAndServe())
}
