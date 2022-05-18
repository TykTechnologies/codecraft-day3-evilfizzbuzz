package three

import (
	"strconv"
)

type Replacer struct{}

func NewBuzzReplacer() *Replacer {
	return &Replacer{}
}

func (r *Replacer) ReplaceListOfStringsWithBuzz(input []string) []string {
	newList := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		number, err := strconv.Atoi(input[i])
		if err != nil {
			newList[i] = input[i]
		}

		if number%5 == 0 {
			newList[i] = "Buzz"
		} else {
			newList[i] = input[i]
		}
	}

	return newList
}

func (r *Replacer) ReplaceListOfIntegersWithBuzz(input []int) []string {
	newList := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		newList[i] = strconv.Itoa(input[i])
	}

	return r.ReplaceListOfStringsWithBuzz(newList)
}
