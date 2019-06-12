package entities

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	others "valuebetsmining/data/Others"
)

//Division ... Each division of each league
type Division struct {
	Name        string          `json:"Name"`
	Matchs      []Match         `json:"Matchs"`
	TeamsLocal  map[string]Team `json:"TeamsLocal"`
	TeamsAway   map[string]Team `json:"TeamsAway"`
	LengthMatch int             `json:"Amount of matches"`
}

//DivisionError ... Struct that handles errors of struct Division
type DivisionError struct {
	ErrorString string
}

const (
	//DirJSONDefault ... Default path to JSON files
	DirJSONDefault = "leagues/JSON/"
	//DirCSVDefault ... Default path to CSV files
	DirCSVDefault = "leagues/CSV/"
	//DirMaster ... Default father path of data files
	DirMaster = "leagues/"
	//ExtFileJSON ... Default extension of json files
	ExtFileJSON = ".json"
	//ExtFileCSV ... Default extension of csv files
	ExtFileCSV = ".csv"
)

var (
	//ErrNotExitDivision ... Error that handles when dont exist a division
	ErrNotExitDivision = &DivisionError{ErrorString: "Dont exist this division"}
	//ErrParsingDivision ... Error that handles an incorrect parsing of a division
	ErrParsingDivision = &DivisionError{ErrorString: "Error parsing division"}
)

//NewDivision ... Creates a new object that represent a division
func NewDivision(name string) (Division, error) {
	if strings.Trim(name, " ") == "" || len(strings.Trim(name, " ")) == 0 {
		return Division{}, ErrParsingDivision
	}
	return Division{
		Name:        name,
		Matchs:      make([]Match, 0),
		TeamsLocal:  make(map[string]Team, 0),
		TeamsAway:   make(map[string]Team, 0),
		LengthMatch: 0,
	}, nil
}

//ParseFiles ... Parse files using a path and the struct Division
func (d *Division) ParseFiles(year Year, country, div string) error {
	for i := year.From; i < year.To; i++ {
		err := d.ParseEachFile(Year{From: i, To: i + 1}, fmt.Sprintf("%s%s/%s_%d%d", DirMaster, country, div, i, i+1))
		if err != nil {
			return err
		}
	}
	marshallJSON, err := json.Marshal(d.Matchs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s%s_%s_%d%d%s", DirJSONDefault, country, div, year.From, year.To, ExtFileJSON), marshallJSON, 0666)
	return nil
}

//ParseEachFile ... Parsing data to create new files and insert into the database. Using struct Division
func (d *Division) ParseEachFile(year Year, path string) error {
	csvFile, err := os.Open(fmt.Sprintf("%s%s", path, ExtFileCSV))
	if err != nil {
		return err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
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
		if t, err := others.AnyoneIsEmpty(line[:6]); err != nil {
			return err
		} else if t {
			continue
		}
		goalsTucked, err := strconv.Atoi(line[4])
		if err != nil {
			return err
		}
		goalsReceived, err := strconv.Atoi(line[5])
		if err != nil {
			return err
		}
		if _, ok := d.TeamsLocal[line[2]]; !ok {
			if _, ok := d.TeamsAway[line[3]]; !ok {
				match, err := NewMatch(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], line[2], line[3])
				if err != nil {
					return err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
			} else {
				match, err := NewMatchReusingAway(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], line[2], d.TeamsAway[line[3]])
				if err != nil {
					return err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
			}
		} else if _, ok := d.TeamsLocal[line[2]]; ok {
			if _, ok := d.TeamsAway[line[3]]; !ok {
				match, err := NewMatchReusingLocal(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], line[3], d.TeamsLocal[line[2]])
				if err != nil {
					return err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
			} else {
				match, err := NewMatchReusingBoth(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], d.TeamsLocal[line[2]], d.TeamsAway[line[3]])
				if err != nil {
					return err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
			}
			count++
		}
	}
	return nil
}

//ParseFilesToCSV ... Parse files using a path and the struct Division creating an unique csv file
func (d *Division) ParseFilesToCSV(year Year, country, div string) error {
	fileCSV, err := os.Create(fmt.Sprintf("%s%s%s%s", DirCSVDefault, country, div, ExtFileCSV))
	if err != nil {
		return err
	}
	defer fileCSV.Close()
	writer := csv.NewWriter(fileCSV)
	defer writer.Flush()
	writer.Write(strings.Split(HeaderLine, ","))
	for i := year.From; i < year.To; i++ {
		strTemp, err := d.ParseEachFileToCSV(Year{From: i, To: i + 1}, fmt.Sprintf("%s%s/%s_%d%d", DirMaster, country, div, i, i+1))
		if err != nil {
			return err
		}
		writer.WriteAll(strTemp)
	}
	if err != nil {
		return err
	}
	return nil
}

//ParseEachFileToCSV ... Parse data of a file to write in other csv all together and parsed
func (d *Division) ParseEachFileToCSV(year Year, path string) ([][]string, error) {
	csvFile, err := os.Open(fmt.Sprintf("%s%s", path, ExtFileCSV))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	count := 1
	reader.Read() //First line
	var final map[string]string
	str := make([][]string, 0, 0)
	if year.To == DefaultMaxYear {
		final = make(map[string]string)
	}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if t, err := others.AnyoneIsEmpty(line[:7]); err != nil {
			return nil, err
		} else if t {
			break
		}
		goalsTucked, err := strconv.Atoi(line[4])
		if err != nil {
			return nil, err
		}
		goalsReceived, err := strconv.Atoi(line[5])
		if err != nil {
			return nil, err
		}
		if _, ok := d.TeamsLocal[line[2]]; !ok {
			if _, ok := d.TeamsAway[line[3]]; !ok {
				match, err := NewMatch(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], line[2], line[3])
				if err != nil {
					return nil, err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
				if year.To == DefaultMaxYear {
					final[match.TeamLocal.Name] = "yes"
					final[match.TeamAway.Name] = "yes"
				}
				target, err := match.StringCSV(count, line, false, false)
				if err != nil {
					return nil, err
				}
				str = append(str, strings.Split(target, ","))
			} else {
				match, err := NewMatchReusingAway(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], line[2], d.TeamsAway[line[3]])
				if err != nil {
					return nil, err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
				if year.To == DefaultMaxYear {
					final[match.TeamLocal.Name] = "yes"
					final[match.TeamAway.Name] = "yes"
				}
				target, err := match.StringCSV(count, line, false, true)
				if err != nil {
					return nil, err
				}
				str = append(str, strings.Split(target, ","))
			}
		} else if _, ok := d.TeamsLocal[line[2]]; ok {
			if _, ok := d.TeamsAway[line[3]]; !ok {
				match, err := NewMatchReusingLocal(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], line[3], d.TeamsLocal[line[2]])
				if err != nil {
					return nil, err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
				if year.To == DefaultMaxYear {
					final[match.TeamLocal.Name] = "yes"
					final[match.TeamAway.Name] = "yes"
				}
				target, err := match.StringCSV(count, line, true, false)
				if err != nil {
					return nil, err
				}
				str = append(str, strings.Split(target, ","))
			} else {
				match, err := NewMatchReusingBoth(count, goalsTucked, goalsReceived, year.From, year.To, line[1], line[6], d.TeamsLocal[line[2]], d.TeamsAway[line[3]])
				if err != nil {
					return nil, err
				}
				d.Matchs = append(d.Matchs, match)
				d.TeamsLocal[line[2]] = match.TeamLocal
				d.TeamsAway[line[3]] = match.TeamAway
				if year.To == DefaultMaxYear {
					final[match.TeamLocal.Name] = "yes"
					final[match.TeamAway.Name] = "yes"
				}
				target, err := match.StringCSV(count, line, true, true)
				if err != nil {
					return nil, err
				}
				str = append(str, strings.Split(target, ","))
			}
		}

		count++
	}
	if year.To == DefaultMaxYear {
		for k := range final {
			match, err := NewMatchReusingBoth(0, 0, 0, year.From, year.To, "test", "test", d.TeamsLocal[k], d.TeamsAway[k])
			if err != nil {
				return nil, err
			}
			d.Matchs = append(d.Matchs, match)
			d.TeamsLocal[k] = match.TeamLocal
			d.TeamsAway[k] = match.TeamAway
			target, err := match.StringCSV(0, []string{"test", "test", k, k, "-1", "-1", "test"}, true, true)
			if err != nil {
				return nil, err
			}
			str = append(str, strings.Split(target, ","))
		}
	}
	return str, nil
}

//UpdateLengthMatch ... Update the length of the match inside the struct
func (d Division) UpdateLengthMatch() {
	d.LengthMatch = len(d.Matchs)
}

//Error ... Func that returns a string representing an error the struct Division
func (dr *DivisionError) Error() string {
	return dr.ErrorString
}
