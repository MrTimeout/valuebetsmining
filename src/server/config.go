package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	//PORT ... Port of the golang program
	PORT = "8080"
	//IPADDR ... ip addr to connect
	IPADDR = "127.0.0.1"
	//DNS ... Name of the host
	DNS = os.Getenv("DNS")
)

//NewConfigServer ... Creates a new config server
func NewConfigServer(rtr *mux.Router) *http.Server {
	return &http.Server{
		Handler: rtr,
		Addr:    fmt.Sprintf(":%s", PORT),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
