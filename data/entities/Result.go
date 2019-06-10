package entities

import (
	"fmt"
	others "valuebetsmining/data/Others"
)

//Result ... these are the results of each team (local and away)
type Result struct {
	Matchs           []int `json:"matchs"`
	StreackWinning   int   `json:"streack winning"`
	StreackNoLosing  int   `json:"streack without losing"`
	StreackTieding   int   `json:"Streack tieding"`
	StreackNoWinning int   `json:"Streack wihout winning"`
	StreackLosing    int   `json:"streack losing"`
}

//ResultError ... This struct handles error of struct Result
type ResultError struct {
	ErrorString string
}

const (
	//DefaultLenResult ... Default length of the arrays of struct result
	DefaultLenResult = 10
	//DefaultPreviousResults ... Default number of matchs before actual one
	DefaultPreviousResults = 1
)

var (
	//ErrNegativeResult ... Error parsing results because they are less than 0
	ErrNegativeResult = &ResultError{ErrorString: "Error parsing results because they are negative"}
	//ErrNilArrResults ... Error parsing array of results because it is nil
	ErrNilArrResults = &ResultError{ErrorString: "Error parsing array of results because it is nil"}
	//ErrTypeResults ... Error parsing array of results because it contains incorrect values
	ErrTypeResults = &ResultError{ErrorString: "Error parsing array of results because it contains incorrect values"}
	//ErrIndexOutOfResult ... Error parsing results of a match that doesnt exit
	ErrIndexOutOfResult = &ResultError{ErrorString: "Error parsing result of a match that doesnt exit"}
	//ErrIndexOutOfRangeResult ... Error parsing array of result because of a negative array
	ErrIndexOutOfRangeResult = &ResultError{ErrorString: "Error parsing array of result because of a negative number"}
)

//NewResults ... creates a new structure of Results of a team
func NewResults(goalsTucked, goalsReceived int) (Result, error) {
	if goalsTucked < 0 || goalsReceived < 0 {
		return Result{}, ErrNegativeGoal
	}
	r := Result{
		Matchs: []int{others.WhoIsBigger(goalsTucked, goalsReceived)},
	}
	err := r.CalFeatures()
	if err != nil {
		return Result{}, err
	}
	return r, nil
}

//NewResultsMatchs ... Return an instance of Result using matchs by param
func NewResultsMatchs(matchs []int) (Result, error) {
	if matchs == nil || len(matchs) == 0 {
		return Result{}, ErrNilArrResults
	}
	if t, err := others.IsAStrangerHere(matchs, []int{1, 0, -1}); err != nil {
		return Result{}, err
	} else if t {
		return Result{}, ErrTypeResults
	}
	r := Result{Matchs: matchs}
	err := r.CalFeatures()
	if err != nil {
		return Result{}, err
	}
	return r, nil
}

//CalFeatures ... Calculate all the features of the struct result
func (r *Result) CalFeatures() error {
	err := r.CalStreackLosing()
	if err != nil {
		return err
	}
	err = r.CalStreackNoLosing()
	if err != nil {
		return err
	}
	err = r.CalStreackNoWinning()
	if err != nil {
		return err
	}
	err = r.CalStreackTieding()
	if err != nil {
		return err
	}
	err = r.CalStreackWinning()
	if err != nil {
		return err
	}
	return nil
}

//Update ... Insert a new match
func (r *Result) Update(goalsTucked, goalsReceived int) error {
	if goalsTucked < 0 || goalsReceived < 0 {
		return ErrNegativeGoal
	}
	r.Matchs = append(r.Matchs, others.WhoIsBigger(goalsTucked, goalsReceived))
	err := r.CalFeatures()
	if err != nil {
		return err
	}
	return nil
}

//CalStreackWinning ...Calculates the streack winning of the team
func (r *Result) CalStreackWinning() error {
	var res int
	var err error
	if len(r.Matchs) >= DefaultLenResult {
		res, err = others.HowManyTimes(r.Matchs[len(r.Matchs)-DefaultLenResult:], true, 1)
	} else {
		res, err = others.HowManyTimes(r.Matchs, true, 1)
	}
	if err != nil {
		return err
	}
	r.StreackWinning = res
	return nil
}

//CalStreackNoLosing ...Calculates the streack without losin of the team
func (r *Result) CalStreackNoLosing() error {
	var res int
	var err error
	if len(r.Matchs) >= DefaultLenResult {
		res, err = others.HowManyTimes(r.Matchs[len(r.Matchs)-DefaultLenResult:], true, 1, 0)
	} else {
		res, err = others.HowManyTimes(r.Matchs, true, 1, 0)
	}
	if err != nil {
		return err
	}
	r.StreackNoLosing = res
	return nil
}

//CalStreackTieding ...Calculates the streack tieding of the team
func (r *Result) CalStreackTieding() error {
	var res int
	var err error
	if len(r.Matchs) >= DefaultLenResult {
		res, err = others.HowManyTimes(r.Matchs[len(r.Matchs)-DefaultLenResult:], true, 0)
	} else {
		res, err = others.HowManyTimes(r.Matchs, true, 0)
	}
	if err != nil {
		return err
	}
	r.StreackTieding = res
	return nil
}

//CalStreackNoWinning ...Calculates the streack no winning of the team
func (r *Result) CalStreackNoWinning() error {
	var res int
	var err error
	if len(r.Matchs) >= DefaultLenResult {
		res, err = others.HowManyTimes(r.Matchs[len(r.Matchs)-DefaultLenResult:], true, -1, 0)
	} else {
		res, err = others.HowManyTimes(r.Matchs, true, -1, 0)
	}
	if err != nil {
		return err
	}
	r.StreackNoWinning = res
	return nil
}

//CalStreackLosing ...Calculates the streack losing of the team
func (r *Result) CalStreackLosing() error {
	var res int
	var err error
	if len(r.Matchs) >= DefaultLenResult {
		res, err = others.HowManyTimes(r.Matchs[len(r.Matchs)-DefaultLenResult:], true, -1)
	} else {
		res, err = others.HowManyTimes(r.Matchs, true, -1)
	}
	if err != nil {
		return err
	}
	r.StreackLosing = res
	return nil
}

//WinTieLose ... Returns a map with won, tieded and lost matchs
func (r *Result) WinTieLose(maxIndex int) map[string]int {
	if maxIndex > len(r.Matchs) {
		maxIndex = len(r.Matchs)
	}
	slices := r.Matchs[len(r.Matchs)-maxIndex:]
	mapWTL := make(map[string]int)
	for _, v := range slices {
		switch v {
		case -1:
			mapWTL["lost"] = mapWTL["lost"] + 1
		case 0:
			mapWTL["tied"] = mapWTL["tied"] + 1
		case 1:
			mapWTL["won"] = mapWTL["won"] + 1
		}
	}
	return mapWTL
}

//PreviousNResultsOfAMatch ... Take an object with values of an specific match n previous matchs
func (r *Result) PreviousNResultsOfAMatch(n int) (Result, error) {
	if n < 0 {
		return Result{}, ErrIndexOutOfRangeResult
	}
	if n == 0 {
		n = DefaultPreviousResults
	}
	diff := len(r.Matchs) - n
	if diff < 0 {
		return Result{}, ErrIndexOutOfResult
	}
	var rr []int
	if diff+1 >= DefaultLenResult {
		rr = r.Matchs[diff+1-DefaultLenResult : diff+1]
	} else {
		rr = r.Matchs[:diff+1]
	}
	result, err := NewResultsMatchs(rr)
	if err != nil {
		return Result{}, err
	}
	err = result.CalFeatures()
	if err != nil {
		return Result{}, err
	}
	result.Matchs = r.Matchs[:diff+1]
	return result, nil
}

//CompareOfResults ... Compare of two instances of struct Result
func (r *Result) CompareOfResults(r2 Result) bool {
	return others.CompareTwoArrs(r.Matchs, r2.Matchs, false) &&
		r.StreackLosing == r2.StreackLosing &&
		r.StreackNoLosing == r2.StreackNoLosing &&
		r.StreackTieding == r2.StreackTieding &&
		r.StreackNoWinning == r2.StreackNoWinning &&
		r.StreackWinning == r2.StreackWinning
}

func (r *Result) String() string {
	mapWTL := r.WinTieLose(DefaultLenResult)
	return fmt.Sprintf("Matchs:\n\tWon: %d\n\tTied: %d\n\tLost: %d \nStreack losing: %d\nStreack no losing: %d\nStreack tieding: %d\nStreack no winning: %d\nStreack winning: %d\n", mapWTL["won"], mapWTL["tied"], mapWTL["lost"], r.StreackLosing, r.StreackNoLosing, r.StreackTieding, r.StreackNoWinning, r.StreackWinning)
}

//StringCSV ... Return a string of attr of struct result
func (r *Result) StringCSV(maxIndex int) string {
	mapWTL := r.WinTieLose(maxIndex)
	return fmt.Sprintf("%d,%d,%d", mapWTL["won"], mapWTL["tied"], mapWTL["lost"])
}

//Error ... Return error of struct Result
func (re ResultError) Error() string {
	return re.ErrorString
}
