package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestNewMatchsLocal(t *testing.T) {
	matchs := make([]Match, 0, 20)
	teamsLocal, teamsAway := make(map[string]Team), make(map[string]Team)
	randomTeams, err := RandomWords(2, 10)
	if err != nil {
		t.Error(err)
	}
	goalsTucked, err := RandomArr(1, 10, 20)
	if err != nil {
		t.Error(err)
	}
	goalsReceibed, err := RandomArr(1, 10, 20)
	if err != nil {
		t.Error(err)
	}
	n, length := 0, 20
	for {
		if n == length {
			break
		}
		if _, ok := teamsLocal[randomTeams[0]]; !ok {
			if _, ok := teamsAway[randomTeams[1]]; !ok {
				match, err := NewMatch(n, goalsTucked[n], goalsReceibed[n], 18, 19, "date", "A", randomTeams[0], randomTeams[1])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[match.TeamLocal.Name] = match.TeamLocal
				teamsAway[match.TeamAway.Name] = match.TeamAway
			} else {
				match, err := NewMatchReusingAway(n, goalsTucked[n], goalsReceibed[n], 18, 19, "date", "A", randomTeams[1], teamsAway[randomTeams[1]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[match.TeamLocal.Name] = match.TeamLocal
				teamsAway[match.TeamAway.Name] = match.TeamAway
			}
		} else if _, ok := teamsLocal[randomTeams[0]]; ok {
			if _, ok := teamsAway[randomTeams[1]]; !ok {
				match, err := NewMatchReusingLocal(n, goalsTucked[n], goalsReceibed[n], 18, 19, "date", "A", randomTeams[1], teamsLocal[randomTeams[0]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[match.TeamLocal.Name] = match.TeamLocal
				teamsAway[match.TeamAway.Name] = match.TeamAway
			} else {
				match, err := NewMatchReusingBoth(n, goalsTucked[n], goalsReceibed[n], 18, 19, "date", "A", teamsLocal[randomTeams[0]], teamsAway[randomTeams[1]])
				if err != nil {
					t.Error(err)
				}
				matchs = append(matchs, match)
				teamsLocal[match.TeamLocal.Name] = match.TeamLocal
				teamsAway[match.TeamAway.Name] = match.TeamAway
			}
		}
		n++
	}
	if len(matchs) == length {
		matchsJSON, err := json.Marshal(matchs)
		if err != nil {
			t.Error(err)
		}
		err = ioutil.WriteFile("test/match_test.json", matchsJSON, 0644)
		if err != nil {
			t.Error(err)
		}
	} else {
		t.Errorf("Want: %d\n Got: %d\n", length, len(matchs))
	}
}

//TestStringCSV ... Testing String CSV
func TestStringCSV(t *testing.T) {
	match, err := NewMatch(1, 1, 2, 10, 11, "date", "A", "local", "away")
	if err != nil {
		t.Error(err)
	}
	t.Error(match.StringCSV(1, []string{"1", "2", "3", "4", "5", "6", "7"}, false, false))
}
