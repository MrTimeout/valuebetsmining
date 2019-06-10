package entities

import (
	"fmt"
)

//Match ... This is where we are going to handle match between local team and away team
type Match struct {
	Number    int    `json:"number"`
	Date      string `json:"date"`
	GoalLocal int    `json:"local goals"`
	GoalAway  int    `json:"away goals"`
	Result    string `json:"result"`
	Years     Year   `json:"years"`
	TeamLocal Team   `json:"local team"`
	TeamAway  Team   `json:"away team"`
}

const (
	//HeaderLine ... First line of each csv file created
	HeaderLine = "Index,Division,Date,LocalTeam,AwayTeam,LocalGoals,AwayGoals,Result,Last10WinningLocalMatchs,Last10TiedingLocalMatchs,Last10LosingLocalMatchs,Last10WinningAwayMatchs,Last10TiedingAwayMatchs,Last10LosingAwayMatchs,Last10StreackWinningLocal,Last10StreackNoLosingLocal,Last10StreackWinningAway,Last10StreackNoLosingAway,Last10GoalsTuckedAmountLocal,Last10GoalsReceivedAmountLocal,Last10GoalsTuckedAmountAway,Last10GoalsReceivedAmountAway,Last10AverageTuckedGoalsLocal,Last10AverageReceivedGoalsLocal,Last10AverageTuckedGoalsAway,Last10AverageReceivedGoalsAway"
	//DefaultNil ... Used to replace nil in programming
	DefaultNil = "-1"
)

//NewMatch ... Creates a new instance of match struct to handle each line of each csv
func NewMatch(number, goalsTucked, goalsReceived, from, to int, date, result, teamLocal, teamAway string) (Match, error) {
	local, err := NewTeam(teamLocal, goalsTucked, goalsReceived)
	if err != nil {
		return Match{}, err
	}
	away, err := NewTeam(teamAway, goalsReceived, goalsTucked)
	if err != nil {
		return Match{}, err
	}
	return Match{
		Number: number,
		Date:   date,
		Years: Year{
			From: from,
			To:   to,
		},
		GoalAway:  goalsReceived,
		GoalLocal: goalsTucked,
		Result:    result,
		TeamLocal: local,
		TeamAway:  away,
	}, nil
}

//NewMatchReusingLocal ... Creates a new instance of match struct to handle each line of each csv. It resuses the object local
func NewMatchReusingLocal(number, goalsTucked, goalsReceived, from, to int, date, result, teamAway string, teamLocal Team) (Match, error) {
	away, err := NewTeam(teamAway, goalsReceived, goalsTucked)
	if err != nil {
		return Match{}, err
	}
	err = teamLocal.Update(goalsTucked, goalsReceived)
	if err != nil {
		return Match{}, err
	}
	return Match{
		Number: number,
		Date:   date,
		Years: Year{
			From: from,
			To:   to,
		},
		GoalAway:  goalsReceived,
		GoalLocal: goalsTucked,
		Result:    result,
		TeamLocal: teamLocal,
		TeamAway:  away,
	}, nil
}

//NewMatchReusingAway ... Creates a new instance of match struct to handle each line of each csv. It reuses the away team
func NewMatchReusingAway(number, goalsTucked, goalsReceived, from, to int, date, result, teamLocal string, teamAway Team) (Match, error) {
	local, err := NewTeam(teamLocal, goalsTucked, goalsReceived)
	if err != nil {
		return Match{}, err
	}
	err = teamAway.Update(goalsReceived, goalsTucked)
	if err != nil {
		return Match{}, err
	}
	return Match{
		Number: number,
		Date:   date,
		Years: Year{
			From: from,
			To:   to,
		},
		GoalAway:  goalsReceived,
		GoalLocal: goalsTucked,
		Result:    result,
		TeamLocal: local,
		TeamAway:  teamAway,
	}, nil
}

//NewMatchReusingBoth ... Creates a new instance of match struct to handle each line of each csv. It reuses the away and local team
func NewMatchReusingBoth(number, goalsTucked, goalsReceived, from, to int, date, result string, teamLocal Team, teamAway Team) (Match, error) {
	err := teamLocal.Update(goalsTucked, goalsReceived)
	if err != nil {
		return Match{}, err
	}
	err = teamAway.Update(goalsReceived, goalsTucked)
	if err != nil {
		return Match{}, err
	}
	return Match{
		Number: number,
		Date:   date,
		Years: Year{
			From: from,
			To:   to,
		},
		GoalAway:  goalsReceived,
		GoalLocal: goalsTucked,
		Result:    result,
		TeamLocal: teamLocal,
		TeamAway:  teamAway,
	}, nil
}

//StringCSV ... Return a string of data in format csv
func (m Match) StringCSV(count int, line []string, tLocal, tAway bool) (string, error) {
	var previousLocal, previousAway Team
	var err error
	allNil := func(count int, line []string) string {
		return fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
			count,
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
			DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil, DefaultNil)
	}
	localNil := func(count int, line []string, previousAway Team) string {
		return fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%d,%d,%s,%s,%d,%d,%s,%s,%.2f,%.2f",
			count,
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
			DefaultNil,
			previousAway.Results.StringCSV(DefaultLenResult),
			DefaultNil,
			DefaultNil,
			previousAway.Results.StreackWinning,
			previousAway.Results.StreackNoLosing,
			DefaultNil,
			DefaultNil,
			previousAway.Goals.GoalsTuckedAmount,
			previousAway.Goals.GoalsReceivedAmount,
			DefaultNil,
			DefaultNil,
			previousAway.Goals.GoalsTuckedAverage,
			previousAway.Goals.GoalsReceivedAverage)
	}
	awayNil := func(count int, line []string, previousLocal Team) string {
		return fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%d,%d,%s,%s,%d,%d,%s,%s,%.2f,%2.f,%s,%s",
			count,
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
			previousLocal.Results.StringCSV(DefaultLenResult),
			DefaultNil,
			previousLocal.Results.StreackWinning,
			previousLocal.Results.StreackNoLosing,
			DefaultNil,
			DefaultNil,
			previousLocal.Goals.GoalsTuckedAmount,
			previousLocal.Goals.GoalsReceivedAmount,
			DefaultNil,
			DefaultNil,
			previousLocal.Goals.GoalsTuckedAverage,
			previousLocal.Goals.GoalsReceivedAverage,
			DefaultNil,
			DefaultNil)
	}
	if !tLocal && !tAway {
		return allNil(count, line), nil
	}
	if tLocal && !tAway {
		previousLocal, err = m.TeamLocal.PreviousNTeamOfAMatch(1)
		if err != nil {
			return "", err
		}
		return awayNil(count, line, previousLocal), nil
	}

	if tAway && !tLocal {
		previousAway, err = m.TeamAway.PreviousNTeamOfAMatch(1)
		if err != nil {
			return "", err
		}
		return localNil(count, line, previousAway), nil
	}

	previousLocal, err = m.TeamLocal.PreviousNTeamOfAMatch(1)
	if err != nil {
		return "", err
	}

	previousAway, err = m.TeamAway.PreviousNTeamOfAMatch(1)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%d,%d,%d,%d,%d,%d,%d,%d,%.2f,%.2f,%.2f,%.2f",
		count,
		line[0],
		line[1],
		line[2],
		line[3],
		line[4],
		line[5],
		line[6],
		previousLocal.Results.StringCSV(DefaultLenResult),
		previousAway.Results.StringCSV(DefaultLenResult),
		previousLocal.Results.StreackWinning,
		previousLocal.Results.StreackNoLosing,
		previousAway.Results.StreackWinning,
		previousAway.Results.StreackNoLosing,
		previousLocal.Goals.GoalsTuckedAmount,
		previousLocal.Goals.GoalsReceivedAmount,
		previousAway.Goals.GoalsTuckedAmount,
		previousAway.Goals.GoalsReceivedAmount,
		previousLocal.Goals.GoalsTuckedAverage,
		previousLocal.Goals.GoalsReceivedAverage,
		previousAway.Goals.GoalsTuckedAverage,
		previousAway.Goals.GoalsReceivedAverage), nil
}
