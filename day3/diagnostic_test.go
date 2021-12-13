package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiagnostic(t *testing.T) {
	//given a list of binary numbers as strings [00100 11110 10110 10111 10101 01111 00111 11100 10000 11001 00010 01010]
	diagnosticReport := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	//when we call Diagnostic()
	power := Diagnostic(diagnosticReport)
	//return a power consumption of 198
	assert.Equal(t, 198, power)
}

func TestGetMostCommonNumber(t *testing.T) {
	//given a list of ints [0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0]
	numbers := []string{"0", "1", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"}

	//when we call GetMostCommonNumber()
	common := GetMostCommonNumber(numbers)

	//then should return 1
	assert.Equal(t, "1", common)
}


