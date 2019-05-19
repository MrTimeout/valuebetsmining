package data

import "errors"

//URL ... this is the object that will handle all the information about the Url that we are going to attack
type URL struct {
	protocol string
	domain   string
	port     int
	folders  []string
	params   []Param
}

//Param ...
type Param struct {
	key   string
	value string
}

//New ... This method creates a new instance of the class URL, testing values passed by params
func (u *URL) New(url string) (URL, error) {
	return URL{}, errors.New("new error on url")
}
