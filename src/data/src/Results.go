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
func (r Result) CalStreackWinning() error {
	res, err := HowManyTimes(r.Matchs, 1)
	if err != nil {
		return err
	}
	r.StreackWinning = res
	return nil
}

//CalStreackNoLosing ...Calculates the streack without losin of the team
func (r Result) CalStreackNoLosing() error {
	res, err := HowManyTimes(r.Matchs, 1, 0)
	if err != nil {
		return err
	}
	r.StreackNoLosing = res
	return nil
}

//CalStreackTieding ...Calculates the streack tieding of the team
func (r Result) CalStreackTieding() error {
	res, err := HowManyTimes(r.Matchs, 0)
	if err != nil {
		return err
	}
	r.StreackTieding = res
	return nil
}

//CalStreackLosing ...Calculates the streack losing of the team
func (r Result) CalStreackLosing() error {
	res, err := HowManyTimes(r.Matchs, -1)
	if err != nil {
		return err
	}
	r.StreackLosing = res
	return nil
}

//CalStreackNoWinning ...Calculates the streack no winning of the team
func (r Result) CalStreackNoWinning() error {
	res, err := HowManyTimes(r.Matchs, -1, 0)
	if err != nil {
		return err
	}
	r.StreackNoWinning = res
	return nil
}
