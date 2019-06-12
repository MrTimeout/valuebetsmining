package server

import (
	"fmt"
	"net/http"
	"valuebetsmining/src/server/entities"

	"github.com/gorilla/mux"
)

var (
	//DefaultDirWEB ... Default dir to follow web
	DefaultDirWEB = "web"
	//DefaultDirRecursos ... Return the path to the recursos folder
	DefaultDirRecursos = fmt.Sprintf("%s/%s/", DefaultDirWEB, "recursos")
	//DefaultDirJS ... Default dir to access js files
	DefaultDirJS = fmt.Sprintf("%s/%s/", DefaultDirWEB, "js")
	//DefaultDirCSS ... Default dir to access css files
	DefaultDirCSS = fmt.Sprintf("%s/%s/", DefaultDirWEB, "css")
	//DefaultErrFile ... Default file to set when an error ocurrs
	DefaultErrFile = fmt.Sprintf("%s/%s", DefaultDirWEB, "error.html")
)

//StablishHanlders ... Stablish all the handlers to the router r
func StablishHanlders(r *mux.Router) {
	r.HandleFunc("/{route:[a-z]+\\.?[a-z]+}", Inicio).Methods("GET")
	r.HandleFunc("/login", entities.Login).Methods("POST")
	r.HandleFunc("/logout", entities.LogOut).Methods("POST")
	r.HandleFunc("/api/v1/countries", Countries).Methods("GET")
	r.HandleFunc("/api/v1/{country}/divisions", Divisions).Methods("GET")
	r.HandleFunc("/api/v1/{country}/{division}/teams", Teams).Methods("GET")
	r.Path("/api/v1/{country}/{division}/{team}/properties").Queries("stadium", "{stadium:local|away|all}").HandlerFunc(PropertiesTeam).Methods("GET")
	r.PathPrefix("/recursos/").Handler(http.StripPrefix("/recursos/", http.FileServer(http.Dir(DefaultDirRecursos)))).Methods("GET")
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(DefaultDirJS)))).Methods("GET")
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(DefaultDirCSS)))).Methods("GET")
	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(DefaultDirWEB))))
}
