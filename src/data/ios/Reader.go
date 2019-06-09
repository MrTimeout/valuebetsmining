package ios

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"valuebetsmining/src/data/entities"
)

//ReadFile ... Read a file to parse the json
func ReadFile(file string) (entities.Config, error) {
	if strings.Trim(file, " ") == "" || len(file) == 0 {
		return entities.Config{}, errors.New("Error parsing file name")
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return entities.Config{}, err
	}
	var config entities.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return entities.Config{}, err
	}
	return config, nil
}
