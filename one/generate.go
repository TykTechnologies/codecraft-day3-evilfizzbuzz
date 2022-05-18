package one

func GenerateInput(min int, max int) []int {
	if min > max {
		return []int{}
	}
	
	if min < 0 || max < 0 {
		return []int{}
	}

	ints := make([]int, max-min+1)

	for i := range ints {
		ints[i] = min + i
	}

	return ints
}
