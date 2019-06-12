package server

import (
	"fmt"
	"net/http"

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
	//DefaultDirHTML ... Default dir to handle html
	DefaultDirHTML = fmt.Sprintf("%s/%s/", DefaultDirWEB, "html")
)

//StablishHanlders ... Stablish all the handlers to the router r
func StablishHanlders(r *mux.Router) {
	r.HandleFunc("/inicio", Inicio)
	r.HandleFunc("/guia", Guide)
	r.HandleFunc("/herramienta", Tool)
	r.HandleFunc("/contacto", Contacto)
	r.HandleFunc("/api/v1/countries", Countries)
	r.HandleFunc("/api/v1/{country}/divisions", Divisions)
	r.HandleFunc("/api/v1/{country}/{division}/teams", Teams)
	r.Path("/api/v1/{country}/{division}/{team}/properties").Queries("stadium", "{local|away|all}").HandlerFunc(PropertiesTeam)
	r.PathPrefix("/recursos/").Handler(http.StripPrefix("/recursos/", http.FileServer(http.Dir(DefaultDirRecursos))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(DefaultDirJS))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(DefaultDirCSS))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(DefaultDirWEB))))
}
