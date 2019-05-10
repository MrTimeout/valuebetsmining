package main
 
import "time"
 
type Team struct {
    Name      string    `json:"name"`
    League    string    `json:"League"`
    Due       time.Time `json:"due"`
}
 
type Teams []Team