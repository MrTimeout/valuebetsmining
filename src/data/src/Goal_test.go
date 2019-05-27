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
	if len(goals.GoalsReceived) != 10 {
		t.Errorf("Error parsing goals length: %d", len(goals.GoalsReceived))
	} else {
		t.Errorf("GOALS: %#v", goals.GoalsReceived)
	}
}
