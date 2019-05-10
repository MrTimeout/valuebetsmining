package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"

    "../Interfaces"
)
 
func getAllteams(w http.ResponseWriter, r *http.Request) {
    
}
 
func getTeamByName(w http.ResponseWriter, r *http.Request) {
    teams := Teams{
        Team { "Barcelona", "Española" },
        Team { "Madrid", "Española" },
    }
 
    if err := json.NewEncoder(w).Encode(Teams); err != nil {
        panic(err)
    }
}
 
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}