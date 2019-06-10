package server

import (
	"fmt"
	"net/http"
	"os"
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
