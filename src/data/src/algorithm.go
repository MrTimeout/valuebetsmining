package data

//ProcessData ... Processing all files from the path endpoints and years
func ProcessData() error {
	config, err := ReadFile("config.json")
	if err != nil {
		return err
	}
	connection := &Connection{config}
	for _, country := range connection.Endpoint {
		for _, key := range country.Keys {
			err := connection.WriteByCountryDivYears(connection.Year, country.Name, key)
			if err != nil {
				return err
			}
			div, err := NewDivision(country.Name)
			if err != nil {
				return err
			}
			err = div.ParseFiles(connection.Year, country.Name, key)
			if err != nil {
				return err
			}
			err = div.ParseFilesToCSV(connection.Year, country.Name, key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
