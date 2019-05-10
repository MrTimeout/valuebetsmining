package Handlers
 
import (
    "encoding/json"
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"

    "../Interfaces"
)
 
func GetAllteams(w http.ResponseWriter, r *http.Request) {
    
}
 
func GetTeamByName(w http.ResponseWriter, r *http.Request) {
    teams := Interfaces.Teams {
        Interfaces.Team { Name: "Barcelona", League: "Española" },
        Interfaces.Team { Name: "Madrid", League: "Española" },
    }
 
    if err := json.NewEncoder(w).Encode(teams); err != nil {
        panic(err)
    }
}
 
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}