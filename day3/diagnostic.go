package day3

import "math"

func Diagnostic(report []int) int {
	return 0
}

func GetDigit(number int, pos int) int {
	r := number % int(math.Pow(10, float64(pos)))
	return r / int(math.Pow(10, float64(pos-1)))}
