package data

//Team ... This is where we store information about each team
type Team struct {
	Name    string `json:"name"`
	Goals   Goal   `json:"goals"`
	Results Result `json:"results"`
}

//NewTeam ... Creates a new instance of Team
func NewTeam(name string, goalsTucked, goalsReceived int) (Team, error) {
	goal, err := NewGoal(goalsTucked, goalsReceived)
	if err != nil {
		return Team{}, err
	}
	results, err := NewResults(goalsTucked, goalsReceived)
	if err != nil {
		return Team{}, err
	}
	return Team{
		Name:    name,
		Goals:   goal,
		Results: results,
	}, nil
}

//Update ... Update values of the properties of each team
func (t Team) Update(goalsTucked, goalsReceived int) error {
	err := t.Goals.Update(goalsTucked, goalsReceived)
	if err != nil {
		return err
	}
	err = t.Results.Update(goalsTucked, goalsReceived)
	if err != nil {
		return err
	}
	return nil
}

//TeamLocal ... This is the object where we will store and calculate the data to insert in mongodb(local)
type TeamLocal struct {
	Team
}

//TeamAway ... This is the object where we will store and calculate the data to insert in mongodb(away)
type TeamAway struct {
	Team
}
