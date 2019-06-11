package logic

import (
	"fmt"
	"valuebetsmining/data/entities"
	"valuebetsmining/data/ios"
	"valuebetsmining/data/network"
)

//ProcessData ... Processing all files from the path endpoints and years
func ProcessData(runConnection, runOutput bool) error {
	if runOutput {
		if err := ios.RemoveContents(entities.DirCSVDefault); err != nil {
			return err
		}
		if err := ios.Restart(entities.DirCSVDefault); err != nil {
			return err
		}
	}
	config, err := entities.ReadFile(entities.ConnectionConfigFile)
	if err != nil {
		return err
	}
	connection := &network.Connection{Config: config}
	for _, country := range connection.Endpoint {
		if runConnection {
			if err := ios.Restart(fmt.Sprintf("%s%s", entities.DirMaster, country.Name)); err != nil {
				return err
			}
		}
		for _, key := range country.Keys {
			if runConnection {
				err := connection.WriteByCountryDivYears(connection.Year, country.Name, key)
				if err != nil {
					return err
				}
			}
			if runOutput {
				div, err := entities.NewDivision(country.Name)
				if err != nil {
					return err
				}
				if entities.ExtFile == "csv" {
					err = div.ParseFilesToCSV(connection.Year, country.Name, key)
					if err != nil {
						return err
					}
				} else if entities.ExtFile == "json" {
					panic("Not supported yet")
				}
			}
		}
	}
	return nil
}
