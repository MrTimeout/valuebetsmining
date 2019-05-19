package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

//ReadFile ...
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
