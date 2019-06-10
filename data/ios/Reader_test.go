package ios

import (
	"log"
	"testing"
	"valuebetsmining/data/entities"
)

//TestEmptyStringReadFile ... test that if we pass a string empty or fill with white backspaces, it will trigger a boolean false.
func TestEmptyStringReadFile(t *testing.T) {
	emptyString := ""
	fillWhiteBackspace := "   "
	normalString := entities.DefaultFileConfig
	if _, err := entities.ReadFile(emptyString); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file name", err)
	}
	if _, err := entities.ReadFile(fillWhiteBackspace); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file name", err)
	}
	if _, err := entities.ReadFile(normalString); err != nil {
		t.Errorf("Expected-> %#v\n\tGet-> %s\n\t", nil, err)
	}
}

//TestNoExistingFileReadFile ... We are trying to test two files: the first one doesnt exists and the second one is correct
func TestNoExistingFileReadFile(t *testing.T) {
	noExistingFile := "abc123"
	existingFile := entities.DefaultFileConfig
	if _, err := entities.ReadFile(noExistingFile); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file", err)
	}
	if _, err := entities.ReadFile(existingFile); err != nil {
		t.Errorf("Expected-> %#v\n\tGet-> %s\n\t", nil, err)
	}
}

//TestValueOfJSON ... Execute a simple Readfile testing if it trigger an error or not.
func TestValueOfJSON(t *testing.T) {
	_, err := entities.ReadFile(entities.DefaultFileConfig)
	if err != nil {
		log.Fatal(err)
	}
}
