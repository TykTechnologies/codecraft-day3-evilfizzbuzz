package fizzbuzz

import "strconv"

func Replace(num int) string {
	if num%15 == 0 {
		return "FizzBuzz"
	}
	return strconv.Itoa(num)
}
