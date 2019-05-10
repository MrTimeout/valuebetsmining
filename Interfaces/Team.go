package Interfaces
 
type Team struct {
    Name      string    `json:"name"`
    League    string    `json:"League"`
}
 
type Teams []Team