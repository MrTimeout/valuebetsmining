package entities

//Team ... This is where we store information about each team
type Team struct {
	Name    string `json:"name"`
	Goals   Goal   `json:"goals"`
	Results Result `json:"results"`
}

const (
	//DefaultPreviousTeam ... Default previous team
	DefaultPreviousTeam = 1
)

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

//PreviousNTeamOfAMatch ... Creates a instance of n prebious matchs
func (t *Team) PreviousNTeamOfAMatch(n int) (Team, error) {
	if n < 0 {
		return Team{}, ErrIndexOutOfRangeResult
	}
	if n == 0 {
		n = DefaultPreviousTeam
	}
	goal, err := t.Goals.PreviousNGoalsOfAMatch(n)
	if err != nil {
		return Team{}, err
	}
	results, err := t.Results.PreviousNResultsOfAMatch(n)
	if err != nil {
		return Team{}, err
	}
	return Team{
		Name:    t.Name,
		Goals:   goal,
		Results: results,
	}, nil
}

//Update ... Update values of the properties of each team
func (t *Team) Update(goalsTucked, goalsReceived int) error {
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
