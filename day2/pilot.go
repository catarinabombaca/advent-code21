package day2

//Todo: refactor

import (
	"fmt"
	"github.com/catarinabombaca/advent-code21/utils"
	"strings"
)

type Step struct {
	direction string
	value int
}

type Submarine struct {
	Position int
	Depth int
	Aim int
}

func (s *Submarine) Move(step Step) {
	switch step.direction {
	case "forward":
		s.Position += step.value
	case "down":
		s.Depth += step.value
	case "up":
		s.Depth -= step.value
	default:
		fmt.Println("unknown direction")
	}
}

func (s *Submarine) MoveWithAim(step Step) {
	switch step.direction {
	case "forward":
		s.Position += step.value
		s.Depth += s.Aim*step.value
	case "down":
		s.Aim += step.value
	case "up":
		s.Aim -= step.value

	default:
		fmt.Println("unknown direction")
	}
}

func Pilot(instructions []string, hasAim bool) int {
	steps := getSteps(instructions)
	submarine := Submarine{0, 0, 0}

	for _, step := range steps {
		if hasAim {
			submarine.MoveWithAim(step)
		} else {
			submarine.Move(step)
		}
	}

	return submarine.Position * submarine.Depth
}

func getSteps(instructions []string) []Step {
	steps := make([]Step, 0, len(instructions))

	for _, instruction := range instructions {
		step := strings.Split(instruction, " ")
		value := utils.GetDigit(step[1])
		steps = append(steps, Step{step[0], value})
	}

	return steps
}

func GetValuesAsString(filename string) []string {
	b := utils.GetFileContent(filename)
	values := strings.Split(string(b), "\n")

	return values
}
