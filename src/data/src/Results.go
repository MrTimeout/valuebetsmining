package data

import (
	"errors"
	"fmt"
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
		return Result{}, errors.New("Error parsing goals of result")
	}
	r := Result{
		Matchs: []int{WhoIsBigger(goalsTucked, goalsReceived)},
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
	if t, err := IsAStrangerHere(matchs, []int{1, 0, -1}); err != nil {
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
		return errors.New("Error parsing goals of result")
	}
	if len(r.Matchs) == 10 {
		r.Matchs = append(r.Matchs[1:], WhoIsBigger(goalsTucked, goalsReceived))
	} else {
		r.Matchs = append(r.Matchs, WhoIsBigger(goalsTucked, goalsReceived))
	}
	err := r.CalFeatures()
	if err != nil {
		return err
	}
	return nil
}

//CalStreackWinning ...Calculates the streack winning of the team
func (r *Result) CalStreackWinning() error {
	res, err := HowManyTimes(r.Matchs, 1)
	if err != nil {
		return err
	}
	r.StreackWinning = res
	return nil
}

//CalStreackNoLosing ...Calculates the streack without losin of the team
func (r *Result) CalStreackNoLosing() error {
	res, err := HowManyTimes(r.Matchs, 1, 0)
	if err != nil {
		return err
	}
	r.StreackNoLosing = res
	return nil
}

//CalStreackTieding ...Calculates the streack tieding of the team
func (r *Result) CalStreackTieding() error {
	res, err := HowManyTimes(r.Matchs, 0)
	if err != nil {
		return err
	}
	r.StreackTieding = res
	return nil
}

//CalStreackLosing ...Calculates the streack losing of the team
func (r *Result) CalStreackLosing() error {
	res, err := HowManyTimes(r.Matchs, -1)
	if err != nil {
		return err
	}
	r.StreackLosing = res
	return nil
}

//CalStreackNoWinning ...Calculates the streack no winning of the team
func (r *Result) CalStreackNoWinning() error {
	res, err := HowManyTimes(r.Matchs, -1, 0)
	if err != nil {
		return err
	}
	r.StreackNoWinning = res
	return nil
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
	return result, nil
}

func (r *Result) String() string {
	return fmt.Sprintf("Matchs: %v") //Incomplete
}

//Error ... Return error of struct Result
func (re ResultError) Error() string {
	return re.ErrorString
}
