package entities

//TeamLocal ... Handles local properties of a local team
type TeamLocal struct {
	Last10WinningMatchs        int     `bson:"Last10WinningLocalMatchs"`
	Last10TiedingMatchs        int     `bson:"Last10TiedingLocalMatchs"`
	Last10LosingMatchs         int     `bson:"Last10LosingLocalMatchs"`
	Last10StreackWinning       int     `bson:"Last10StreackWinningLocal"`
	Last10StreackNoLosing      int     `bson:"Last10StreackNoLosingLocal"`
	Last10GoalsTuckedAmount    int     `bson:"Last10GoalsTuckedAmountLocal"`
	Last10GoalsReceivedAmount  int     `bson:"Last10GoalsReceivedAmountLocal"`
	Last10AverageTuckedGoals   float64 `bson:"Last10AverageTuckedGoalsLocal"`
	Last10AverageReceivedGoals float64 `bson:"Last10AverageReceivedGoalsLocal"`
}

//TeamAway ... Handles away properties of an away team
type TeamAway struct {
	Last10WinningMatchs        int     `bson:"Last10WinningAwayMatchs"`
	Last10TiedingMatchs        int     `bson:"Last10TiedingAwayMatchs"`
	Last10LosingMatchs         int     `bson:"Last10LosingAwayMatchs"`
	Last10StreackWinning       int     `bson:"Last10StreackWinningAway"`
	Last10StreackNoLosing      int     `bson:"Last10StreackNoLosingAway"`
	Last10GoalsTuckedAmount    int     `bson:"Last10GoalsTuckedAmountAway"`
	Last10GoalsReceivedAmount  int     `bson:"Last10GoalsReceivedAmountAway"`
	Last10AverageTuckedGoals   float64 `bson:"Last10AverageTuckedGoalsAway"`
	Last10AverageReceivedGoals float64 `bson:"Last10AverageReceivedGoalsAway"`
}

//Team ... Handles all the information of the team
type Team struct {
	local TeamLocal
	away  TeamAway
}

//NewTeam ... Return a new team with local and away properties
func NewTeam(r Result) Team {
	return Team{
		local: NewLocalTeam(r),
		away:  NewAwayTeam(r),
	}
}

//NewAwayTeam ... Return an instance of am away team using result by param
func NewAwayTeam(r Result) TeamAway {
	return TeamAway{
		Last10WinningMatchs:        r.Last10WinningAwayMatchs,
		Last10TiedingMatchs:        r.Last10TiedingAwayMatchs,
		Last10LosingMatchs:         r.Last10LosingAwayMatchs,
		Last10StreackWinning:       r.Last10StreackWinningAway,
		Last10StreackNoLosing:      r.Last10StreackNoLosingAway,
		Last10GoalsTuckedAmount:    r.Last10GoalsTuckedAmountAway,
		Last10GoalsReceivedAmount:  r.Last10GoalsReceivedAmountAway,
		Last10AverageTuckedGoals:   r.Last10AverageTuckedGoalsAway,
		Last10AverageReceivedGoals: r.Last10AverageReceivedGoalsAway,
	}
}

//NewLocalTeam ... Return an instance of a local team using result by param
func NewLocalTeam(r Result) TeamLocal {
	return TeamLocal{
		Last10WinningMatchs:        r.Last10WinningLocalMatchs,
		Last10TiedingMatchs:        r.Last10TiedingLocalMatchs,
		Last10LosingMatchs:         r.Last10LosingLocalMatchs,
		Last10StreackWinning:       r.Last10StreackWinningLocal,
		Last10StreackNoLosing:      r.Last10StreackNoLosingLocal,
		Last10GoalsTuckedAmount:    r.Last10GoalsTuckedAmountLocal,
		Last10GoalsReceivedAmount:  r.Last10GoalsReceivedAmountLocal,
		Last10AverageTuckedGoals:   r.Last10AverageTuckedGoalsLocal,
		Last10AverageReceivedGoals: r.Last10AverageReceivedGoalsLocal,
	}
}
