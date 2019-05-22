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
			data, err := connection.GetAllByCountryDiv(country.Name, key)
			if err != nil {
				return err
			}
			err = ParseData(data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//ParseData ... Parsing data to create new files and insert into the database
func ParseData([]string) error {
	return nil
}
