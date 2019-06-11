package entities

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

//ConfigFile ... Config of the file in the output
type ConfigFile struct {
	Lines     int    `json:"lines"`
	Extension string `json:"extension"`
}

//ErrorConfigFile ... Struct that hanldes errors
type ErrorConfigFile struct {
	ErrorString string
}

const (
	//DefaultFileConfigOutput ... Default file to use to extract all the data
	DefaultFileConfigOutput = "config/defaultOutput.json"
	//DefaultFileConfigOutputSchema ... Default file of config where it stored a default config file
	DefaultFileConfigOutputSchema = "config/output.json"
)

var (
	//DefaultsExt ... Default extension of the output file
	DefaultsExt = []string{"json", "csv"}
	//MaxLines number of lines to be parsed before
	MaxLines int
	//ExtFile ... Extension of the output of the file
	ExtFile string
	//OutputConfigFile ... Output file config
	OutputConfigFile string
)

var (
	//ErrInvalidLine ... Error parsing argument line
	ErrInvalidLine = &ErrorConfigFile{ErrorString: "Error parsing lines because it is greater than maxLine or min than 0"}
	//ErrInvalidExtension ... Error parsing extension of the output of the file
	ErrInvalidExtension = &ErrorConfigFile{ErrorString: "Error parsing extension file between the defaults"}
)

//NewConfigFile ... Return an instance of the config file
func NewConfigFile(l int, ext string) (ConfigFile, error) {
	if l <= 0 || l >= 20 {
		return ConfigFile{}, ErrInvalidLine
	}
	for _, v := range DefaultsExt {
		if v == ext {
			return ConfigFile{Lines: l, Extension: ext}, nil
		}
	}
	return ConfigFile{}, ErrInvalidExtension
}

//SkeletonConfigFile ... Return a json of the structure of the file json
func SkeletonConfigFile() (string, error) {
	n, _ := NewConfigFile(10, "csv")
	json, err := json.MarshalIndent(n, "", "	")
	if err != nil {
		return "", err
	}
	return string(json), nil
}

//ReadFileTakingArgs ... Read a file to parse the json
func ReadFileTakingArgs(file string) (ConfigFile, error) {
	if strings.Trim(file, " ") == "" || len(file) == 0 {
		return ConfigFile{}, errors.New("Error parsing file name")
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return ConfigFile{}, err
	}
	var config ConfigFile
	err = json.Unmarshal(data, &config)
	if err != nil {
		return ConfigFile{}, err
	}
	MaxLines, ExtFile = config.Lines, config.Extension
	return config, nil
}

//ObtainLines ...
func ObtainLines() int {
	return MaxLines
}

//Error ...
func (ecf *ErrorConfigFile) Error() string {
	return ecf.ErrorString
}
