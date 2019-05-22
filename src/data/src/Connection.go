package data

import (
	"errors"
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
func (c Connection) GetAllByCountryDiv(country, div string) ([]string, error) {
	if resul, err := c.ExistsCountry(country); err != nil {
		return []string{}, err
	} else if !resul {
		return []string{}, errors.New("Errors getting country")
	}
	if resul, err := c.ExistsDivision(div); err != nil {
		return []string{}, err
	} else if !resul {
		return []string{}, errors.New("Errors getting division")
	}
	result := []string{}
	for _, value := range c.Year.GetYears() {
		res, err := http.Get(fmt.Sprintf("%s/%s/%s.csv", c.Path, value, div))
		if err != nil {
			return []string{}, err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return []string{}, err
		}
		result = append(result, string(body))
	}
	return result, nil
}
