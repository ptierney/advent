package day13

import (
	"advent/common"

	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var MemorySize int = 10000000

type Computer struct {
	STD_INPUT    []int64
	inputCounter int

	STD_OUTPUT       io.Writer
	STD_OUTPUT_INT64 []int64

	IP      int64
	RelBase int64

	Program []int64

	ErrorReg bool
	HaltFlag bool

	StartingInput string
}

func NewComputer() *Computer {
	c := new(Computer)

	c.NullVariables()

	c.Program = make([]int64, 0)

	c.StartingInput = ""

	return c
}

type ParameterMode int64

const (
	Position  ParameterMode = 0
	Immediate ParameterMode = 1
	Relative  ParameterMode = 2
)

type Opcode int64

const (
	Add         Opcode = 1
	Multiply    Opcode = 2
	Input       Opcode = 3
	Output      Opcode = 4
	JumpIfTrue  Opcode = 5
	JumpIfFalse Opcode = 6
	LessThan    Opcode = 7
	Equals      Opcode = 8
	AdjRelBase  Opcode = 9
	Halt        Opcode = 99
)

type Parameter struct {
	Value int64
	Mode  ParameterMode
}

func NewParameter(value int64, mode ParameterMode) *Parameter {
	p := new(Parameter)

	p.Value = value
	p.Mode = mode

	return p
}

type Instruction struct {
	Opcode     Opcode
	Parameters []*Parameter
}

func NewInstruction(opcode Opcode, params []*Parameter) *Instruction {
	i := new(Instruction)

	i.Opcode = opcode
	i.Parameters = params

	return i
}

func OpcodeFromInt64(o int64) Opcode {
	return Opcode(o)
}

func ParameterModeFromInt64(p int64) ParameterMode {
	return ParameterMode(p)
}

// ParseInstruction parses the instruction starting at the c.IP, and returns
// the amount that the c.IP should be incremented by
func (c *Computer) ParseInstruction() (*Instruction, int64) {
	splitOpcode := splitNum(c.Program[c.IP])

	opcodeInt64 := splitOpcode[0] + splitOpcode[1]*10
	opcode := OpcodeFromInt64(opcodeInt64)

	parameters := make([]*Parameter, 0)

	var numParams int64

	switch opcode {
	case Add, Multiply, LessThan, Equals:
		numParams = 3
	case JumpIfTrue, JumpIfFalse:
		numParams = 2
	case Input, Output, AdjRelBase:
		numParams = 1
	case Halt:
		numParams = 0
	default:
		c.ErrorReg = true
		return nil, 0
	}

	for i := int64(1); i <= numParams; i++ {
		paramValue := c.Program[c.IP+i]
		paramModeInt64 := splitOpcode[i+1]
		paramMode := ParameterModeFromInt64(paramModeInt64)

		param := NewParameter(paramValue, paramMode)

		parameters = append(parameters, param)
	}

	ipOffset := numParams + 1

	return NewInstruction(opcode, parameters), ipOffset
}

func splitNum(num int64) []int64 {
	splitted := make([]int64, 5)

	splitted[0] = num % 10
	splitted[1] = (num / 10) % 10
	splitted[2] = (num / 100) % 10
	splitted[3] = (num / 1000) % 10
	splitted[4] = (num / 10000) % 10

	return splitted
}

func (c *Computer) trueValueOfParameter(param *Parameter) int64 {
	if param.Mode == Immediate {
		return param.Value
	} else if param.Mode == Position {
		return c.Program[param.Value]
	} else if param.Mode == Relative {
		return c.Program[param.Value+c.RelBase]
	} else {
		panic("Unknown parameter mode")
	}

	return 0
}

func (c *Computer) writeValueToParameter(value int64, param *Parameter) {
	if param.Mode == Relative {
		c.Program[param.Value+c.RelBase] = value
	} else if param.Mode == Position {
		c.Program[param.Value] = value
	} else {
		panic("Trying to write to immediate mode")
	}
}

// executec.Program returns true if it should continue running
func (c *Computer) executeInstruction() bool {
	instruction, offset := c.ParseInstruction()

	// Check for fundamental error conditions
	if instruction == nil || c.ErrorReg == true {
		return false
	}

	if instruction.Opcode == Halt {
		c.HaltFlag = true
		return false
	}

	opcode := instruction.Opcode

	if opcode == Add || opcode == Multiply || opcode == LessThan || opcode == Equals {
		aParam := instruction.Parameters[0]
		bParam := instruction.Parameters[1]
		// This is a write parameter, so it must be in position mode
		cParam := instruction.Parameters[2]

		aValue := c.trueValueOfParameter(aParam)
		bValue := c.trueValueOfParameter(bParam)

		var result int64

		switch opcode {
		case Add:
			result = aValue + bValue
		case Multiply:
			result = aValue * bValue
		case LessThan:
			if aValue < bValue {
				result = 1
			} else {
				result = 0
			}
		case Equals:
			if aValue == bValue {
				result = 1
			} else {
				result = 0
			}
		}

		c.writeValueToParameter(result, cParam)
	} else if opcode == Input {
		aParam := instruction.Parameters[0]
		i := c.GetNextInput()
		c.writeValueToParameter(i, aParam)
	} else if opcode == Output {
		aParam := instruction.Parameters[0]
		value := c.trueValueOfParameter(aParam)
		//fmt.Fprintf(c.STD_OUTPUT, "%v\n", value)
		c.STD_OUTPUT_INT64 = append(c.STD_OUTPUT_INT64, value)
	} else if opcode == JumpIfTrue || opcode == JumpIfFalse {
		aParam := instruction.Parameters[0]
		bParam := instruction.Parameters[1]

		aValue := c.trueValueOfParameter(aParam)
		bValue := c.trueValueOfParameter(bParam)

		if opcode == JumpIfTrue {
			if aValue != 0 {
				c.IP = bValue
				offset = 0 // null the offset so we don't increase c.IP
			}
		} else if opcode == JumpIfFalse {
			if aValue == 0 {
				c.IP = bValue
				offset = 0 // null the offset so we don't increase c.IP
			}
		}
	} else if opcode == AdjRelBase {
		aParam := instruction.Parameters[0]
		aValue := c.trueValueOfParameter(aParam)

		c.RelBase += aValue
	} else {
		c.ErrorReg = true
		return false
	}

	c.IP += offset

	return true
}

func (c *Computer) SetInput(i int64) {
	c.STD_INPUT = make([]int64, 1)
	c.STD_INPUT[0] = i
	c.inputCounter = 0
}

func (c *Computer) AddInput(i int64) {
	c.STD_INPUT = append(c.STD_INPUT, i)
}

func (c *Computer) GetNextInput() int64 {
	if len(c.STD_INPUT) == 0 {
		c.ErrorReg = true
		panic("Could not get next input")
		return 0
	}

	i := c.STD_INPUT[c.inputCounter]

	if c.inputCounter <= (len(c.STD_INPUT) - 2) {
		c.inputCounter++
	}

	return i
}

func (c *Computer) SetProgram(prog string) {
	c.StartingInput = prog
}

func (c *Computer) LoadProgramString(prog string) {
	c.SetProgram(prog)

	c.loadProgramString(c.StartingInput)
}

func (c *Computer) LoadProgramFromInput(input []string) {
	c.SetProgram(input[0])

	c.loadProgramString(c.StartingInput)
}

func (c *Computer) loadProgramString(startInput string) {
	programStr := strings.Split(startInput, ",")

	c.Program = make([]int64, MemorySize)

	for i, str := range programStr {
		result, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}

		c.Program[i] = int64(result)
	}
}

func (c *Computer) NullVariables() {
	c.STD_INPUT = make([]int64, 0)
	c.inputCounter = 0

	c.STD_OUTPUT = os.Stdout
	c.STD_OUTPUT_INT64 = make([]int64, 0)

	c.IP = 0
	c.RelBase = 0

	c.ErrorReg = false
	c.HaltFlag = false
}

func (c *Computer) ResetProgram() {
	c.loadProgramString(c.StartingInput)

	c.NullVariables()
}

func (c *Computer) SetProgramAlarm() {
	c.SetProgramNounVerb(12, 2)
}

func (c *Computer) SetProgramNounVerb(noun, verb int64) {
	c.Program[1] = noun
	c.Program[2] = verb
}

func (c *Computer) ExecuteProgram() {
	for {
		if c.executeInstruction() == false {
			break
		}
	}
}

func (c *Computer) ClearOutput() {
	c.STD_OUTPUT_INT64 = make([]int64, 0)
}

func (c *Computer) ExecuteUntilOutput() (int64, bool, error) {
	for {
		if c.executeInstruction() == false {
			if c.ErrorReg == true {
				return 0, false, fmt.Errorf("Encountered error during execution")
			} else {
				return 0, true, nil
			}
		}

		if len(c.STD_OUTPUT_INT64) > 0 {
			break
		}
	}

	if c.ErrorReg == true {
		err := fmt.Errorf("Encountered Error during execution")
		panic(err)
		return 0, false, err
	}

	return c.STD_OUTPUT_INT64[0], false, nil
}

func (c *Computer) SetMemoryAtAddress(addr int64, value int64) {
	c.Program[addr] = value
}

func (c *Computer) SolveProblem() int64 {
	c.IP = 0

	c.ExecuteProgram()

	return c.Program[0]
}

func PrintInts(ints []int64) {
	for _, i := range ints {
		fmt.Printf("%v\n", i)
	}
}

func (c *Computer) GetBOOSTKeycode() {
	input := common.GetInput("day09/input")

	c.SetProgram(input[0])
	c.ResetProgram()

	c.AddInput(1)

	fmt.Printf("Starting\n")

	c.ExecuteProgram()

	PrintInts(c.STD_OUTPUT_INT64)
}

func (c *Computer) GetCoords() {
	input := common.GetInput("day09/input")

	c.SetProgram(input[0])
	c.ResetProgram()

	c.AddInput(2)

	fmt.Printf("Starting\n")

	c.ExecuteProgram()

	PrintInts(c.STD_OUTPUT_INT64)
}
