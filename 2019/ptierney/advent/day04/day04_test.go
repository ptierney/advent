package day04

import (
	"testing"
)

func TestValueCorrect1(t *testing.T) {
	if valueCorrect(111111) == false {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrect2(t *testing.T) {
	if valueCorrect(223450) == true {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrect3(t *testing.T) {
	if valueCorrect(123789) == true {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrect4(t *testing.T) {
	if valueCorrect(334567) == false {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrectPart2_1(t *testing.T) {
	if valueCorrectPart2(112233) == false {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrectPart2_2(t *testing.T) {
	if valueCorrectPart2(123444) == true {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrectPart2_3(t *testing.T) {
	if valueCorrectPart2(111122) == false {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrectPart2_4(t *testing.T) {
	if valueCorrectPart2(111134) == true {
		t.Fatalf("Got wrong answer")
	}
}

func TestValueCorrectPart2_5(t *testing.T) {
	if valueCorrectPart2(111111) == true {
		t.Fatalf("Got wrong answer")
	}
}
