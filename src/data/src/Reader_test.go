package data

import (
	"testing"
)

//TestEmptyStringReadFile ... test that if we pass a string empty or fill with white backspaces, it will trigger a boolean false.
func TestEmptyStringReadFile(t *testing.T) {
	emptyString := ""
	fillWhiteBackspace := "   "
	normalString := "config.json"
	if _, err := ReadFile(emptyString); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file name", err)
	}
	if _, err := ReadFile(fillWhiteBackspace); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file name", err)
	}
	if _, err := ReadFile(normalString); err != nil {
		t.Errorf("Expected-> %#v\n\tGet-> %s\n\t", nil, err)
	}
}

//TestNoExistingFileReadFile ... We are trying to test two files: the first one doesnt exists and the second one is correct
func TestNoExistingFileReadFile(t *testing.T) {
	noExistingFile := "abc123"
	existingFile := "config.json"
	if _, err := ReadFile(noExistingFile); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file", err)
	}
	if _, err := ReadFile(existingFile); err != nil {
		t.Errorf("Expected-> %#v\n\tGet-> %s\n\t", nil, err)
	}
}

//TestErrorInJSONReadFile ... We are trying to test two JSON: the first one contains errors and the second one is correct
func TestErrorInJSONReadFile(t *testing.T) {
	failedJSON := "config_test_1.json"
	successJSON := "config.json"
	if _, err := ReadFile(failedJSON); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing file because of failed JSON", err)
	}
	if _, err := ReadFile(successJSON); err != nil {
		t.Errorf("Expected-> %#v\n\tGet-> %s\n\t", nil, err)
	}
}
