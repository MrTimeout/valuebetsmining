package data

//Match ... This is where we are going to handle match between local team and away team
type Match struct {
	Number    int       `json:"number"`
	Date      string    `json:"date"`
	Years     Year      `json:"years"`
	GoalAway  int       `json:"local goals"`
	GoalLocal int       `json:"away goals"`
	Result    string    `json:"result"`
	TeamLocal TeamLocal `json:"local team"`
	TeamAway  TeamAway  `json:"away team"`
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
		TeamLocal: TeamLocal{local},
		TeamAway:  TeamAway{away},
	}, nil
}

//NewMatchReusingLocal ... Creates a new instance of match struct to handle each line of each csv. It resuses the object local
func NewMatchReusingLocal(number, goalsTucked, goalsReceived, from, to int, date, result, teamAway string, teamLocal TeamLocal) (Match, error) {
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
		TeamAway:  TeamAway{away},
	}, nil
}

//NewMatchReusingAway ... Creates a new instance of match struct to handle each line of each csv. It reuses the away team
func NewMatchReusingAway(number, goalsTucked, goalsReceived, from, to int, date, result, teamLocal string, teamAway TeamAway) (Match, error) {
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
		TeamLocal: TeamLocal{local},
		TeamAway:  teamAway,
	}, nil
}

//NewMatchReusingBoth ... Creates a new instance of match struct to handle each line of each csv. It reuses the away and local team
func NewMatchReusingBoth(number, goalsTucked, goalsReceived, from, to int, date, result string, teamLocal TeamLocal, teamAway TeamAway) (Match, error) {
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
