package calc

import (
	"errors"
	"fmt"
)

//This is a shared library, so the package is calc

func  AlgebricAdd( float642 float64, float643 float64) (float64) {
	fmt.Println("called calc.AlgebricAdd")
    return float642+float643
}

func AlgebricMinus(float642 float64, float643 float64)(float64, error) {
	fmt.Println("called calc.AlgebricMinus")
	if float642 - float643 < 0 {
		return 0, errors.New("Negative result by err")
	}
	return float642 - float643, nil
}

func SeveralInts (numbers ...int) {
	fmt.Println("Varaidic func", numbers)
}

func Average(nums ...float64) float64{
	var sum float64 =0
	for _, number := range nums {
        sum = sum+number
	}
	return sum/float64 (len(nums))
}
