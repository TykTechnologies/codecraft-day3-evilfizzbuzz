package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceDivisibleByThreeOrFiveInSequence(t *testing.T) {
	input := []int{1, 3, 6, 15, 26, 30}
	expected := []string{"1", "3", "6", "FizzBuzz", "26", "FizzBuzz"}
	assert.Equal(t, expected, ReplaceDivisibleByThreeOrFive(input))
}
