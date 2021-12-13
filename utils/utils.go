package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func GetFileContent(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
	}

	return b
}

func GetDigit(number string) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		fmt.Errorf("error converting to value to int: %v", err)
	}
	return value
}
