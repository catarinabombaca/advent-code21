package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSonar(t *testing.T) {
	assert := assert.New(t)
	//given 5 depth measurements [199, 200, 208, 210, 200]
	depthMeasurements := []int{199,200,208,210,200}

	//when we call Sonar
	count := Sonar(depthMeasurements)
	//should return 3 as a count
	assert.Equal(3, count)
}

func TestSonarSlidingWindow(t *testing.T) {
	assert := assert.New(t)
	//given 6 depth measurements [199, 200, 208, 210, 200, 207]
	depthMeasurements := []int{199, 200, 208, 210, 200, 207}

	//when we call SonarSlidingWindow
	count := SonarSlidingWindow(depthMeasurements)

	//should return 1 as a count
	assert.Equal(1, count)
}