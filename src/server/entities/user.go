package entities

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

//User ... Handles info about user
type User struct {
	Name   string
	Cookie bool
}

var (
	//Key ... Super secret key
	Key = []byte("dont-tell-anyone")
	//Store ... This is where the key is stored
	Store = sessions.NewCookieStore(Key)
)

//Login ... This is where the user is logged
func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "X-auttentication")
	log.Println(r.FormValue("user"), r.FormValue("password"), r.URL.RawPath)
	if r.FormValue("user") == "root" && r.FormValue("password") == "abc123." {
		session.Values["authenticated"] = true
		session.Values["name"] = r.FormValue("username")
		session.Save(r, w)
		http.Redirect(w, r, r.URL.EscapedPath(), 302)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		tmplt := template.New("Error")
		tmplt, _ = tmplt.ParseFiles("web/errors.html")
		errHTTP, err := NewErrHTTP(ErrInternalServerError, http.StatusInternalServerError)
		if err == nil {
			tmplt.Execute(w, errHTTP)
		}
	}

}

//LogOut ... This is where the user log out
func LogOut(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "X-auttentication")
	log.Println(session.Values["authenticated"])
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/inicio", 302)
}
