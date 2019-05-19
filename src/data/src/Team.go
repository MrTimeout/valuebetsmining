package data

//Team ... This is where we store information about each team
type Team struct {
	Name   string `json:"name"`
	Goals  Goal   `json:"goals"`
	Matchs []int  `json:"matchs"`
}

//TeamLocal ... This is the object where we will store and calculate the data to insert in mongodb(local)
type TeamLocal struct {
	Name  string `json:"name"`
	Goals []int  `json:"goals"`
	Match []int  `json:"matchs"`
}

//TeamAway ... This is the object where we will store and calculate the data to insert in mongodb(away)
type TeamAway struct {
	Name  string `json:"name"`
	Goals []int  `json:"goals"`
	Match []int  `json:"matchs"`
}
