package data

import "fmt"

//Goal ... This is where we store the goals of each team
type Goal struct {
	GoalsTucked          []int   `json:"goals tucked"`
	GoalsReceived        []int   `json:"goals received"`
	GoalsTuckedAverage   float64 `json:"goals tucked average"`
	GoalsReceivedAverage float64 `json:"goals received average"`
	GoalsTuckedMode      []int   `json:"goals tucked mode"`
	GoalsReceivedMode    []int   `json:"goals received mode"`
}

//GoalError ... Struct that handles errors with Goal struct
type GoalError struct {
	ErrorString string
}

const (
	//DefaultLen ... Default length of the arrays of struct goals
	DefaultLen = 10
	//DefaultPreviousGoals ... Default number of matchs before actual one
	DefaultPreviousGoals = 1
)

var (
	//ErrNegative ... Error parsing goals because they are less than 0
	ErrNegative = &GoalError{ErrorString: "Error parsing goals because they are negative"}
	//ErrIndexOutOfGoal ... Error parsing goals of a match that doesnt exit
	ErrIndexOutOfGoal = &GoalError{ErrorString: "Error parsing goals of a match that doesnt exit"}
	//ErrIndexOutOfRange ... Error parsing array of goals because of a negative array
	ErrIndexOutOfRange = &GoalError{ErrorString: "Error parsing array of goals because of a negative number"}
)

//NewGoal ... This is the 'constructor of the struct goal'
func NewGoal(goalsTucked int, goalsReceived int) (Goal, error) {
	if goalsTucked < 0 || goalsReceived < 0 {
		return Goal{}, ErrNegative
	}
	return Goal{
		GoalsTucked:          []int{goalsTucked},
		GoalsReceived:        []int{goalsReceived},
		GoalsTuckedAverage:   float64(goalsTucked),
		GoalsReceivedAverage: float64(goalsReceived),
		GoalsTuckedMode:      []int{goalsTucked},
		GoalsReceivedMode:    []int{goalsReceived},
	}, nil
}

//Update ... Updates the propertues of Goal object
func (g *Goal) Update(goalsTucked, goalsReceived int) error {
	if goalsTucked < 0 || goalsReceived < 0 {
		return ErrNegative
	}
	err := g.AppendGoalsReceived(goalsReceived)
	if err != nil {
		return err
	}
	err = g.AppendGoalsTucked(goalsTucked)
	if err != nil {
		return err
	}
	return nil
}

//AppendGoalsTucked ... Append to the GoalsTucked array a new goal.
func (g *Goal) AppendGoalsTucked(goals int) error {
	if goals < 0 {
		return ErrNegative
	}
	if len(g.GoalsTucked) == DefaultLen {
		g.GoalsTucked = append(g.GoalsTucked[1:len(g.GoalsTucked)], goals)
	} else {
		g.GoalsTucked = append(g.GoalsTucked, goals)
	}
	err := g.CalculateGoalsTuckedAverage()
	if err != nil {
		return err
	}
	err = g.CalculateGoalsTuckedMode()
	if err != nil {
		return err
	}
	return nil
}

//AppendGoalsReceived ... Append to the GoalsReceived array a new goal.
func (g *Goal) AppendGoalsReceived(goals int) error {
	if goals < 0 {
		return ErrNegative
	}
	if len(g.GoalsReceived) == DefaultLen {
		g.GoalsReceived = append(g.GoalsReceived[1:len(g.GoalsReceived)], goals)
	} else {
		g.GoalsReceived = append(g.GoalsReceived, goals)
	}
	err := g.CalculateGoalsReceivedAverage()
	if err != nil {
		return err
	}
	err = g.CalculateGoalsReceivedMode()
	if err != nil {
		return err
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

//CalculateGoalsTuckedMode ... Calculates the mode of the goals tucked
func (g *Goal) CalculateGoalsTuckedMode() error {
	mode, err := Mode(g.GoalsTucked...)
	if err != nil {
		return err
	}
	g.GoalsTuckedMode = mode
	return nil
}

//CalculateGoalsReceivedMode ... Calculates the mode of the goals received
func (g *Goal) CalculateGoalsReceivedMode() error {
	mode, err := Mode(g.GoalsReceived...)
	if err != nil {
		return err
	}
	g.GoalsReceivedMode = mode
	return nil
}

//PreviousNGoalsOfAMatch ... Take an object with values of an specific match n previous matchs
func (g *Goal) PreviousNGoalsOfAMatch(n int) (Goal, error) {
	if n < 0 {
		return Goal{}, ErrIndexOutOfRange
	}
	if n == 0 {
		n = DefaultPreviousGoals
	}
	diff := len(g.GoalsTucked) - n
	if diff < 0 {
		return Goal{}, ErrIndexOutOfGoal
	}
	averageTucked, err := Average(g.GoalsTucked[:diff], false)
	if err != nil {
		return Goal{}, err
	}
	averageReceived, err := Average(g.GoalsReceived[:diff], false)
	if err != nil {
		return Goal{}, err
	}
	modeTucked, err := Mode(g.GoalsTucked[:diff]...)
	if err != nil {
		return Goal{}, err
	}
	modeReceived, err := Mode(g.GoalsReceived[:diff]...)
	if err != nil {
		return Goal{}, err
	}
	return Goal{
		GoalsTucked:          []int{g.GoalsTucked[diff]},
		GoalsReceived:        []int{g.GoalsReceived[diff]},
		GoalsTuckedAverage:   averageTucked,
		GoalsReceivedAverage: averageReceived,
		GoalsTuckedMode:      modeTucked,
		GoalsReceivedMode:    modeReceived,
	}, nil
}

//CompareOfGoals ... Compare actual goal struct and other passed by param
func (g *Goal) CompareOfGoals(g2 Goal) bool {
	return CompareTwoArrs(g.GoalsReceived, g2.GoalsReceived) &&
		CompareTwoArrs(g.GoalsTucked, g2.GoalsTucked) &&
		CompareTwoArrs(g.GoalsReceivedMode, g2.GoalsReceivedMode) &&
		CompareTwoArrs(g.GoalsTuckedMode, g2.GoalsTuckedMode) &&
		g.GoalsReceivedAverage == g2.GoalsReceivedAverage &&
		g.GoalsTuckedAverage == g2.GoalsTuckedAverage
}

//String ... Return a form more readable of the struct goal
func (g *Goal) String() string {
	return fmt.Sprintf("%")
}

//Error ... Returns the error of the struct goal
func (g *GoalError) Error() string {
	return g.ErrorString
}
