package logic

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"valuebetsmining/src/data/entities"
)

//TestWriteJSONToAFile ... Write random info of the stucts team* to a file to understand how it works
func TestWriteJSONToAFile(t *testing.T) {
	csvFile, err := os.Open("test/SP1_1819.csv")
	if err != nil {
		t.Errorf("Error: %#v", err)
	}
	reader := csv.NewReader(csvFile)
	teamsLocal, teamsAway := make(map[string]entities.Team), make(map[string]entities.Team)
	matchs := []entities.Match{}
	from, to := 18, 19
	count := 1
	reader.Read() //First line
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}
		goalsTucked, err := strconv.Atoi(line[4])
		if err != nil {
			t.Error(err)
		}
		goalsReceived, err := strconv.Atoi(line[5])
		if err != nil {
			t.Error(err)
		}
		if _, ok := teamsLocal[line[2]]; !ok {
			if _, ok := teamsAway[line[3]]; !ok {
				match, err := entities.NewMatch(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[2], line[3])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			} else {
				match, err := entities.NewMatchReusingAway(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[2], teamsAway[line[3]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			}
		} else if _, ok := teamsLocal[line[2]]; ok {
			if _, ok := teamsAway[line[3]]; !ok {
				match, err := entities.NewMatchReusingLocal(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[3], teamsLocal[line[2]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[line[2]] = match.TeamLocal
				teamsAway[line[3]] = match.TeamAway
			} else {
				match, err := entities.NewMatchReusingBoth(count, goalsTucked, goalsReceived, from, to, line[1], line[6], teamsLocal[line[2]], teamsAway[line[3]])
				if err != nil {
					t.Error(err)
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
		t.Error(err)
	}
	err = ioutil.WriteFile("test/SP1_1819.json", matchsJSON, 0644)
	if err != nil {
		t.Error(err)
	}
}

//TestProcessData ... Process data testing the func processData
func TestProcessData(t *testing.T) {
	err := ProcessData()
	if err != nil {
		t.Error(err)
	}
}
