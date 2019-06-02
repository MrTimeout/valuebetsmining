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
			if err := config.ExistsDivision(div); err != nil {
				t.Errorf("Error: %#v", err)
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
		if err := config.ExistsCountry(countries.Name); err != nil {
			t.Errorf("Error:%#v", err)
		}
	}
}
