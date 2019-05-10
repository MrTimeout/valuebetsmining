package main
 
import ( 
    "net/http"

    "./Handlers"
)
 
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
 
type Routes []Route
 
var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Handlers.GetAllteams,
    },
    Route{
        "TodoIndex",
        "GET",
        "/todos",
        Handlers.GetTeamByName,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        Handlers.TodoShow,
    },
}