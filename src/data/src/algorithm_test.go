package data

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

//TestWriteJSONToAFile ... Write random info of the stucts team* to a file to understand how it works
func TestWriteJSONToAFile(t *testing.T) {
	csvFile, err := os.Open("test/SP1_1819.csv")
	if err != nil {
		t.Errorf("Error: %#v", err)
	}
	reader := csv.NewReader(csvFile)
	teamsLocal, teamsAway := make(map[string]TeamLocal), make(map[string]TeamAway)
	matchs := []Match{}
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
				match, err := NewMatch(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[2], line[3])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
			} else {
				match, err := NewMatchReusingAway(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[2], teamsAway[line[3]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
			}
		} else if _, ok := teamsLocal[line[2]]; ok {
			if _, ok := teamsAway[line[3]]; !ok {
				match, err := NewMatchReusingLocal(count, goalsTucked, goalsReceived, from, to, line[1], line[6], line[3], teamsLocal[line[2]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
			} else {
				match, err := NewMatchReusingBoth(count, goalsTucked, goalsReceived, from, to, line[1], line[6], teamsLocal[line[2]], teamsAway[line[3]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
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
