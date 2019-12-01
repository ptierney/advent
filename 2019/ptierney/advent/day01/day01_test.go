package day01

import (
	"testing"
)

func TestGetFuelAndSubFuel(t *testing.T) {
	if getFuelAndSubFuel(100756) != 50346 {
		t.Errorf("Incorrect sub fuel calculation: 50346")
	}

	if getFuelAndSubFuel(1969) != 966 {
		t.Errorf("Incorrect sub fuel calculation: 1969")
	}
}
