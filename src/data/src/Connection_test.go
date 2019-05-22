package data

import (
	"regexp"
	"testing"
)

//TestObjectConfigGet ... Testing the object created in config
func TestObjectConfigGet(t *testing.T) {

}

//TestConnectionGetResponse ... Testing connection via get to the Path and getting response
func TestConnectionGetResponse(t *testing.T) {
	config, err := ReadFile("config.json")
	if err != nil {
		t.Error(err)
	}
	connection := &Connection{config}
	res, err := connection.Get()
	if err != nil {
		t.Error(err)
	}
	if match, _ := regexp.MatchString("<html>", res); match {
		t.Errorf("\nWant:\n\t%t\nGot\n\t%t\nResponse:\n%s", false, true, res)
	}
}
