package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

const (
	//ConfigJSONFile ... Path to access config.json file
	ConfigJSONFile = "config/config.json"
	//MaxYear ... MaxYear of the file config
	MaxYear = 19
	//MinYear ... MInyear of the file config
	MinYear = 10
)

var (
	//ErrNotExistCountry ... Error that handles when dont exist a country
	ErrNotExistCountry = &CountryError{ErrorString: "Dont exist this country"}
	//ErrParsingCountry ... Error that handles an incorrect parsing of a country
	ErrParsingCountry = &CountryError{ErrorString: "Error parsing Country"}
)

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
		return errors.New("asdasd")
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
	return errors.New("asdasdasd")
}

//Error ... Func that return a string representing an error of country
func (cr *ConfigError) Error() string {
	return cr.ErrorString
}

//Error ... Func that return a string representing an error of country
func (cr *CountryError) Error() string {
	return cr.ErrorString
}

func main() {
	trying()
	//praticseArgs()
}

func trying() {
	fileType := flag.String("type-file", "csv", "You have to choose the file extension to parse data. Possible values: 'json' or 'csv'")
	yearFrom := flag.Int("year-from", 10, "Year from to search about teams and divisions. By default is 10 an is also the minimun, representing 2010(is inclusive)")
	yearTo := flag.Int("year-to", 19, "Year to to seach about teams and divisions. By default is 19 an is also the maximun, representing 2019(inclusive)")
	country := flag.String("country", "all", "Country to search in the machine. By default is Spain and it neccesary to use nomenclature of the file params.md to use it well. If you dont write any country, it will cacth all countries")
	division := flag.String("division", "all", "Divisions to search about a country. Possible values are all|other division. See params.md to more info")
	skeleton := flag.Bool("skeleton-json", false, "It prints the skeleton of the json to show the structure of the file to pass like a param.")
	configFile := flag.String("config-file", "", "If you use this option, use it alone and pass a correct path to the param. By default is the config/config.json file. Review it to know and understand how to work with it")

	flag.Parse()

	shell, err := NewShell()
	if err != nil {
		log.Panic(err)
	}
	if *skeleton {
		log.Println(shell.skeleton)
	} else if *configFile != "" {
		log.Println(*configFile)
	} else {
		log.Println(shell.IsFileType(*fileType))
		log.Println(shell.IsYearFrom(*yearFrom))
		log.Println(shell.IsYearTo(*yearTo))
		log.Println(shell.IsCountry(*country))
		log.Println(shell.IsDivision(*division))
	}
}

//Shell ... Handles all params of the program
type Shell struct {
	fileType   string
	yearFrom   int
	yearTo     int
	country    string
	division   string
	skeleton   string
	configFile string
	config     Config
}

//ErrorShell ... Handles errors of shell struct
type ErrorShell struct {
	ErrorString string
}

const (
	//DefaultFileConfig ... Default file of config json
	DefaultFileConfig = "config/config.json"
	//DefaultExtFile ... Default extension of the file result of the algorithm
	DefaultExtFile = "csv"
	//DefaultValueCountry ... Default value of the country: all
	DefaultValueCountry = "all"
	//DefaultValueDivision ... Default value of the division: all
	DefaultValueDivision = "all"
)

var (
	//ErrInvalidFileType ... Invalid type file because it is different from the defaults
	ErrInvalidFileType = &ErrorShell{ErrorString: "Error on type file because is distinct of json and csv"}
	//ErrInvalidYear ... Error because it an incorrect year
	ErrInvalidYear = &ErrorShell{ErrorString: "Error filtering year because it is incorrect"}
	//ErrInvalidCountry ... Error parsing country
	ErrInvalidCountry = &ErrorShell{ErrorString: "Error filtering country because it doesnt exists or it doesnt appear to exists ;)"}
	//ErrInvalidDivision ... Error parsing Division
	ErrInvalidDivision = &ErrorShell{ErrorString: "Error filtering division because it doesnt exists or it doesnt appear to exists ;)"}
)

//NewShell ... Creates a new shell with default values
func NewShell() (Shell, error) {
	sque, err := Skeleton()
	if err != nil {
		return Shell{}, err
	}
	c, err := ReadFile(DefaultFileConfig)
	if err != nil {
		return Shell{}, err
	}
	return Shell{
		fileType:   DefaultExtFile,
		yearFrom:   c.Year.From,
		yearTo:     c.Year.To,
		country:    "all",
		division:   "all",
		skeleton:   sque,
		configFile: DefaultFileConfig,
		config:     c,
	}, nil
}

//IsFileType ... Sets file type if it is correct
func (s *Shell) IsFileType(n string) error {
	if n == "json" || n == "csv" {
		s.fileType = n
		return nil
	}
	return ErrInvalidFileType
}

//IsYearFrom ... Sets year from if it appears to look good
func (s *Shell) IsYearFrom(n int) error {
	if n < s.config.Year.From || n > s.config.Year.To {
		return ErrInvalidYear
	}
	if n >= s.yearTo {
		return ErrInvalidYear
	}
	s.yearFrom = n
	return nil
}

//IsYearTo ... Sets year to if it appears to look good
func (s *Shell) IsYearTo(n int) error {
	if n < s.config.Year.From || n > s.config.Year.To {
		return ErrInvalidYear
	}
	if n <= s.yearFrom {
		return ErrInvalidYear
	}
	s.yearTo = n
	return nil
}

//IsCountry ... Sets country if it is exists
func (s *Shell) IsCountry(n string) error {
	if n == DefaultValueCountry {
		s.country = n
		return nil
	}
	for _, value := range s.config.Endpoint {
		if n == value.Name {
			s.country = n
			return nil
		}
	}
	return ErrInvalidCountry
}

//IsDivision ... Sets division if it is exists
func (s *Shell) IsDivision(n string) error {
	if n == DefaultValueDivision {
		s.division = n
		return nil
	}
	for _, value := range s.config.Endpoint {
		if value.Name == s.country {
			for _, v := range value.Keys {
				if v == n {
					s.division = n
					return nil
				}
			}
			return ErrInvalidDivision
		}
	}
	return ErrInvalidCountry
}

//Skeleton ... Return an squeleton of the config json file
func Skeleton() (string, error) {
	c := &Config{
		Endpoint: []Country{Country{Name: "Country_name_1", Keys: []string{"division_1", "division_2"}}, Country{Name: "Country_name_2", Keys: []string{"division_1", "division_2"}}},
		Path:     "Path to attack, Use the default path because is something ultra necessary",
		Year:     Year{From: 10, To: 19},
	}
	json, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

//Error ... Return errors handles
func (es *ErrorShell) Error() string {
	return es.ErrorString
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
