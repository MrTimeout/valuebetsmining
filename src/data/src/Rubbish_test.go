package data

import (
	"testing"
)

//TestEmptyStringReadFile ... test that if we pass an empty array, it will trigger an error.
func TestEmptyArrAverage(t *testing.T) {
	emptyArray := []int{}
	if _, err := Average(emptyArray, true); err == nil {
		t.Errorf("Expected-> %s\n\tGet-> %s\n\t", "Error parsing array of integers", err)
	}
}

//TestEmptyStringReadFile ... test that if we pass an array, it will return the correct result.
func TestNegativeArrAverage(t *testing.T) {
	var array = []int{}
	for index := -1; index >= -10; index-- {
		array = append(array, index)
	}
	if avg, _ := Average(array, true); avg != -5.5 {
		t.Errorf("Expected-> %s\n\tGet-> %f\n\t", "-5.5", avg)
	}
	if avg, _ := Average(array, false); avg != 5.5 {
		t.Errorf("Expected-> %s\n\tGet-> %f\n\t", "5.5", avg)
	}
}

//TestEmptyStringReadFile ... test that if we pass an array, it will return the correct result.
func TestPositiveArrAverage(t *testing.T) {
	var array = []int{}
	for index := 1; index <= 10; index++ {
		array = append(array, index)
	}
	if avg, _ := Average(array, true); avg != 5.5 {
		t.Errorf("Expected-> %s\n\tGet-> %f\n\t", "5.5", avg)
	}
	if avg, _ := Average(array, false); avg != 5.5 {
		t.Errorf("Expected-> %s\n\tGet-> %f\n\t", "5.5", avg)
	}
}

//TestEmptyStringReadFile ... test that if we pass an array, it will return the correct result.
func TestMixtArrAverage(t *testing.T) {
	var array = []int{}
	for index := -1; index >= -10; index-- {
		array = append(array, []int{index, -index}...)
	}
	if avg, _ := Average(array, true); avg != 0 {
		t.Errorf("Expected-> %s\n\tGet-> %f\n\t", "0", avg)
	}
	if avg, _ := Average(array, false); avg != 5.5 {
		t.Errorf("Expected-> %s\n\tGet-> %f\n\t", "5.5", avg)
	}
}
