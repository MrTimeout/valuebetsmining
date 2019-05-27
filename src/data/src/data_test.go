package data

import (
	"testing"
)

func TestExistsDivision(t *testing.T) {
	config, err := ReadFile("config.json")
	if err != nil {
		t.Errorf("Error:%#v", err)
	}
	for _, divs := range config.Endpoint {
		for _, div := range divs.Keys {
			if resul, err := config.ExistsDivision(div); err != nil {
				t.Errorf("Error: %#v", err)
			} else if !resul {
				t.Errorf("\nCase:%s\nWant:\n\t%t\nGot:\n\t%t", div, true, false)
			}
		}
	}
}

func TestExistsCountry(t *testing.T) {
	config, err := ReadFile("config.json")
	if err != nil {
		t.Errorf("Error:%#v", err)
	}
	for _, countries := range config.Endpoint {
		if resul, err := config.ExistsCountry(countries.Name); err != nil {
			t.Errorf("Error:%#v", err)
		} else if !resul {
			t.Errorf("\nCase:%s\nWant:\n\t%t\nGot:\n\t%t", countries.Name, true, false)
		}
	}
}
