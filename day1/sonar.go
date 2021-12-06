package day1

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
