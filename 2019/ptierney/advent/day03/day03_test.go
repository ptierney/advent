package day03

import (
	"advent/common"

	"testing"
)

func TestExecute1(t *testing.T) {
	input := common.GetInput("day03/input_test1")
	result := SolveProblem(input)

	if result != 6 {
		t.Fatalf("Received incorrect result on test 1: %v", result)
	}
}

func TestExecute2(t *testing.T) {
	input := common.GetInput("day03/input_test2")
	result := SolveProblem(input)

	if result != 159 {
		t.Fatalf("Received incorrect result on test 2: %v", result)
	}
}

func TestExecute3(t *testing.T) {
	input := common.GetInput("day03/input_test3")
	result := SolveProblem(input)

	if result != 135 {
		t.Fatalf("Received incorrect result on test 3: %v", result)
	}
}

func TestExecute1_2(t *testing.T) {
	input := common.GetInput("day03/input_test1")
	result := SolvePart2(input)

	if result != 30 {
		t.Fatalf("Received incorrect result on test 1_2: %v", result)
	}
}

func TestExecute2_2(t *testing.T) {
	input := common.GetInput("day03/input_test2")
	result := SolvePart2(input)

	if result != 610 {
		t.Fatalf("Received incorrect result on test 2_2: %v", result)
	}
}

func TestExecute3_2(t *testing.T) {
	input := common.GetInput("day03/input_test3")
	result := SolvePart2(input)

	if result != 410 {
		t.Fatalf("Received incorrect result on test 3_2: %v", result)
	}
}
