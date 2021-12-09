package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiagnostic(t *testing.T) {
	//given a list of binary numbers [00100 11110 10110 10111 10101 01111 00111 11100 10000 11001 00010 01010]
	diagnosticReport := []int{00100, 11110, 10110, 10111, 10101, 01111, 00111, 11100, 10000, 11001, 00010, 01010}
	//when we call Diagnostic()
	power := Diagnostic(diagnosticReport)
	//return a power consumption of 198
	assert.Equal(t, 198, power)
}

func TestGetDigit(t *testing.T) {
	//00100 an int and a position 2
	number := 00100
	pos := 4
	//when we call
	digit := GetDigit(number, pos)
	//return 0
	assert.Equal(t, 0, digit)
}
