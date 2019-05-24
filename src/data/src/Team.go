package data

//Team ... This is where we store information about each team
type Team struct {
	Name    string `json:"name"`
	Goals   Goal   `json:"goals"`
	Results Result `json:"results"`
}

//NewTeam ... Creates a new instance of Team
func NewTeam(name string, goalTucked, goalReceived int) (Team, error) {
	goal, err := NewGoal(goalTucked, goalReceived)
	if err != nil {
		return Team{}, err
	}
	results, err := NewResults(goalTucked, goalReceived)
	if err != nil {

	}
	return Team{
		Name:    name,
		Goals:   goal,
		Results: results,
	}, nil
}

//TeamLocal ... This is the object where we will store and calculate the data to insert in mongodb(local)
type TeamLocal struct {
	Team
}

//TeamAway ... This is the object where we will store and calculate the data to insert in mongodb(away)
type TeamAway struct {
	Team
}
