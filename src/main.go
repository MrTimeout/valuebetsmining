package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	//PORT ... Port of the golang program
	PORT = os.Getenv("PORT")
	//IPADDR ... ip addr to connect
	IPADDR = os.Getenv("IPADDR")
	//DNS ... Name of the host
	DNS = "valuebetsmining"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/inicio", func(w http.ResponseWriter, r *http.Request) {

	}).Host(DNS).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", IPADDR, PORT),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
