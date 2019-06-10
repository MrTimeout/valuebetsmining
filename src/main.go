package main

import (
	"fmt"
	"io/ioutil"
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
	DNS = os.Getenv("DNS")
	//DefaultDirWEB ... Default dir to follow web
	DefaultDirWEB = "./web"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(DefaultDirWEB))))
	r.HandleFunc("/inicio", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("web/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "<h1>Hey</h1>")
	})
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", PORT),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
