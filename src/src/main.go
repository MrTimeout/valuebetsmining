package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	//PORT ... Port of the golang program
	PORT = os.Getenv("PORT")
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
