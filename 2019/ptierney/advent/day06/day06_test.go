package day06

import (
	"advent/common"

	"testing"
)

func TestDirectAndIndirectInputs(t *testing.T) {
	input := common.GetInput("day06/testinput1")

	orbits := GetDirectAndIndirectInputs(input)

	if orbits != 42 {
		t.Errorf("Expected 42 orbits, received %v", orbits)
	}
}
