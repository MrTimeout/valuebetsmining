package data

//Division ... Each division of each league
type Division struct {
	Name        string  `json:"Name"`
	Matchs      []Match `json:"Matchs"`
	LengthMatch int     `json:"Amount of matches"`
}

//UpdateLengthMatch ... Update the length of the match inside the struct
func (d Division) UpdateLengthMatch() {
	d.LengthMatch = len(d.Matchs)
}
