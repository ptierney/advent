package day00

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	result := SumNumbers()

	if result != 433 {
		t.Errorf("Unexpected result")
	}
}
