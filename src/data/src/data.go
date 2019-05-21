package data

import "fmt"

//Config ... Struct created to parse JSON config.json and get information
type Config struct {
	Path     string    `json:"Path"`
	Year     Year      `json:"Year"`
	Endpoint []Country `json:"Endpoints"`
}

//Year ... Struct created to parse JSON config.json and get information
type Year struct {
	From int `json:"from"`
	To   int `json:"to"`
}

//Country ... Strcut created to pase JSON config.json and get information
type Country struct {
	Name string   `json:"name"`
	Keys []string `json:"keys"`
}

//GetYears ... Returns years [from, to] in an array of strings in the format: from(from+1),...
func (y Year) GetYears() []string {
	res := []string{}
	index := 1
	for {
		if index == 19 {
			break
		}
		res = append(res, fmt.Sprintf("%d%d", index, index+1))
		index++
	}
	return res
}
