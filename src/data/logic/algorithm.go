package logic

import (
	"valuebetsmining/src/data/entities"
	"valuebetsmining/src/data/ios"
	"valuebetsmining/src/data/network"
)

//ProcessData ... Processing all files from the path endpoints and years
func ProcessData() error {
	config, err := ios.ReadFile("config.json")
	if err != nil {
		return err
	}
	connection := &network.Connection{Config: config}
	for _, country := range connection.Endpoint {
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
