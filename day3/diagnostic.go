package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Diagnostic(report []string) int {
	var gammaNumbers []string
	var epsilonNumbers []string
	for i := 0; i < len(report[0]); i++ {
		var numbers []string
		for _, binary := range report {
			numbers = append(numbers, string(binary[i]))
		}
		mostCommon := GetMostCommonNumber(numbers)
		leastCommon := GetLeastCommonNumber(numbers)
		gammaNumbers = append(gammaNumbers, mostCommon)
		epsilonNumbers = append(epsilonNumbers, leastCommon)
	}

	gammaDecimal, _ := convertBinaryToDecimal(strings.Join(gammaNumbers, ""))
	epsilonDecimal, _ := convertBinaryToDecimal(strings.Join(epsilonNumbers, ""))
	return gammaDecimal * epsilonDecimal
}

func GetMostCommonNumber(numbers []string) string {
		countMap := make(map[string]int)
		for _, value := range numbers {
			countMap[value] += 1
		}

		if countMap["0"] > countMap["1"] {
			return "0"
		} else {
			return "1"
		}
}

func GetLeastCommonNumber(numbers []string) string {
	countMap := make(map[string]int)
	for _, value := range numbers {
		countMap[value] += 1
	}

	if countMap["0"] > countMap["1"] {
		return "1"
	} else {
		return "0"
	}
}

func convertBinaryToDecimal(binary string) (int, error) {
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(output), nil
}



