package entities

import (
	"encoding/json"
)

//Shell ... Handles all params of the program
type Shell struct {
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
	//DefaultDirConfig ... Default dir where all config files are stored
	DefaultDirConfig = "config"
	//DefaultFileConfig ... Default file of config json
	DefaultFileConfig = "config/config.json"
	//DefaultExtFile ... Default extension of the file result of the algorithm
	DefaultExtFile = "csv"
	//DefaultValueCountry ... Default value of the country: all
	DefaultValueCountry = "all"
	//DefaultValueDivision ... Default value of the division: all
	DefaultValueDivision = "all"
	//DefaultMaxYear ... Default value to max year
	DefaultMaxYear = 19
	//DefaultMinYear ... Default value to min year
	DefaultMinYear = 10
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
	//ErrInvalidConfigFile ... Error parsing Config because it dont have a correct struct or its name config/config.json
	ErrInvalidConfigFile = &ErrorShell{ErrorString: "Error parsing config file"}
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
		yearFrom:   c.Year.From,
		yearTo:     c.Year.To,
		country:    "all",
		division:   "all",
		skeleton:   sque,
		configFile: DefaultFileConfig,
		config:     c,
	}, nil
}

/*
//IsFileType ... Sets file type if it is correct
func (s *Shell) IsFileType(n string) error {
	if n == "json" || n == "csv" {
		s.fileType = n
		return nil
	}
	return ErrInvalidFileType
}*/

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
	json, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return "", err
	}
	return string(json), nil
}

//Error ... Return errors handles
func (es *ErrorShell) Error() string {
	return es.ErrorString
}
