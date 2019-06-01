package data

import (
	"testing"
)

func TestGoalsFunc(t *testing.T) {
	goals, err := NewGoal(1, 2)
	if err != nil {
		t.Error(err)
	}
	rands, err := RandomArr(0, 10, 10)
	for key, value := range rands {
		goals.Update(value, key)
	}
	if len(goals.GoalsReceived) <= 10 {
		t.Errorf("Error parsing goals length: %d", len(goals.GoalsReceived))
	}
}

//TestPreviousNGoalsOfAMatch ... Testing function PreviousNGoalsOfAMatch to visualize if it returns correct values
func TestPreviousNGoalsOfAMatch(t *testing.T) {
	goals, err := NewGoal(1, 2)
	if err != nil {
		t.Error(err)
	}
	randsLocal, err := RandomArr(0, 10, 10)
	randsAway, err := RandomArr(0, 10, 10)
	goalsArr := []Goal{goals}
	for i := 0; i < len(randsLocal); i++ {
		err = goals.Update(randsLocal[i], randsAway[i])
		if err != nil {
			t.Error(err)
			break
		}
		goalsArr = append(goalsArr, goals)
	}
	for i := 1; i < len(goals.GoalsReceived); i++ {
		goalsPrevious, err := goals.PreviousNGoalsOfAMatch(i)
		if err != nil {
			t.Error(err)
			break
		}
		if !goalsPrevious.CompareOfGoals(goalsArr[len(goalsArr)-i]) {
			t.Errorf("Time:%d\nWant->\n%s\nGot->\n%s\n", i, goalsArr[len(goalsArr)-i].String(), goalsPrevious.String())
			break
		}
	}
}
