package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlignPosition(t *testing.T) {
	//given a slice with the crabs positions
	crabsPosition := []int{16,1,2,0,4,2,7,1,2,14}
	//when we can AlignPosition()
	fuel := AlignPosition(crabsPosition)
	//then it should return the least amount of fuel for the align
	assert.Equal(t, 37, fuel)
}
