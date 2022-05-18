package fizzbuzz

import (
	"fmt"
)

func ReplaceDivisibleByThreeOrFive(input []int) []string {
	out := make([]string, len(input))
	for i, val := range input {
		if val%15 == 0 {
			out[i] = "FizzBuzz"
			continue
		}
		out[i] = fmt.Sprintf("%d", val)
	}
	return out
}
