package day09

import (
	"advent/common"

	"testing"
)

func TestDay2Part1(t *testing.T) {
	input := common.GetInput("day02/testinput1")
	LoadProgram(input)
	result := SolveProblem()

	if result != 2 {
		t.Fatalf("Received incorrect result on day 2 test 1: %v", result)
	}
}

func TestDay2Part4(t *testing.T) {
	input := common.GetInput("day02/testinput4")
	LoadProgram(input)
	result := SolveProblem()

	if result != 30 {
		t.Fatalf("Received incorrect result on day 2 test 4: %v", result)
	}
}

func TestDay5Part1(t *testing.T) {
	input := common.GetInput("day05/testinput1")
	LoadProgram(input)
	SetProgramAlarm()
	result := SolveProblem()

	if result != 2692315 {
		t.Fatalf("Received incorrect result on test 1: %v", result)
	}
}

type StringWriter struct {
	Value string
}

func NewStringWriter() *StringWriter {
	sw := new(StringWriter)
	sw.Value = ""
	return sw
}

func (sw *StringWriter) Write(p []byte) (n int, err error) {
	sw.Value = string(p)

	return len(p), nil
}

func TestDay5Part3(t *testing.T) {
	input := common.GetInput("day05/testinput3")
	LoadProgram(input)

	sw := NewStringWriter()

	IP = 0
	STD_INPUT = 50
	STD_OUTPUT = sw
	STD_OUTPUT_INT64 = make([]int64, 0)

	ExecuteProgram()

	if STD_OUTPUT_INT64[0] != 1001 {
		t.Errorf("Expected 1001, received %v", STD_OUTPUT_INT64[0])
	}

	LoadProgram(input)

	IP = 0
	STD_INPUT = 7
	sw = NewStringWriter()
	STD_OUTPUT = sw
	STD_OUTPUT_INT64 = make([]int64, 0)

	ExecuteProgram()

	if STD_OUTPUT_INT64[0] != 999 {
		t.Errorf("Expected 999, received %v", STD_OUTPUT_INT64[0])
	}

	IP = 0
	STD_INPUT = 8
	sw = NewStringWriter()
	STD_OUTPUT = sw
	STD_OUTPUT_INT64 = make([]int64, 0)

	ExecuteProgram()

	if STD_OUTPUT_INT64[0] != 1000 {
		t.Errorf("Expected 1000, received %v", STD_OUTPUT_INT64[0])
	}
}

func TestCopyProgram(t *testing.T) {
	input := common.GetInput("day09/copyprograminput")

	SetProgram(input[0])
	ResetProgram()

	ExecuteProgram()

	if STD_OUTPUT_INT64[0] != 109 ||
		STD_OUTPUT_INT64[1] != 1 ||
		STD_OUTPUT_INT64[2] != 204 ||
		STD_OUTPUT_INT64[3] != -1 ||
		STD_OUTPUT_INT64[15] != 99 {
		t.Errorf("Expected program itself, received %v", STD_OUTPUT_INT64)
	}
}

func TestOutputMiddle(t *testing.T) {
	input := common.GetInput("day09/middleoutputtest")

	SetProgram(input[0])
	ResetProgram()

	ExecuteProgram()

	if STD_OUTPUT_INT64[0] != 1125899906842624 {
		t.Errorf("Expected 1125899906842624, received %v", STD_OUTPUT_INT64[0])
	}

}

func Test16Dig(t *testing.T) {
	input := common.GetInput("day09/sixteen")

	SetProgram(input[0])
	ResetProgram()

	ExecuteProgram()

	if STD_OUTPUT_INT64[0] != 1219070632396864 {
		t.Errorf("Expected sixteen diget number, received %v", STD_OUTPUT_INT64[0])
	}
}
