package main

import (
	"fmt"
	"github.com/catarinabombaca/advent-code21/day1"
	"github.com/catarinabombaca/advent-code21/day2"
	"github.com/catarinabombaca/advent-code21/day3"
	"github.com/catarinabombaca/advent-code21/day7"
)

func main() {
	//Day 1 results
	measurements := day1.GetInts("data.txt")
	count1 := day1.Sonar(measurements)
	count2 := day1.SonarSlidingWindow(measurements)

	fmt.Printf("Day 1: %v/%v\n", count1, count2)

	//Day 2 results
	instructions := day2.GetValuesAsString("data_day2.txt")
	product1 := day2.Pilot(instructions, false)
	product2 := day2.Pilot(instructions, true)

	fmt.Printf("Day 2: %v/%v\n", product1, product2)

	//Day 3 results
	binaries := day2.GetValuesAsString("data_day3.txt")
	power1 := day3.Diagnostic(binaries)

	fmt.Printf("Day 3: %v/...\n", power1)

	crabsPos := day7.GetCrabsPosition()
	minFuel1 := day7.AlignPosition(crabsPos)
	minFuel2 := day7.AlignPositionExponential(crabsPos)

	fmt.Printf("Day 7: %v/%v\n", minFuel1, minFuel2)
}


