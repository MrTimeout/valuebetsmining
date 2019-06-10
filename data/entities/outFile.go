package entities

import "encoding/json"

//ConfigFile ... Config of the file in the output
type ConfigFile struct {
	Lines     int    `json:"lines"`
	Extension string `json:"extension"`
}

//ErrorConfigFile ... Struct that hanldes errors
type ErrorConfigFile struct {
	ErrorString string
}

var (
	//DefaultsExt ... Default extension of the output file
	DefaultsExt = []string{"json", "csv"}
	//DefaultFileConfigFile ... Default file to use to extract all the data
	DefaultFileConfigFile = "config/configFile.json"
	//DefaultFileConfigFileSchema ... Default file of config where it stored a default config file
	DefaultFileConfigFileSchema = "config/file.json"
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
	return ConfigFile{}, nil
}

//SkeletonConfigFile ... Return a json of the structure of the file json
func SkeletonConfigFile() (string, error) {
	n, _ := NewConfigFile(10, "csv")
	json, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

//Error
func (ecf *ErrorConfigFile) Error() string {
	return ecf.ErrorString
}
