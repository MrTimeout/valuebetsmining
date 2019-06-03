package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	DefaultDirTestResults  = "test/Results"
	DefaultNameTestResults = "Testing"
	DefaultExtensionCSV    = ".csv"
)

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

func TestNewResults(t *testing.T) {
	r, err := NewResults(0, 1)
	if err != nil {
		t.Error(err)
	}
	randGoalsLocal, err := RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	randsGoalsAway, err := RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	data := make([][]string, 0, 100)
	for i := 0; i < len(randGoalsLocal); i++ {
		r.Update(randGoalsLocal[i], randsGoalsAway[i])
		data = append(data, strings.Split(fmt.Sprintf("%d,%d,%d,%s", i, randGoalsLocal[i], randsGoalsAway[i], r.StringCSV(DefaultLenResult)), ","))
	}
	file, err := os.Create(fmt.Sprintf("%s/%s%s_%d_%d%s", DefaultDirTestResults, DefaultNameTestResults, FuncName(), time.Now().Hour(), time.Now().Minute(), DefaultExtensionCSV))
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Index", "goals local", "goals away", "local winnings", "local tiedings", "local losings"})

	writer.WriteAll(data)
}

func TestNewResultsStringCSVAll(t *testing.T) {
	r, err := NewResults(0, 1)
	if err != nil {
		t.Error(err)
	}
	randGoalsLocal, err := RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	randsGoalsAway, err := RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	data := make([][]string, 0, 100)
	for i := 0; i < len(randGoalsLocal); i++ {
		r.Update(randGoalsLocal[i], randsGoalsAway[i])
		data = append(data, strings.Split(fmt.Sprintf("%d,%d,%d,%s,%d,%d", i, randGoalsLocal[i], randsGoalsAway[i], r.StringCSV(DefaultLenResult), r.StreackWinning, r.StreackNoLosing), ","))
	}
	file, err := os.Create(fmt.Sprintf("%s/%s%s_%d_%d%s", DefaultDirTestResults, DefaultNameTestResults, FuncName(), time.Now().Hour(), time.Now().Minute(), DefaultExtensionCSV))
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Index", "goals local", "goals away", "local winnings", "local tiedings", "local losings", "local winning streack", "local no losing streack"})

	writer.WriteAll(data)
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

func TestCalsFeatures(t *testing.T) {

}
