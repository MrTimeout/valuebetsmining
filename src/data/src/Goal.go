package data

import "errors"

//Goal ... This is where we store the goals of each team
type Goal struct {
	GoalsTucked          []int   `json:"goals tucked"`
	GoalsReceived        []int   `json:"goals received"`
	GoalsTuckedAverage   float64 `json:"goals tucked average"`
	GoalsReceivedAverage float64 `json:"goals received average"`
	GoalsTuckedMode      []int   `json:"goals tucked mode"`
	GoalsReceivedMode    []int   `json:"goals received mode"`
}

//NewGoal ... This is the 'constructor of the struct goal'
func (g *Goal) NewGoal(goalsTucked int, goalsReceived int) (Goal, error) {
	if goalsTucked < 0 || goalsReceived < 0 {
		return Goal{}, errors.New("Error parsing goals")
	}
	goal := Goal{
		GoalsTucked:          []int{goalsTucked},
		GoalsReceived:        []int{goalsReceived},
		GoalsTuckedAverage:   float64(goalsTucked),
		GoalsReceivedAverage: float64(goalsReceived),
		GoalsTuckedMode:      []int{goalsTucked},
		GoalsReceivedMode:    []int{goalsReceived},
	}
	return goal, nil
}

//AppendGoalsTucked ... Append to the GoalsTucked array a new goal.
func (g *Goal) AppendGoalsTucked(goals int) error {
	if goals < 0 {
		return errors.New("Error parsing goals argument")
	}
	if len(g.GoalsTucked) == 10 {
		g.GoalsTucked = append(g.GoalsTucked[1:len(g.GoalsTucked)], goals)
	} else {
		g.GoalsTucked = append(g.GoalsTucked, goals)
	}
	return nil
}

//AppendGoalsReceived ... Append to the GoalsReceived array a new goal.
func (g *Goal) AppendGoalsReceived(goals int) error {
	if goals < 0 {
		return errors.New("Error parsing goals argument")
	}
	if len(g.GoalsReceived) == 10 {
		g.GoalsReceived = append(g.GoalsReceived[1:len(g.GoalsReceived)], goals)
	} else {
		g.GoalsReceived = append(g.GoalsReceived, goals)
	}
	return nil
}

//CalculateGoalsTuckedAverage ... Calculates the average of the goals tucked
func (g *Goal) CalculateGoalsTuckedAverage() error {
	average, err := Average(g.GoalsTucked, false)
	if err != nil {
		return err
	}
	g.GoalsTuckedAverage = average
	return nil
}

//CalculateGoalsReceivedAverage ... Calculates the average of the goals received
func (g *Goal) CalculateGoalsReceivedAverage() error {
	average, err := Average(g.GoalsReceived, false)
	if err != nil {
		return err
	}
	g.GoalsReceivedAverage = average
	return nil
}
