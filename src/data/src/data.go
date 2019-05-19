package data

//Path ... Struct created to parse JSON config.json and get information
type Path struct {
	Path     string    `json:"Path"`
	Endpoint []Country `json:"Endpoints"`
}

//Country ... Strcut created to pase JSON config.json and get information
type Country struct {
	Name string   `json:"name"`
	Keys []string `json:"keys"`
}
