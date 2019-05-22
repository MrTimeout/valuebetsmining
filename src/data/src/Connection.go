package data

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Connection ... Struct used to connect to an url and download it
type Connection struct {
	Config
}

//Get ... Download the content of an endpoint
func (c Connection) Get() (string, error) {
	fmt.Println(fmt.Sprintf("%s/%s/%s.csv", c.Path, c.Year.GetYears()[0], c.Endpoint[0].Keys[0]))
	res, err := http.Get(fmt.Sprintf("%s/%s/%s.csv", c.Path, c.Year.GetYears()[0], c.Endpoint[0].Keys[0]))
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

//GetAllByCountryDiv ... Download all the content of all years of a country/div
func (c Connection) GetAllByCountryDiv(country, div string) (string, error) {

	return "", nil
}
