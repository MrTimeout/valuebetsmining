package entities

import (
	"fmt"
	others "valuebetsmining/data/Others"
)

//Goal ... This is where we store the goals of each team
type Goal struct {
	GoalsTucked          []int   `json:"goals tucked"`
	GoalsReceived        []int   `json:"goals received"`
	GoalsTuckedAmount    int     `json:"goals tucked amount"`
	GoalsReceivedAmount  int     `json:"goals received amount"`
	GoalsTuckedAverage   float64 `json:"goals tucked average"`
	GoalsReceivedAverage float64 `json:"goals received average"`
	GoalsTuckedMode      []int   `json:"goals tucked mode"`
	GoalsReceivedMode    []int   `json:"goals received mode"`
}

//GoalError ... Struct that handles errors with Goal struct
type GoalError struct {
	ErrorString string
}

var (
	//DefaultLenGoal ... Default length of the arrays of struct goals
	DefaultLenGoal = 10
	//DefaultPreviousGoals ... Default number of matchs before actual one
	DefaultPreviousGoals = 1
)

var (
	//ErrNegativeGoal ... Error parsing goals because they are less than 0
	ErrNegativeGoal = &GoalError{ErrorString: "Error parsing goals because they are negative"}
	//ErrIndexOutOfGoal ... Error parsing goals of a match that doesnt exit
	ErrIndexOutOfGoal = &GoalError{ErrorString: "Error parsing goals of a match that doesnt exit"}
	//ErrIndexOutOfRangeGoal ... Error parsing array of goals because of a negative array
	ErrIndexOutOfRangeGoal = &GoalError{ErrorString: "Error parsing array of goals because of a negative number"}
)

//NewGoal ... This is the 'constructor of the struct goal'
func NewGoal(goalsTucked int, goalsReceived int) (Goal, error) {
	if goalsTucked < 0 || goalsReceived < 0 {
		return Goal{}, ErrNegativeGoal
	}
	return Goal{
		GoalsTucked:          []int{goalsTucked},
		GoalsReceived:        []int{goalsReceived},
		GoalsTuckedAmount:    goalsTucked,
		GoalsReceivedAmount:  goalsReceived,
		GoalsTuckedAverage:   float64(goalsTucked),
		GoalsReceivedAverage: float64(goalsReceived),
		GoalsTuckedMode:      []int{goalsTucked},
		GoalsReceivedMode:    []int{goalsReceived},
	}, nil
}

//Update ... Updates the propertues of Goal object
func (g *Goal) Update(goalsTucked, goalsReceived int) error {
	if goalsTucked < 0 || goalsReceived < 0 {
		return ErrNegativeGoal
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
		return ErrNegativeGoal
	}
	g.GoalsTucked = append(g.GoalsTucked, goals)
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
		return ErrNegativeGoal
	}
	g.GoalsReceived = append(g.GoalsReceived, goals)
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

//CalculateProperties ... Calculate properties in general
func (g *Goal) CalculateProperties() error {
	g.CalculateGoalsReceivedAmount()
	if err := g.CalculateGoalsReceivedAverage(); err != nil {
		return err
	}
	if err := g.CalculateGoalsReceivedMode(); err != nil {
		return err
	}
	g.CalculateGoalsTuckedAmount()
	if err := g.CalculateGoalsTuckedAverage(); err != nil {
		return err
	}
	if err := g.CalculateGoalsTuckedMode(); err != nil {
		return err
	}
	return nil
}

//CalculateGoalsTuckedAmount ... Calculates the amount of the tucked goals
func (g *Goal) CalculateGoalsTuckedAmount() {
	g.GoalsTuckedAmount = others.Sum(false, g.GoalsTucked...)
}

//CalculateGoalsReceivedAmount ... Calculates the amount of the tucked goals
func (g *Goal) CalculateGoalsReceivedAmount() {
	g.GoalsReceivedAmount = others.Sum(false, g.GoalsReceived...)
}

//CalculateGoalsTuckedAverage ... Calculates the average of the goals tucked
func (g *Goal) CalculateGoalsTuckedAverage() error {
	var average float64
	var err error
	if len(g.GoalsTucked) >= DefaultLenGoal {
		average, err = others.Average(g.GoalsTucked[len(g.GoalsTucked)-DefaultLenGoal:len(g.GoalsTucked)], false)
	} else {
		average, err = others.Average(g.GoalsTucked, false)
	}
	if err != nil {
		return err
	}
	g.GoalsTuckedAverage = average
	return nil
}

//CalculateGoalsReceivedAverage ... Calculates the average of the goals received
func (g *Goal) CalculateGoalsReceivedAverage() error {
	var average float64
	var err error
	if len(g.GoalsReceived) >= DefaultLenGoal {
		average, err = others.Average(g.GoalsReceived[len(g.GoalsReceived)-DefaultLenGoal:], false)
	} else {
		average, err = others.Average(g.GoalsReceived, false)
	}
	if err != nil {
		return err
	}
	g.GoalsReceivedAverage = average
	return nil
}

//CalculateGoalsTuckedMode ... Calculates the mode of the goals tucked
func (g *Goal) CalculateGoalsTuckedMode() error {
	var mode []int
	var err error
	if len(g.GoalsTucked) >= DefaultLenGoal {
		mode, err = others.Mode(g.GoalsTucked[len(g.GoalsTucked)-DefaultLenGoal : len(g.GoalsTucked)]...)
	} else {
		mode, err = others.Mode(g.GoalsTucked...)
	}
	if err != nil {
		return err
	}
	g.GoalsTuckedMode = mode
	return nil
}

//CalculateGoalsReceivedMode ... Calculates the mode of the goals received
func (g *Goal) CalculateGoalsReceivedMode() error {
	var mode []int
	var err error
	if len(g.GoalsReceived) >= DefaultLenGoal {
		mode, err = others.Mode(g.GoalsReceived[len(g.GoalsReceived)-DefaultLenGoal : len(g.GoalsReceived)]...)
	} else {
		mode, err = others.Mode(g.GoalsReceived...)
	}
	if err != nil {
		return err
	}
	g.GoalsReceivedMode = mode
	return nil
}

//PreviousNGoalsOfAMatch ... Take an object with values of an specific match n previous matchs
func (g *Goal) PreviousNGoalsOfAMatch(n int) (Goal, error) {
	if n < 0 {
		return Goal{}, ErrIndexOutOfRangeGoal
	}
	if n == 0 {
		n = DefaultPreviousGoals
	}
	if len(g.GoalsTucked) == 1 {
		return Goal{}, ErrIndexOutOfGoal
	}
	diff := len(g.GoalsTucked) - n //9 - 1
	if diff < 0 {
		return Goal{}, ErrIndexOutOfGoal
	}
	var gt, gr []int
	if diff >= DefaultLenGoal { //1 >= 1
		gt = g.GoalsTucked[diff-DefaultLenGoal : len(g.GoalsTucked)-DefaultPreviousGoals]
		gr = g.GoalsReceived[diff-DefaultLenGoal : len(g.GoalsReceived)-DefaultPreviousGoals]
	} else {
		gt = g.GoalsTucked[0 : len(g.GoalsTucked)-DefaultPreviousGoals]
		gr = g.GoalsReceived[0 : len(g.GoalsReceived)-DefaultPreviousGoals]
	}
	gTemp := Goal{
		GoalsTucked:   gt,
		GoalsReceived: gr,
	}
	err := gTemp.CalculateProperties()
	if err != nil {
		return Goal{}, err
	}
	return gTemp, nil
}

//CompareOfGoals ... Compare actual goal struct and other passed by param
func (g *Goal) CompareOfGoals(g2 Goal) bool {
	return others.CompareTwoArrs(g.GoalsReceived, g2.GoalsReceived, false) &&
		others.CompareTwoArrs(g.GoalsTucked, g2.GoalsTucked, false) &&
		others.CompareTwoArrs(g.GoalsReceivedMode, g2.GoalsReceivedMode, true) &&
		others.CompareTwoArrs(g.GoalsTuckedMode, g2.GoalsTuckedMode, true) &&
		g.GoalsReceivedAverage == g2.GoalsReceivedAverage &&
		g.GoalsTuckedAverage == g2.GoalsTuckedAverage
}

//String ... Return a form more readable of the struct goal
func (g *Goal) String() string {
	return fmt.Sprintf("Goals Tucked: %v\nGoals Received: %v\nAverage of goals tucked: %f\nAverage of goals received: %f\nMode of goals tucked: %v\nMode of goals received: %v\n", g.GoalsTucked, g.GoalsReceived, g.GoalsTuckedAverage, g.GoalsReceivedAverage, g.GoalsTuckedMode, g.GoalsReceivedMode)
}

//StringCSV ... Return a string of attr of struct goal
func (g *Goal) StringCSV() string {
	return fmt.Sprintf("%.2f,%.2f", g.GoalsTuckedAverage, g.GoalsReceivedAverage)
}

//Error ... Returns the error of the struct goal
func (g *GoalError) Error() string {
	return g.ErrorString
}
