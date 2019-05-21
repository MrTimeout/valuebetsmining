package data

import "errors"

//Team ... This is where we store information about each team
type Team struct {
	Name   string `json:"name"`
	Goals  Goal   `json:"goals"`
	Matchs []int  `json:"matchs"`
}

//InsertMatch ... Insert other result of the team
func (t Team) InsertMatch(a int) error {
	if a > 1 || a < -1 {
		return errors.New("Error parsing result of a team")
	}
	if len(t.Matchs) == 10 {
		t.Matchs = append(t.Matchs, a)
	}
	return nil
}

//TeamLocal ... This is the object where we will store and calculate the data to insert in mongodb(local)
type TeamLocal struct {
	Team
}

//TeamAway ... This is the object where we will store and calculate the data to insert in mongodb(away)
type TeamAway struct {
	Team
}
