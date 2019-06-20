package logic

import (
	"fmt"
	"log"
	"valuebetsmining/data/entities"
	"valuebetsmining/data/ios"
	"valuebetsmining/data/network"
)

//ProcessData ... Processing all files from the path endpoints and years
func ProcessData(runConnection, runOutput bool) error {
	log.Println("asdasasd")
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
					if err := ios.RestartFile(fmt.Sprintf("%s%s%s%s", entities.DirCSVDefault, country.Name, key, entities.ExtFileCSV)); err != nil {
						return err
					}
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
