package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWithBuzz(t *testing.T) {
	t.Run("replace with buzz when divisible by 5", func(t *testing.T) {
		replacer := NewBuzzReplacer()
		expectedResult := []string{"1", "2", "3", "4", "Buzz"}

		t.Run("with string slice input", func(t *testing.T) {
			input := []string{"1", "2", "3", "4", "5"}
			actualResult := replacer.ReplaceListOfStringsWithBuzz(input)
			assert.Equal(t, expectedResult, actualResult)
		})

		t.Run("with int slice input", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			actualResult := replacer.ReplaceListOfIntegersWithBuzz(input)
			assert.Equal(t, expectedResult, actualResult)
		})
	})
}
