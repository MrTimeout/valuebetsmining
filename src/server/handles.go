package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"valuebetsmining/src/mongodb/models"

	"github.com/gorilla/mux"
)

const (
	//DefaultBegin ... Default file of begin
	DefaultBegin = "index.html"
	//DefaultContact ... Default file of contact
	DefaultContact = "contacto.html"
	//DefaultGuide ... Default file of guide
	DefaultGuide = "guia.html"
	//DefaultTool ... Default file of tool
	DefaultTool = "herramienta.html"
)

//Inicio ... Handler that return the page index.html
func Inicio(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("%s/%s", DefaultDirWEB, DefaultBegin)
	if _, err := os.Stat(fileName); err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	}
	http.ServeFile(w, r, fileName)
}

//Contacto ... Handler that return the page contacto.html
func Contacto(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("%s/%s", DefaultDirWEB, DefaultContact)
	if _, err := os.Stat(fileName); err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	}
	http.ServeFile(w, r, fileName)
}

//Guide ... Handler that return the page guia.html
func Guide(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("%s/%s", DefaultDirWEB, DefaultGuide)
	if _, err := os.Stat(fileName); err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	}
	http.ServeFile(w, r, fileName)
}

//Tool ... Handler that return the page herramienta.html
func Tool(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("%s/%s", DefaultDirWEB, DefaultTool)
	if _, err := os.Stat(fileName); err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	}
	http.ServeFile(w, r, fileName)
}

//PropertiesTeam ... Return data that is in the database corresponding to the properties of the team
func PropertiesTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resl, err := models.GetPropertiesOfATeam(fmt.Sprintf("%s%s", vars["country"], vars["division"]), vars["team"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		marshall, err := json.Marshal(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
		} else {
			fmt.Fprintf(w, string(marshall))
		}
	}
}

//Countries ... Return data that is in the database corresponding to the countries
func Countries(w http.ResponseWriter, r *http.Request) {
	resl, err := models.Countries()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	} else {
		js, err := json.Marshal(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(js))
		}
	}
}

//Divisions ... Return data that is in the database corresponding to the divisions
func Divisions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resl, err := models.Divisions(vars["country"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	} else {
		js, err := json.Marshal(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(js))
		}
	}
}

//Teams ... Return data that is in the database corresponding to the teams
func Teams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resl, err := models.GetAllTeamName(fmt.Sprintf("%s%s", vars["country"], vars["division"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	} else {
		js, err := json.Marshal(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(js))
		}
	}
}

//Error404 ... Error not found
func Error404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	})
}

//Error403 ... Error forbidden
func Error403(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "403.html"))
}
