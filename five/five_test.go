package five_test

import (
	"testing"

	"github.com/matryer/is"

	five "github.com/TykTechnologies/codecraft-day3-evilfizzbuzz/five"
)

func TestOutputAsString(t *testing.T) {
	is := is.New(t)
	input := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}
	output := five.OutputAsString(input)

	is.Equal(output, "1,2,Fizz,4,Buzz,Fizz,7")
}
