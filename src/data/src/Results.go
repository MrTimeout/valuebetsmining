package data

import "errors"

//Result ... these are the results of each team (local and away)
type Result struct {
	Matchs           []int `json:"matchs"`
	StreackWinning   int   `json:"streack winning"`
	StreackNoLosing  int   `json:"streack without losing"`
	StreackTieding   int   `json:"Streack tieding"`
	StreackNoWinning int   `json:"Streack wihout winning"`
	StreackLosing    int   `json:"streack losing"`
}

//New ... creates a new structure of Results of a team
func (r Result) New(goalsTucked, goalsReceived int) (Result, error) {
	if goalsTucked < 0 || goalsReceived < 0 {
		return Result{}, errors.New("Error parsing goals of result")
	}
	return Result{
		Matchs: []int{WhoIsBigger(goalsTucked, goalsReceived)},
	}, nil
}

//CalStreackWinning ...Calculates the streack winning of the team
func (r Result) CalStreackWinning() {

}
