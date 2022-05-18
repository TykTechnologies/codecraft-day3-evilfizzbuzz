package codecraft_day3_evilfizzbuzz

func generateInput() []int {
	ints := make([]int, 100)

	for i := range ints {
		ints[i] = i + 1
	}

	return ints
}
