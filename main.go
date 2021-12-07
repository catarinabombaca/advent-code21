package main

import (
	"fmt"
	"github.com/catarinabombaca/advent-code21/day1"
	"github.com/catarinabombaca/advent-code21/day2"
	"github.com/catarinabombaca/advent-code21/day7"
)

func main() {
	//Day 1 results
	measurements := day1.GetMeasurements()
	count1 := day1.Sonar(measurements)
	count2 := day1.SonarSlidingWindow(measurements)

	fmt.Printf("Day 1: %v/%v\n", count1, count2)

	//Day 2 results
	instructions := day2.GetInstructions()
	product1 := day2.Pilot(instructions)

	fmt.Printf("Day 2: %v/...\n", product1)

	crabsPos := day7.GetCrabsPosition()
	minFuel1 := day7.AlignPosition(crabsPos)
	minFuel2 := day7.AlignPositionExponential(crabsPos)

	fmt.Printf("Day 7: %v/%v\n", minFuel1, minFuel2)
}
