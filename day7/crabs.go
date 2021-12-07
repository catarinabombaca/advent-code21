package day7

import (
	"fmt"
	"github.com/catarinabombaca/advent-code21/utils"
	"math"
	"strconv"
	"strings"
)

func AlignPosition(crabsPosition []int) int {
	var fuelSpent []int

	for _, pos := range crabsPosition {
		var totalFuel float64
		for _, nextPos := range crabsPosition {
			totalFuel += math.Abs(float64(pos - nextPos))
		}
		fuelSpent = append(fuelSpent, int(totalFuel))
	}

	return findMin(fuelSpent)
}

func AlignPositionExponential(crabsPosition []int) int {
	var fuelSpent []int

	for _, pos := range crabsPosition {
		var totalFuel float64
		for _, nextPos := range crabsPosition {
			delta := math.Abs(float64(pos - nextPos))
			totalFuel += delta*(delta+1)/2
		}
		fuelSpent = append(fuelSpent, int(totalFuel))
	}

	return findMin(fuelSpent)
}

func findMin(values []int) int {
	var minValue int
	for i, value := range values {
		if i == 0 || minValue > value {minValue = value}
	}

	return minValue
}

func GetCrabsPosition() []int {
	b := utils.GetFileContent("data_day7.txt")

	values := strings.Split(string(b), ",")
	data := make([]int, 0, len(values))

	for _, value := range values {
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		position, err := strconv.Atoi(value)
		if err != nil {
			fmt.Errorf("error converting to measumerement to int: %v", err)
		}

		data = append(data, position)
	}

	return data
}