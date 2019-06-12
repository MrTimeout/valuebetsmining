package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"valuebetsmining/src/mongodb/models"
	"valuebetsmining/src/server/entities"

	"github.com/gorilla/mux"
)

const (
	//DefaultExtHTML ... Default extension of html files
	DefaultExtHTML = "html"
)

//Inicio ... Handler that return the page index.html
func Inicio(w http.ResponseWriter, r *http.Request) {
	ree, path := regexp.MustCompile(`^/?(inicio|guia|contacto|herramienta)(\.htm(l)?)?$`), mux.Vars(r)
	if str := ree.FindAllStringSubmatch(path["route"], -1); len(str) == 1 && (len(str[0]) >= 2 || len(str[0]) <= 3) {
		fileName := fmt.Sprintf("%s/%s.%s", DefaultDirWEB, str[0][1], DefaultExtHTML)
		if _, err := os.Stat(fileName); err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			tmplt := template.New("Error")
			tmplt, _ = tmplt.ParseFiles(DefaultErrFile)
			errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
			if err == nil {
				tmplt.Execute(w, errHTTP)
			}
		} else {
			session, _ := entities.Store.Get(r, "X-auttentication")
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				if str[0][1] == "herramienta" {
					tmplt, _ := template.ParseFiles(DefaultErrFile)
					errHTTP, err := entities.NewErrHTTP(entities.ErrForbidden, http.StatusForbidden)
					if err == nil {
						tmplt.Execute(w, errHTTP)
					}
				} else {
					var buff bytes.Buffer
					t, _ := template.ParseFiles("web/nav.html")
					user := entities.User{Name: "", Cookie: false}
					if err := t.Execute(&buff, user); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						tmplt, _ := template.ParseFiles(DefaultErrFile)
						errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
						if err == nil {
							tmplt.Execute(w, errHTTP)
						}
					}
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "text/html; charset=utf-8")
					h := map[string]interface{}{
						"Header": template.HTML(buff.String()),
					}
					t, _ = template.ParseFiles(fileName)
					if err := t.Execute(w, h); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						tmplt := template.New("Error")
						tmplt, _ = tmplt.ParseFiles(DefaultErrFile)
						errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
						if err == nil {
							tmplt.Execute(w, errHTTP)
						}
					}
				}
			} else {
				var buff bytes.Buffer
				t, _ := template.ParseFiles("web/nav.html")
				user := entities.User{Name: session.Values["name"].(string), Cookie: session.Values["authenticated"].(bool)}
				if err := t.Execute(&buff, user); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					tmplt, _ := template.ParseFiles(DefaultErrFile)
					errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
					if err == nil {
						tmplt.Execute(w, errHTTP)
					}
				}
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				h := map[string]interface{}{
					"Header": template.HTML(buff.String()),
				}
				t, _ = template.ParseFiles(fileName)
				if err := t.Execute(w, h); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					tmplt, _ := template.ParseFiles(DefaultErrFile)
					errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
					if err == nil {
						tmplt.Execute(w, errHTTP)
					}
				}
			}
		}
	}
}

//Tool ... Handler that return the page herramienta.html
func Tool(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("%s/%s", DefaultDirWEB, "herramienta.html")
	if _, err := os.Stat(fileName); err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	}
	http.ServeFile(w, r, fileName)
}

//PropertiesTeam ... Return data that is in the database corresponding to the properties of the team
func PropertiesTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stadium := r.FormValue("stadium")
	if stadium == "local" || stadium == "away" || stadium == "all" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
	switch stadium {
	case "local":
		resl, err := models.GetPropertiesOfALocalTeam(fmt.Sprintf("%s%s", vars["country"], vars["division"]), vars["team"])
		if err == nil {
			w.WriteHeader(http.StatusFound)
			err := json.NewEncoder(w).Encode(resl)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
			}
		}
	case "away":
		resl, err := models.GetPropertiesOfAAwayTeam(fmt.Sprintf("%s%s", vars["country"], vars["division"]), vars["team"])
		if err == nil {
			w.WriteHeader(http.StatusFound)
			err := json.NewEncoder(w).Encode(resl)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
			}
		}
	case "all":
		resl, err := models.GetPropertiesOfATeam(fmt.Sprintf("%s%s", vars["country"], vars["division"]), vars["team"])
		fmt.Fprintf(w, "%+q, %v", err, resl)
		if err == nil {
			w.WriteHeader(http.StatusNotFound)
			err := json.NewEncoder(w).Encode(resl)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
			}
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	}
}

//Countries ... Return data that is in the database corresponding to the countries
func Countries(w http.ResponseWriter, r *http.Request) {
	resl, err := models.Countries()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		if len(resl) != 0 {
			w.WriteHeader(http.StatusFound)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		err := json.NewEncoder(w).Encode(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
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
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if len(resl) != 0 {
			w.WriteHeader(http.StatusFound)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		err := json.NewEncoder(w).Encode(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
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
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if len(resl) != 0 {
			w.WriteHeader(http.StatusFound)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		err := json.NewEncoder(w).Encode(resl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, fmt.Sprintf("%s/%s", DefaultDirWEB, "404.html"))
		}
	}
}

//Error404 ... Error not found
func Error404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var buff, buf bytes.Buffer
		t, _ := template.ParseFiles("web/nav.html")
		user := entities.User{Name: "", Cookie: false}
		if err := t.Execute(&buff, user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmplt, _ := template.ParseFiles(DefaultErrFile)
			errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
			if err == nil {
				tmplt.Execute(w, errHTTP)
			}
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		h := map[string]interface{}{
			"Header": template.HTML(buff.String()),
		}
		log.Println(buff.String())
		t, _ = template.ParseFiles("web/error.html")
		errHTTP, _ := entities.NewErrHTTP(entities.ErrNotFound, http.StatusNotFound)
		t.Execute(&buf, errHTTP)
		log.Println("asdasdasd")
		if err := t.Execute(w, h); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			tmplt, _ := template.ParseFiles(DefaultErrFile)
			errHTTP, err := entities.NewErrHTTP(entities.ErrInternalServerError, http.StatusInternalServerError)
			if err == nil {
				tmplt.Execute(w, errHTTP)
			}
		}
	})
}

//Error403 ... Error forbidden
func Error403(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	tmplt, _ := template.ParseFiles(DefaultErrFile)
	errHTTP, _ := entities.NewErrHTTP(entities.ErrForbidden, http.StatusForbidden)
	tmplt.Execute(w, errHTTP)
}
