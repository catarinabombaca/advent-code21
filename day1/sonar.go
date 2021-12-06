package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
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


func GetMeasurements() []int {

	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
	}

	lines := strings.Split(string(b), "\n")
	data := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 { continue }
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		measurement, err := strconv.Atoi(l)
		if err != nil {
			fmt.Errorf("error converting to measumerement to int: %v", err)
		}

		data = append(data, measurement)
	}

	return data
}
