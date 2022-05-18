package codecraft_day3_evilfizzbuzz

func generateInput(min int, max int) []int {
	ints := make([]int, max-min+1)

	for i := range ints {
		ints[i] = min + i
	}

	return ints
}
