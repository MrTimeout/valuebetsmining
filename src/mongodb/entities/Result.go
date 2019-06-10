package entities

import "fmt"

//Result ... Struct that handles information of the last match of each team
type Result struct {
	Index                           int     `bson:"Index"`
	Division                        string  `bson:"Division"`
	Date                            string  `bson:"Date"`
	LocalTeam                       string  `bson:"LocalTean"`
	AwayTeam                        string  `bson:"AwayTeam"`
	Result                          string  `bson:"Result"`
	LocalGoals                      int     `bson:"LocalGoals"`
	AwayGoals                       int     `bson:"AwayGoals"`
	Last10WinningLocalMatchs        int     `bson:"Last10WinningLocalMatchs"`
	Last10TiedingLocalMatchs        int     `bson:"Last10TiedingLocalMatchs"`
	Last10LosingLocalMatchs         int     `bson:"Last10LosingLocalMatchs"`
	Last10WinningAwayMatchs         int     `bson:"Last10WinningAwayMatchs"`
	Last10TiedingAwayMatchs         int     `bson:"Last10TiedingAwayMatchs"`
	Last10LosingAwayMatchs          int     `bson:"Last10LosingAwayMatchs"`
	Last10StreackWinningLocal       int     `bson:"Last10StreackWinningLocal"`
	Last10StreackNoLosingLocal      int     `bson:"Last10StreackNoLosingLocal"`
	Last10StreackWinningAway        int     `bson:"Last10StreackWinningAway"`
	Last10StreackNoLosingAway       int     `bson:"Last10StreackNoLosingAway"`
	Last10GoalsTuckedAmount         int     `bson:"Last10GoalsTuckedAmount"`
	Last10GoalsReceivedAmount       int     `bson:"Last10GoalsReceivedAmount"`
	Last10AverageTuckedGoalsLocal   float64 `bson:"Last10AverageTuckedGoalsLocal"`
	Last10AverageReceivedGoalsLocal float64 `bson:"Last10AverageReceivedGoalsLocal"`
	Last10AverageTuckedGoalsAway    float64 `bson:"Last10AverageTuckedGoalsAway"`
	Last10AverageReceivedGoalsAway  float64 `bson:"Last10AverageReceivedGoalsAway"`
}

//String ... Return all the properties in a easy way to understand
func (r *Result) String() string {
	return fmt.Sprintf("Index: %d\nDiv: %s\nDate: %s\nTeams: %s - %s\nResult: %s\n Goals: %d - %d\nLast10LocalMatchs: %d - %d - %d --- [w, t, l] --- %d - %d - %d\n Last10Streack: %d - %d --- [w, noL] --- %d - %d\n Last10AverageGoals: %2.f - %2.f --- [Tucked, received] --- %2.f - %2.f",
		r.Index,
		r.Division,
		r.Date,
		r.LocalTeam,
		r.AwayTeam,
		r.Result,
		r.LocalGoals,
		r.AwayGoals,
		r.Last10WinningLocalMatchs,
		r.Last10TiedingLocalMatchs,
		r.Last10LosingLocalMatchs,
		r.Last10WinningAwayMatchs,
		r.Last10TiedingAwayMatchs,
		r.Last10LosingAwayMatchs,
		r.Last10StreackWinningLocal,
		r.Last10StreackNoLosingLocal,
		r.Last10StreackWinningAway,
		r.Last10StreackNoLosingAway,
		r.Last10AverageTuckedGoalsLocal,
		r.Last10AverageReceivedGoalsLocal,
		r.Last10AverageTuckedGoalsAway,
		r.Last10AverageReceivedGoalsAway,
	)
}
