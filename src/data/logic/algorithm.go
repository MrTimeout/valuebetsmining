package logic

import (
	"fmt"
	"valuebetsmining/src/data/entities"
	"valuebetsmining/src/data/ios"
	"valuebetsmining/src/data/network"
)

//ProcessData ... Processing all files from the path endpoints and years
func ProcessData() error {
	if err := ios.Restart(entities.DirMaster); err != nil {
		return err
	}
	if err := ios.Restart(entities.DirCSVDefault); err != nil {
		return err
	}
	if err := ios.Restart(entities.DirJSONDefault); err != nil {
		return err
	}
	config, err := ios.ReadFile(entities.ConfigJSONFile)
	if err != nil {
		return err
	}
	connection := &network.Connection{Config: config}
	for _, country := range connection.Endpoint {
		if err := ios.Restart(fmt.Sprintf("%s%s", entities.DirMaster, country.Name)); err != nil {
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
