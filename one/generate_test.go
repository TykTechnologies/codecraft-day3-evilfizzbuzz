package one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateInput(t *testing.T) {
	t.Run("Generates a list of integers from 1 to 100", func(t *testing.T) {
		fizzbuzzInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

		assert.Equal(t, fizzbuzzInput, GenerateInput(1, 100))
	})

	t.Run("Generates a list of integers from 10 to 20", func(t *testing.T) {
		tenToTwenty := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

		assert.Equal(t, tenToTwenty, GenerateInput(10, 20))
	})

	t.Run("Generates a list of a 0 integer", func(t *testing.T) {
		assert.Equal(t, []int{0}, GenerateInput(0, 0))
	})

	t.Run("Doesn't generate anything for a range of 20,10", func(t *testing.T) {
		assert.Equal(t, []int{}, GenerateInput(20, 10))
	})

	t.Run("Doesn't generate anything for negative values", func(t *testing.T) {
		assert.Equal(t, []int{}, GenerateInput(-1, -1))
	})
}
