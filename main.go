package main

import (
	"github.com/TykTechnologies/codecraft-day3-evilfizzbuzz/one"
	"github.com/TykTechnologies/codecraft-day3-evilfizzbuzz/three"
	"github.com/TykTechnologies/codecraft-day3-evilfizzbuzz/two"
)

func main() {
	fromOne := one.GenerateInput(1, 100)

	fromTwo := two.ReplaceDivisibleByThree(fromOne)

	br := three.NewBuzzReplacer()
	sendToFour := br.ReplaceListOfStringsWithBuzz(fromTwo)

	_ = sendToFour // TODO
}
