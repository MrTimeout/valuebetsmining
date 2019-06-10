package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

//ConfigError ... Struct that handles error of struct Config
type ConfigError struct {
	ErrorString string
}

//CountryError ... Struct that handles errors of struct country
type CountryError struct {
	ErrorString string
}

var (
	//ErrNotExistCountry ... Error that handles when dont exist a country
	ErrNotExistCountry = &CountryError{ErrorString: "Dont exist this country"}
	//ErrParsingCountry ... Error that handles an incorrect parsing of a country
	ErrParsingCountry = &CountryError{ErrorString: "Error parsing Country"}
)

//NewConfig ... Return an object to perform config file
func NewConfig(countr, division string, yearFrom, yearTo int) (Config, error) {
	c, err := ReadFile(DefaultFileConfig)
	if err != nil {
		return Config{}, err
	}
	if countr == "all" {
		c.Year = Year{From: yearFrom, To: yearTo}
		return c, nil
	}
	country := Country{
		Name: countr,
		Keys: []string{},
	}
	if division == "all" {
		for _, v := range c.Endpoint {
			if v.Name == country.Name {
				for _, va := range v.Keys {
					country.Keys = append(country.Keys, va)
				}
				break
			}
		}
	} else {
		country.Keys = append(country.Keys, division)
	}
	return Config{
		Path:     c.Path,
		Endpoint: []Country{country},
		Year:     Year{From: yearFrom, To: yearTo},
	}, nil

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
func (c Config) ExistsCountry(country string) error {
	if strings.Trim(country, " ") == "" || len(strings.Trim(country, " ")) == 0 {
		return ErrParsingCountry
	}
	index := 0
	for {
		if index >= len(c.Endpoint) {
			break
		}
		if match, _ := regexp.MatchString("^"+country+"$", c.Endpoint[index].Name); match {
			return nil
		}
		index++
	}
	return ErrNotExistCountry
}

//ExistsDivision ... Determinates if exists a specific division
func (c Config) ExistsDivision(division string) error {
	if strings.Trim(division, " ") == "" || len(strings.Trim(division, " ")) == 0 {
		return ErrParsingDivision
	}
	index := 0
	for {
		if index >= len(c.Endpoint) {
			break
		}
		for _, value := range c.Endpoint[index].Keys {
			if match, _ := regexp.MatchString("^"+division+"$", value); match {
				return nil
			}
		}
		index++
	}
	return ErrNotExitDivision
}

//ReadFile ... Read a file to parse the json
func ReadFile(file string) (Config, error) {
	if strings.Trim(file, " ") == "" || len(file) == 0 {
		return Config{}, errors.New("Error parsing file name")
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

//Error ... Func that return a string representing an error of country
func (cr *ConfigError) Error() string {
	return cr.ErrorString
}

//Error ... Func that return a string representing an error of country
func (cr *CountryError) Error() string {
	return cr.ErrorString
}
