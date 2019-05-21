package data

import (
	"io/ioutil"
	"net/http"
)

//Connection ... Struct used to connect to an url and download it
type Connection struct {
	Config
}

//Get ... Download the content of an endpoint
func (c Connection) Get() (string, error) {
	res, err := http.Get("%s/")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
