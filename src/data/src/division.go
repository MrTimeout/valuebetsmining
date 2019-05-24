package data

//Division ... Each division of each league
type Division struct {
	Years       []int   `json:"Years"`
	Matchs      []Match `json:"Matchs"`
	LengthMatch int     `json:"Amount of matches"`
}

//New ... Creates an instance of the struct Division
func (d Division) New() (Division, error) {

}
