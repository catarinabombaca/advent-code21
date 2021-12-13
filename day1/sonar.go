package day1

//Todo: refactor

import (
	"github.com/catarinabombaca/advent-code21/utils"
	"strings"
)

func Sonar(measurements []int) int {
	var count int

	for i, measurement := range measurements {
		if i == 0 {
			continue
		}

		delta := measurement - measurements[i-1]
		if delta > 0 {
			count++
		}
	}

	return count
}

func SonarSlidingWindow(measurements []int) int {
	var sums []int

	for i, measurement := range measurements {
		if i+2 < len(measurements) {
			sum := measurement + measurements[i+1] + measurements[i+2]
			sums = append(sums, sum)
		}
	}

	return Sonar(sums)
}


func GetInts(fileName string) []int {

	b := utils.GetFileContent(fileName)

	lines := strings.Split(string(b), "\n")
	data := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 { continue }
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		measurement := utils.GetDigit(l)

		data = append(data, measurement)
	}

	return data
}
