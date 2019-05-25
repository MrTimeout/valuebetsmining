package data

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
