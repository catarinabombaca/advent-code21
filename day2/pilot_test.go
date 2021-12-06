package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPilot(t *testing.T) {
	assert := assert.New(t)

	t.Run("pilot", func(t *testing.T) {
		//given a slice with steps [forward 5 down 5 forward 8 up 3 down 8 forward 2]
		instructions := []string{"forward 5", "down 5","forward 8", "up 3", "down 8", "forward 2"}
		//when we call Pilot()
		product := Pilot(instructions)
		//then the product of the coordinates should be 150
		assert.Equal(150, product)
	})
}
