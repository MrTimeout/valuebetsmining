package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
		fmt.Println(i)
		err := d.ParseEachFile(Year{From: i, To: i + 1}, fmt.Sprintf("./leagues/%s/%s_%d%d", country, div, i, i+1))
		if err != nil {
			return err
		}
	}
	marshallJSON, err := json.Marshal(d.Matchs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./leagues/JSON/%s_%s_%d%d.json", country, div, year.From, year.To), marshallJSON, 0666)
	return nil
}

//ParseFilesToCSV ... Parse files using a path and the struct Division creating an unique csv file
func (d *Division) ParseFilesToCSV(year Year, country, div string) error {
	for i := year.From; i < year.To; i++ {
		err := d.ParseEachFile(Year{From: i, To: i + 1}, fmt.Sprintf("./leagues/%s/%s_%d%d", country, div, i, i+1))
		if err != nil {
			return err
		}
	}
	//err = ioutil.WriteFile(fmt.Sprintf("./leagues/CSV/%s_%s_%d%d.csv", country, div, year.From, year.To), csvData, 0666)
	return nil
}

//ParseEachFile ... Parsing data to create new files and insert into the database. Using struct Division
func (d *Division) ParseEachFile(year Year, path string) error {
	fmt.Println(path)
	csvFile, err := os.Open(fmt.Sprintf("%s.csv", path))
	if err != nil {
		return err
	}
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
		if strings.Trim(line[0], " ") == "" {
			break
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

//UpdateLengthMatch ... Update the length of the match inside the struct
func (d Division) UpdateLengthMatch() {
	d.LengthMatch = len(d.Matchs)
}

//Error ... Func that returns a string representing an error the struct Division
func (dr *DivisionError) Error() string {
	return dr.ErrorString
}
