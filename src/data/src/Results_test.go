package data

import "testing"

func TestResultsFunc(t *testing.T) {
	results, err := NewResults(5, 2)
	if err != nil {
		t.Error(err)
	}
	rands, err := RandomArr(1, 3, 4)
	for _, value := range rands {
		results.Update(value, -value)
	}
	if len(results.Matchs) >= 100 {
		t.Errorf("Error parsing goals length: %d", len(results.Matchs))
	} else {
		t.Errorf("MATCHS: %#v", results.Matchs)
	}
}
