package data

import "testing"

func TestConnection(t *testing.T) {
	config, err := ReadFile("config.json")
	if err != nil {
		t.Error(err)
	}
	connection := &Connection{config}
	connection.Get()
}
