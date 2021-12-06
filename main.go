package main

import (
	"fmt"
	"github.com/catarinabombaca/advent-code21/day1"
)

func main() {
	//Day 1 results
	measurements := day1.GetMeasurements()
	count1 := day1.Sonar(measurements)
	count2 := day1.SonarSlidingWindow(measurements)

	fmt.Printf("Day 1: %v/%v\n", count1, count2)

}
