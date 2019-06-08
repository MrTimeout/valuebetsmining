package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	mongo "valuebetsmining/src/src/mongodb"

	"github.com/gorilla/mux"
)

var (
	//PORT ... Port of the golang program
	PORT = os.Getenv("PORT")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{country}/{division}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		driver, err := mongo.ConnectDB()
		if err != nil {
			io.WriteString(w, fmt.Sprintf("%d", http.StatusNotFound))
		}
		str, err := driver.GetAllTeamName(vars["country"] + vars["division"] + "1019")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("%#q", err))
		}
		for _, v := range str {
			io.WriteString(w, fmt.Sprintf("<h1>%s</h1><br>", v))
		}
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
