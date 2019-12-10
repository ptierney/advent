package day05

import (
	"advent/common"

	"bufio"
	"os"
	"testing"
)

func TestExecuteDay2_1(t *testing.T) {
	input := common.GetInput("day02/testinput1")
	LoadProgram(input)
	result := SolveProblem()

	if result != 2 {
		t.Fatalf("Received incorrect result on day 2 test 1: %v", result)
	}
}

func TestExecuteDay2_4(t *testing.T) {
	input := common.GetInput("day02/testinput4")
	LoadProgram(input)
	result := SolveProblem()

	if result != 30 {
		t.Fatalf("Received incorrect result on day 2 test 4: %v", result)
	}
}

func TestExecute1(t *testing.T) {
	input := common.GetInput("day05/testinput1")
	LoadProgram(input)
	SetProgramAlarm()
	result := SolveProblem()

	if result != 2692315 {
		t.Fatalf("Received incorrect result on test 1: %v", result)
	}
}

func TestExecuteFinalPart1(t *testing.T) {
	input := common.GetInput("day05/input")
	LoadProgram(input)

	f, err := os.OpenFile("output_part1", os.O_RDWR|os.O_CREATE, 0660)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := bufio.NewWriter(f)

	IP = 0
	STD_INPUT = 1
	STD_OUTPUT = w

	ExecuteProgram()

	w.Flush()

	f.Close()
}

func TestExecutePart2(t *testing.T) {
	input := common.GetInput("day05/testinput3")
	LoadProgram(input)

	f, err := os.OpenFile("output_testpart2", os.O_RDWR|os.O_CREATE, 0660)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := bufio.NewWriter(f)

	IP = 0
	STD_INPUT = 50
	STD_OUTPUT = w

	ExecuteProgram()

	w.Flush()

	f.Close()
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

func TestExecutePart2_2(t *testing.T) {
	input := common.GetInput("day05/testinput3")
	LoadProgram(input)

	sw := NewStringWriter()

	IP = 0
	STD_INPUT = 50
	STD_OUTPUT = sw

	ExecuteProgram()

	if sw.Value != "1001\n" {
		t.Errorf("Expected 1001, received %v", sw.Value)
	}

	LoadProgram(input)

	IP = 0
	STD_INPUT = 7
	sw = NewStringWriter()
	STD_OUTPUT = sw

	ExecuteProgram()

	if sw.Value != "999\n" {
		t.Errorf("Expected 999, received %v", sw.Value)
	}

}

func TestExecuteFinalPart2(t *testing.T) {
	input := common.GetInput("day05/input")
	LoadProgram(input)

	f, err := os.OpenFile("output_part2", os.O_RDWR|os.O_CREATE, 0660)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := bufio.NewWriter(f)

	IP = 0
	STD_INPUT = 5
	STD_OUTPUT = w

	ExecuteProgram()

	w.Flush()

	f.Close()
}
