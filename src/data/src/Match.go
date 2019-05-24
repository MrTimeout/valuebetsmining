package data

//Match ... This is where we are going to handle match between local team and away team
type Match struct {
	Number    int       `json:"number"`
	Date      string    `json:"date"`
	Years     Year      `json:"years"`
	TeamLocal TeamLocal `json:"local team"`
	TeamAway  TeamAway  `json:"away team"`
	GoalAway  int       `json:"local goals"`
	GoalLocal int       `json:"away goals"`
	Result    int       `json:"result"`
}

//NewMatch ... Creates a new instance of match struct to handle each line of each csv
func NewMatch(number, goalTucked, goalReceived int, date, teamLocal, teamAway string) (Match, error) {

	return Match{}, nil
}
