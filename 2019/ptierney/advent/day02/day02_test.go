package day02

import (
	"advent/common"

	"testing"
)

func TestExecute1(t *testing.T) {
	input := common.GetInput("day02/testinput1")
	LoadProgram(input)
	result := SolveProblem()

	if result != 2 {
		t.Fatalf("Received incorrect result on test 1: %v", result)
	}
}

func TestExecute4(t *testing.T) {
	input := common.GetInput("day02/testinput4")
	LoadProgram(input)
	result := SolveProblem()

	if result != 30 {
		t.Fatalf("Received incorrect result on test 1: %v", result)
	}
}
