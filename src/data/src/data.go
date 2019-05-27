package data

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//Config ... Struct created to parse JSON config.json and get information
type Config struct {
	Path     string    `json:"Path"`
	Year     Year      `json:"Year"`
	Endpoint []Country `json:"Endpoints"`
}

//Year ... Struct created to parse JSON config.json and get information
type Year struct {
	From int `json:"from"`
	To   int `json:"to"`
}

//Country ... Strcut created to pase JSON config.json and get information
type Country struct {
	Name string   `json:"name"`
	Keys []string `json:"keys"`
}

//GetYears ... Returns years [from, to] in an array of strings in the format: from(from+1),...
func (y Year) GetYears() []string {
	res := []string{}
	index := 1
	for {
		if index == y.To-y.From {
			break
		}
		res = append(res, fmt.Sprintf("%d%d", y.From+index, y.From+index+1))
		index++
	}
	return res
}

//ExistsCountry ... Determinates if exists a specific country
func (c Config) ExistsCountry(country string) (bool, error) {
	if strings.Trim(country, " ") == "" || len(strings.Trim(country, " ")) == 0 {
		return false, errors.New("Error parsing country")
	}
	index := 0
	for {
		if index >= len(c.Endpoint) {
			break
		}
		if match, _ := regexp.MatchString("^"+country+"$", c.Endpoint[index].Name); match {
			return true, nil
		}
		index++
	}
	return false, nil
}

//ExistsDivision ... Determinates if exists a specific division
func (c Config) ExistsDivision(division string) (bool, error) {
	if strings.Trim(division, " ") == "" || len(strings.Trim(division, " ")) == 0 {
		return false, errors.New("Error parsing division")
	}
	index := 0
	for {
		if index >= len(c.Endpoint) {
			break
		}
		for _, value := range c.Endpoint[index].Keys {
			if match, _ := regexp.MatchString("^"+division+"$", value); match {
				return true, nil
			}
		}
		index++
	}
	return false, nil
}
