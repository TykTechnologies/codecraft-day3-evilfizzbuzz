package replacethree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceDivisibleByThreeInSequence(t *testing.T) {
	input := []int{1, 3, 6, 9, 19}
	expected := []string{"1", "Fizz", "Fizz", "Fizz", "19"}
	assert.Equal(t, expected, ReplaceDivisibleByThree(input))
}

func TestDoesntReplaceDivisibleByFiveInSequence(t *testing.T) {
	input := []int{1, 3, 6, 15, 20, 18, 19}
	expected := []string{"1", "Fizz", "Fizz", "15", "20", "Fizz", "19"}
	assert.Equal(t, expected, ReplaceDivisibleByThree(input))
}
