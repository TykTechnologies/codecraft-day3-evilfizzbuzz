package one

func GenerateInput(min int, max int) []int {
	if min > max {
		return []int{}
	}

	ints := make([]int, max-min+1)

	for i := range ints {
		ints[i] = min + i
	}

	return ints
}
