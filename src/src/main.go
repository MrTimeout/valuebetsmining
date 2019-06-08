package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	mongo "valuebetsmining/src/src/mongodb"
)

var (
	//PORT ... Port of the golang program
	PORT = os.Getenv("PORT")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		driver, err := mongo.ConnectDB()
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Status code: %d", http.StatusNotFound))
		}
		str, err := driver.GetAllCollectionNames(mongo.DBDbase)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Status code: %d", http.StatusNotFound))
		}
		for k, v := range str {
			log.Println(k, v)
			io.WriteString(w, fmt.Sprintf("%d: %#q", k, v))
		}
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
