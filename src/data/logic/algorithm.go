package logic

import (
	"fmt"
	"os"
	"valuebetsmining/src/data/entities"
	"valuebetsmining/src/data/ios"
	"valuebetsmining/src/data/network"
)

//ProcessData ... Processing all files from the path endpoints and years
func ProcessData() error {
	err := os.Mkdir(entities.DirMaster, 0777)
	if err != nil {
		return err
	}
	err = os.Mkdir(entities.DirCSVDefault, 0777)
	if err != nil {
		return err
	}
	err = os.Mkdir(entities.DirJSONDefault, 0777)
	if err != nil {
		return err
	}
	config, err := ios.ReadFile(entities.ConfigJSONFile)
	if err != nil {
		return err
	}
	connection := &network.Connection{Config: config}
	for _, country := range connection.Endpoint {
		err := os.Mkdir(fmt.Sprintf("%s%s", entities.DirMaster, country.Name), 0777)
		if err != nil {
			return err
		}
		for _, key := range country.Keys {
			err := connection.WriteByCountryDivYears(connection.Year, country.Name, key)
			if err != nil {
				return err
			}
			div, err := entities.NewDivision(country.Name)
			if err != nil {
				return err
			}
			/*err = div.ParseFiles(connection.Year, country.Name, key)
			if err != nil {
				return err
			}*/
			err = div.ParseFilesToCSV(connection.Year, country.Name, key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
