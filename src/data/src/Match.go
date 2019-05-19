package data

//Match ... This is where we are going to handle match between local team and away team
type Match struct {
	TeamLocal `json:"local team"`
	TeamAway  `json:"away team"`
	GoalAway  int `json:"local goals"`
	GoalLocal int `json:"away goals"`
	Result    int `json:"result"`
}
