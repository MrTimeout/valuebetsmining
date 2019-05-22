package data

import (
	"log"
	"testing"
)

func TestExistsDivision(t *testing.T) {
	config, err := ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, divs := range config.Endpoint {
		for _, div := range divs.Keys {
			if resul, err := config.ExistsDivision(div); err != nil {
				t.Errorf("Error: %#v", err)
			} else if !resul {
				t.Errorf("Case:%s\nWant:\n\t%t\nGot:\n\t%t", div, true, false)
			}
		}
	}
}
