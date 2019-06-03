package data

import (
	"math/rand"
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

//TestMode ... Testing mode method
func TestMode(t *testing.T) {
	i, cases, mode := 0, 20, 4
	arr := make([]int, 0, cases)
	for {
		if i == cases {
			break
		}
		arr = append(arr, mode)
		i++
	}
	if v, err := Mode(arr...); err != nil {
		t.Error(err)
	} else if len(v) != 1 {
		t.Errorf("Length:\nWant:\n\r%d\nGot:\n\r%d", 1, len(v))
	} else if v[0] != mode {
		t.Errorf("Mode:\nWant:\n\r%d\nGot:\n\r%d", mode, v[0])
	}
}

//TestAmIHere ... Testing if the functions works well passing corret params
func TestAmIHere(t *testing.T) {
	ints := []int{
		1,
		3,
		4,
		5,
	}
	testInts := []int{
		1,
		6,
		3,
		10,
	}
	wants := []bool{
		true,
		false,
		true,
		false,
	}
	index := 0
	for {
		if index >= len(ints) {
			break
		}
		if resul, err := AmIHere(ints, testInts[index]); err != nil {
			t.Errorf("Error:%#v", err)
		} else if resul != wants[index] {
			t.Errorf("\nCase:%d\nWant:\n\t%t\nGot:\n\t%t", testInts[index], wants[index], resul)
		}
		index++
	}
}

func TestHowManyTimes(t *testing.T) {
	ints, want := []int{1, 1, 1, 1, 0, -1, 1, 1, 0}, []int{0, 1}
	expected := 3
	got, err := HowManyTimes(ints, true, want...)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("\nCase:%d\nWant:\n\t%d\nGot:\n\t%d", 1, expected, got)
	}
}

//TestRadomLetter ... Random letter testing that are in a range of [97,122]
func TestRadomLetter(t *testing.T) {
	i, cases := 0, 100
	for {
		if i == cases {
			break
		}
		r := []rune(RandomLetter())
		if ri := int(r[0]); ri < 97 || ri > 122 {
			t.Error("Error creating random letter")
		}
		i++
	}
}

//TestRandomWords ... Random words of an specific length
func TestRandomWords(t *testing.T) {
	i, cases := 0, 100
	for {
		if i == cases {
			break
		}
		r, err := RandomWord(rand.Intn(cases-i+1) + i)
		if err != nil {
			t.Error(err)
		}
		if _, err := IsASCIICodeArr([]rune(r), 97, 122); err != nil {
			t.Error(err)
		}
		i++
	}
}
