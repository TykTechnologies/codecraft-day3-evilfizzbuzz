package two

import "fmt"

func ReplaceDivisibleByThree(input []int) []string {
	out := make([]string, len(input))
	for i, val := range input {
		if val%3 == 0 && val%5 != 0 {
			out[i] = "Fizz"
			continue
		}
		out[i] = fmt.Sprintf("%d", val)
	}
	return out
}
