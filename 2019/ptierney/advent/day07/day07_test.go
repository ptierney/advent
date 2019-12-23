package day07

import (
	"advent/common"

	"testing"
)

func TestDay2Part1(t *testing.T) {
	input := common.GetInput("day02/testinput1")

	c := NewComputer()

	c.LoadProgram(input)
	result := c.SolveProblem()

	if result != 2 {
		t.Fatalf("Received incorrect result on day 2 test 1: %v", result)
	}
}

func TestDay2Part4(t *testing.T) {
	input := common.GetInput("day02/testinput4")

	c := NewComputer()

	c.LoadProgram(input)
	result := c.SolveProblem()

	if result != 30 {
		t.Fatalf("Received incorrect result on day 2 test 4: %v", result)
	}
}

func TestDay5Part1(t *testing.T) {
	input := common.GetInput("day05/testinput1")

	c := NewComputer()

	c.LoadProgram(input)
	c.SetProgramAlarm()
	result := c.SolveProblem()

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

	c := NewComputer()

	c.LoadProgram(input)

	sw := NewStringWriter()

	c.IP = 0
	c.SetInput(50)
	c.STD_OUTPUT = sw
	c.STD_OUTPUT_INT64 = make([]int64, 0)

	c.ExecuteProgram()

	if c.STD_OUTPUT_INT64[0] != 1001 {
		t.Errorf("Expected 1001, received %v", c.STD_OUTPUT_INT64[0])
	}

	c.LoadProgram(input)

	c.IP = 0
	c.SetInput(7)
	sw = NewStringWriter()
	c.STD_OUTPUT = sw
	c.STD_OUTPUT_INT64 = make([]int64, 0)

	c.ExecuteProgram()

	if c.STD_OUTPUT_INT64[0] != 999 {
		t.Errorf("Expected 999, received %v", c.STD_OUTPUT_INT64[0])
	}

	c.IP = 0
	c.SetInput(8)
	sw = NewStringWriter()
	c.STD_OUTPUT = sw
	c.STD_OUTPUT_INT64 = make([]int64, 0)

	c.ExecuteProgram()

	if c.STD_OUTPUT_INT64[0] != 1000 {
		t.Errorf("Expected 1000, received %v", c.STD_OUTPUT_INT64[0])
	}
}

func TestCopyProgram(t *testing.T) {
	input := common.GetInput("day09/copyprograminput")

	c := NewComputer()

	c.SetProgram(input[0])
	c.ResetProgram()

	c.ExecuteProgram()

	if c.STD_OUTPUT_INT64[0] != 109 ||
		c.STD_OUTPUT_INT64[1] != 1 ||
		c.STD_OUTPUT_INT64[2] != 204 ||
		c.STD_OUTPUT_INT64[3] != -1 ||
		c.STD_OUTPUT_INT64[15] != 99 {
		t.Errorf("Expected program itself, received %v", c.STD_OUTPUT_INT64)
	}
}

func TestOutputMiddle(t *testing.T) {
	input := common.GetInput("day09/middleoutputtest")

	c := NewComputer()

	c.SetProgram(input[0])
	c.ResetProgram()

	c.ExecuteProgram()

	if c.STD_OUTPUT_INT64[0] != 1125899906842624 {
		t.Errorf("Expected 1125899906842624, received %v", c.STD_OUTPUT_INT64[0])
	}
}

func Test16Dig(t *testing.T) {
	input := common.GetInput("day09/sixteen")

	c := NewComputer()

	c.SetProgram(input[0])
	c.ResetProgram()

	c.ExecuteProgram()

	if c.STD_OUTPUT_INT64[0] != 1219070632396864 {
		t.Errorf("Expected sixteen diget number, received %v", c.STD_OUTPUT_INT64[0])
	}
}

func TestAmplifier1(t *testing.T) {
	input := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	ms, ps := FindMaxThrusterSignal(input)

	if ms != 43210 {
		t.Errorf("Expected Max Signal 43210, received %v, with settings %v", ms, ps)
	}
}

func TestAmplifier2(t *testing.T) {
	input := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	ms, ps := FindMaxThrusterSignal(input)

	if ms != 54321 {
		t.Errorf("Expected Max Signal 54321, received %v, with settings %v", ms, ps)
	}
}

func TestAmplifier3(t *testing.T) {
	input := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	ms, ps := FindMaxThrusterSignal(input)

	if ms != 65210 {
		t.Errorf("Expected Max Signal 65210, received %v, with settings %v", ms, ps)
	}
}

func TestFeedbackAmp1(t *testing.T) {
	input := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	ms, ps := FindMaxFeedbackSignal(input)

	if ms != 139629729 {
		t.Errorf("Expected Max Signal 139629729, received %v, with settings %v", ms, ps)
	}
}

func TestFeedbackAmp2(t *testing.T) {
	input := "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
	ms, ps := FindMaxFeedbackSignal(input)

	if ms != 18216 {
		t.Errorf("Expected Max Signal 18216, received %v, with settings %v", ms, ps)
	}
}
