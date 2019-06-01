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

//TestPreviousNResultsOfAMatch ... Testing function PreviousNResultsOfAMatch to visualize if it returns correct values
func TestPreviousNResultsOfAMatch(t *testing.T) {
	r, err := NewResults(1, 2)
	if err != nil {
		t.Error(err)
	}
	randsLocal, err := RandomArr(0, 5, 10)
	randsAway, err := RandomArr(0, 5, 10)
	resultsArr := []Result{r}
	for i := 0; i < len(randsLocal); i++ {
		err = r.Update(randsLocal[i], randsAway[i])
		if err != nil {
			t.Error(err)
			break
		}
		resultsArr = append(resultsArr, r)
	}
	for i := 1; i < len(r.Matchs); i++ {
		resultsPrevious, err := r.PreviousNResultsOfAMatch(i)
		if err != nil {
			t.Error(err)
			break
		}
		if !resultsPrevious.CompareOfResults(resultsArr[len(resultsArr)-i]) {
			t.Errorf("Time:%d\nWant->\n%s\nGot->\n%s\n", i, resultsArr[len(resultsArr)-i].String(), resultsPrevious.String())
			break
		}
	}
}
