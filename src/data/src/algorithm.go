package data

/*
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
}*/
/*
//ParseData ... Parsing data to create new files and insert into the database
func ParseData([]string) error {
	teamsLocal, teamsAway := make(map[string]Team), make(map[string]Team)
	matchs := []Match{}
	count := 1
	reader.Read() //First line
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		goalsTucked, err := strconv.Atoi(line[4])
		if err != nil {
			return err
		}
		goalsReceived, err := strconv.Atoi(line[5])
		if err != nil {
			return err
		}
		if _, ok := teamsLocal[line[2]]; !ok {
			if _, ok := teamsAway[line[3]]; !ok {
				match, err := NewMatch(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[2], line[3])
				if err != nil {
					return err
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			} else {
				match, err := NewMatchReusingAway(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[2], teamsAway[line[3]])
				if err != nil {
					return err
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			}
		} else if _, ok := teamsLocal[line[2]]; ok {
			if _, ok := teamsAway[line[3]]; !ok {
				match, err := NewMatchReusingLocal(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[3], teamsLocal[line[2]])
				if err != nil {
					return err
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			} else {
				match, err := NewMatchReusingBoth(count, goalsTucked, goalsReceived, from, to, line[1], line[6], teamsLocal[line[2]], teamsAway[line[3]])
				if err != nil {
					return err
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			}
		}
		count++
	}
	matchsJSON, err := json.Marshal(matchs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("test/SP1_1819.json", matchsJSON, 0644)
	if err != nil {
		return err
	}
}
*/
