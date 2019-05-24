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
func NewMatch(number, goalTucked, goalReceived, from, to int, date, result, teamLocal, teamAway string) (Match, error) {
	local, err := NewTeam(teamLocal, goalTucked, goalReceived)
	if err != nil {
		return Match{}, err
	}
	away, err := NewTeam(teamAway, goalReceived, goalTucked)
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
		GoalAway:  goalReceived,
		GoalLocal: goalTucked,
		Result:    result,
		TeamLocal: TeamLocal{local},
		TeamAway:  TeamAway{away},
	}, nil
}
