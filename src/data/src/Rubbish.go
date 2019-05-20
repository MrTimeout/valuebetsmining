package data

import (
	"errors"
	"math"
)

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

func Mode(nums ...int) {

}
