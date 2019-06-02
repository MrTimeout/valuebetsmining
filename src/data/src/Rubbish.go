package data

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"strings"
)

//Pair ... Custom pair of values: key=value
type Pair struct {
	Key   int
	Value int
}

//PairList ... Custom list of Pairs
type PairList []Pair

var (
	//ErrMinimalLength ... Error parsing length because is less than or equal to 0
	ErrMinimalLength = errors.New("Error parsing length because is negative or 0")
	//ErrNilArr ... Error parsing arr because it is nil or has length og 0
	ErrNilArr = errors.New("Error parsing arr because it is nil or has len of 0")
)

//Len ... Returns length og the PairList
func (p PairList) Len() int { return len(p) }

//Less ... Returns if second value is greater than first or not
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

//Swap ... Swap two elements inside PairList
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//Average ... Calculate average of an array of integers.\n arr[] int is the sequence of numbers.\n negative is true, when you want to count negative numbers like negative numbers, and false otherwise.
func Average(arr []int, negative bool) (float64, error) {
	if len(arr) == 0 {
		return 0.0, errors.New("Error parsing array of integers")
	}
	sum := 0.0
	if negative {
		for _, value := range arr {
			sum += float64(value)
		}
	} else {
		for _, value := range arr {
			sum += math.Abs(float64(value))
		}
	}
	return sum / float64(len(arr)), nil
}

//Mode ... Calculate the mode of a range of numbers
func Mode(nums ...int) ([]int, error) {
	if len(nums) == 0 {
		return []int{}, errors.New("Error parsing array of integers")
	}
	m := make(map[int]int)

	for _, value := range nums {
		m[value]++
	}

	result, err := MaxsOfAMap(m)
	if err != nil {
		return []int{}, err
	}

	return result, nil
}

//MaxsOfAMap ... Order a map by its values
func MaxsOfAMap(m map[int]int) ([]int, error) {
	if len(m) == 0 {
		return []int{}, errors.New("Error parsing map of integers")
	}
	resul := make([]int, 0)
	_, values := ParseMapIntToKeysValues(m)
	max, err := MaxArr(values)
	if err != nil {
		return []int{}, errors.New("Error parsing values of arr values")
	}
	for key, value := range m {
		if value == max {
			resul = append(resul, key)
		}
	}
	return resul, nil
}

//ParseMapIntToKeysValues ... Return an array of the keys and values
func ParseMapIntToKeysValues(m map[int]int) ([]int, []int) {
	keys, values := make([]int, 0, len(m)), make([]int, 0, len(m))
	for k, v := range m {
		keys, values = append(keys, k), append(values, v)

	}
	return keys, values
}

//MaxArr ... Max number of an array of integer
func MaxArr(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Error parsing arr")
	}
	i, max := 0, math.MinInt64
	for {
		if i == len(arr) {
			break
		}
		if arr[i] >= max {
			max = arr[i]
		}
		i++
	}
	return max, nil
}

//WhoIsBigger ... returns -1|0|1 if a < b|a == b|a > b
func WhoIsBigger(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

//HowManyTimes ... Returns how many times a value or an array of values are inside an amount of data
func HowManyTimes(data []int, values ...int) (int, error) {
	if len(data) == 0 || len(values) == 0 {
		return -1, errors.New("Error parsing array of integers")
	}
	hmt := 0
	for index := 0; index < len(data); index++ {
		b, err := AmIHere(values, data[index])
		if err != nil {
			return -1, err
		}
		if !b {
			break
		} else {
			hmt++
		}
	}
	return hmt, nil
}

//AmIHere ... Returns bool is a element exists in an array
func AmIHere(data []int, value int) (bool, error) {
	if len(data) == 0 {
		return false, errors.New("Error parsing array of integers")
	}
	for _, val := range data {
		if value == val {
			return true, nil
		}
	}
	return false, nil
}

//RandomArr ... Random arrays of integers in [from, to]{amount}
func RandomArr(from, to, amount int) ([]int, error) {
	if from < 0 || to < 0 || amount <= 0 || from >= to {
		return nil, errors.New("Error parsing params")
	}
	rands := make([]int, 0, amount)
	for i := 0; i < amount; i++ {
		rands = append(rands, rand.Intn(to-from)+from)
	}
	return rands, nil
}

//RandomArrNegative ... Random arrays of integers in [from, to]{amount}
func RandomArrNegative(from, to, amount int) ([]int, error) {
	if amount <= 0 || from >= to {
		return nil, errors.New("Error parsing params")
	}
	rands := make([]int, 0, amount)
	for i := 0; i < amount; i++ {
		rands = append(rands, (rand.Intn(to-from)+from)*RandomSign())
	}
	return rands, nil
}

//RandomSign ... returns -1 or 1, dependends on if it is odd or even
func RandomSign() int {
	if r := rand.Int(); r%2 == 0 {
		return 1
	}
	return -1
}

//RandomWords ... Return a random arr of words
func RandomWords(len, lenWord int) ([]string, error) {
	if len <= 0 {
		return []string{}, errors.New("Error parsing len of the array")
	}
	if lenWord <= 0 {
		lenWord = rand.Intn(rand.Int()) + rand.Int()
	}
	str, i := make([]string, 0, len), 0
	for {
		if i == len {
			break
		}
		s, err := RandomWord(lenWord)
		if err != nil {
			return []string{}, err
		}
		str = append(str, s)
		i++
	}
	return str, nil
}

//RandomWord ... Returns a random word of an specific len
func RandomWord(len int) (string, error) {
	if len <= 0 {
		return "", errors.New("Error parsing length of the word")
	}
	s, i := "", 0
	for {
		if i == len {
			break
		}
		s += RandomLetter()
		i++
	}
	return s, nil
}

//RandomLetter ... It returns a latin letter
func RandomLetter() string {
	return string(rand.Intn(122-97) + 97)
}

//IsASCIICode ... Returns true if it is between and b
func IsASCIICode(min, max int, char rune) (bool, error) {
	if min <= 0 || max <= 0 {
		return false, errors.New("Error parsing params")
	}
	return int(char) >= min && int(char) <= max, nil
}

//IsASCIICodeArr ... Return true if it is between min and max
func IsASCIICodeArr(runes []rune, min, max int) (bool, error) {
	if min <= 0 || max <= 0 {
		return false, errors.New("Error parsing params")
	}
	for _, v := range runes {
		if int(v) >= min && int(v) <= max {
			return false, nil
		}
	}
	return true, nil
}

//CompareTwoArrs ... Compare two arrays returning bool value. True if they are equals and false otherwise
func CompareTwoArrs(arr, arr2 []int, sor bool) bool {
	if len(arr) != len(arr2) {
		return false
	}
	if sor {
		sort.Ints(arr)
		sort.Ints(arr2)
	}
	index := 0
	for {
		if index == len(arr) {
			return true
		}
		if arr[index] != arr2[index] {
			return false
		}
		index++
	}
}

//RandomArrWithValues ... Creates a random arr with values passed by param
func RandomArrWithValues(length int, values []int) ([]int, error) {
	if length <= 0 {
		return []int{}, ErrMinimalLength
	}
	if values == nil || len(values) == 0 {
		return []int{}, ErrNilArr
	}
	var arr = []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(length-0)+0)
	}
	return arr, nil
}

//IsAStrangerHere ... Test if there are not other values than the ones inside the helper
func IsAStrangerHere(target, helper []int) (bool, error) {
	if target == nil || helper == nil || len(target) == 0 || len(helper) == 0 {
		return false, ErrNilArr
	}
	for i := 0; i < len(target); i++ {
		if b, err := AmIHere(helper, target[i]); err != nil {
			return false, err
		} else if !b {
			return true, nil
		}
	}
	return false, nil
}

//AnyoneIsEmpty ... Returns a bool value! false in case of != "" an true otherwise
func AnyoneIsEmpty(arr []string) (bool, error) {
	if arr == nil || len(arr) == 0 {
		return false, ErrNilArr
	}
	for i := 0; i < len(arr); i++ {
		if strings.Trim(arr[i], " ") == "" || len(strings.Trim(arr[i], " ")) == 0 {
			return true, nil
		}
	}
	return false, nil
}

/*
//DisorderAMap ... Disorder a map of integers
func DisorderAMap(m map[int]int) (map[int]int, error) {
	if len(m) == 0 {
		return m, errors.New("Error parsing map of integers")
	}
	temp := []int{}
	for _
	for index := 0; index <= len(m)/2; index++ {
		m[index:]
	}
	return m, nil
}
*/
