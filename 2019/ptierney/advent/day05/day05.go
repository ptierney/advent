package day05

import (
	"advent/common"

	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var STD_INPUT int
var STD_OUTPUT io.Writer

var IP int

var Program []int
var StartingProgram []int

var ErrorReg bool

var StartingInput string

type ParameterMode int

const (
	Position  ParameterMode = 0
	Immediate ParameterMode = 1
)

type Opcode int

const (
	Add         Opcode = 1
	Multiply    Opcode = 2
	Input       Opcode = 3
	Output      Opcode = 4
	JumpIfTrue  Opcode = 5
	JumpIfFalse Opcode = 6
	LessThan    Opcode = 7
	Equals      Opcode = 8
	Halt        Opcode = 99
)

type Parameter struct {
	Value int
	Mode  ParameterMode
}

func NewParameter(value int, mode ParameterMode) *Parameter {
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

func OpcodeFromInt(o int) Opcode {
	return Opcode(o)
}

func ParameterModeFromInt(p int) ParameterMode {
	return ParameterMode(p)
}

// ParseInstruction parses the instruction starting at the IP, and returns
// the amount that the IP should be incremented by
func ParseInstruction() (*Instruction, int) {
	splitOpcode := splitNum(Program[IP])

	opcodeInt := splitOpcode[0] + splitOpcode[1]*10
	opcode := OpcodeFromInt(opcodeInt)

	parameters := make([]*Parameter, 0)

	var numParams int

	switch opcode {
	case Add, Multiply, LessThan, Equals:
		numParams = 3
	case JumpIfTrue, JumpIfFalse:
		numParams = 2
	case Input, Output:
		numParams = 1
	case Halt:
		numParams = 0
	default:
		ErrorReg = true
		return nil, 0
	}

	for i := 1; i <= numParams; i++ {
		paramValue := Program[IP+i]
		paramModeInt := splitOpcode[i+1]
		paramMode := ParameterModeFromInt(paramModeInt)

		param := NewParameter(paramValue, paramMode)

		parameters = append(parameters, param)
	}

	ipOffset := numParams + 1

	return NewInstruction(opcode, parameters), ipOffset
}

func splitNum(num int) []int {
	splitted := make([]int, 5)

	splitted[0] = num % 10
	splitted[1] = (num / 10) % 10
	splitted[2] = (num / 100) % 10
	splitted[3] = (num / 1000) % 10
	splitted[4] = (num / 10000) % 10

	return splitted
}

func trueValueOfParameter(param *Parameter) int {
	if param.Mode == Immediate {
		return param.Value
	}

	if param.Mode != Position {
		panic("Unknown parameter mode")
	}

	return Program[param.Value]
}

// executeProgram returns true if it should continue running
func executeInstruction() bool {
	instruction, offset := ParseInstruction()

	// Check for fundamental error conditions
	if instruction == nil || ErrorReg == true {
		return false
	}

	if instruction.Opcode == Halt {
		return false
	}

	opcode := instruction.Opcode

	if opcode == Add || opcode == Multiply || opcode == LessThan || opcode == Equals {
		aParam := instruction.Parameters[0]
		bParam := instruction.Parameters[1]
		// This is a write parameter, so it must be in position mode
		cParam := instruction.Parameters[2]

		aValue := trueValueOfParameter(aParam)
		bValue := trueValueOfParameter(bParam)

		var result int

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

		Program[cParam.Value] = result
	} else if opcode == Input {
		aParam := instruction.Parameters[0]
		Program[aParam.Value] = STD_INPUT
	} else if opcode == Output {
		aParam := instruction.Parameters[0]
		value := trueValueOfParameter(aParam)
		fmt.Fprintf(STD_OUTPUT, "%v\n", value)
	} else if opcode == JumpIfTrue || opcode == JumpIfFalse {
		aParam := instruction.Parameters[0]
		bParam := instruction.Parameters[1]

		aValue := trueValueOfParameter(aParam)
		bValue := trueValueOfParameter(bParam)

		if opcode == JumpIfTrue {
			if aValue != 0 {
				IP = bValue
				offset = 0 // null the offset so we don't increase IP
			}
		} else if opcode == JumpIfFalse {
			if aValue == 0 {
				IP = bValue
				offset = 0 // null the offset so we don't increase IP
			}
		}
	} else {
		ErrorReg = true
		return false
	}

	IP += offset

	return true
}

func LoadProgram(input []string) {
	StartingInput = input[0]

	loadProgramString(StartingInput)
}

func loadProgramString(startInput string) {
	programStr := strings.Split(startInput, ",")

	Program = make([]int, len(programStr))

	for i, str := range programStr {
		result, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}

		Program[i] = result
	}

	StartingProgram = Program
}

func ResetProgram() {
	loadProgramString(StartingInput)
	IP = 0
	ErrorReg = false
}

func SetProgramAlarm() {
	SetProgramNounVerb(12, 2)
}

func SetProgramNounVerb(noun, verb int) {
	Program[1] = noun
	Program[2] = verb
}

func ExecuteProgram() {
	for {
		if executeInstruction() == false {
			break
		}
	}
}

func SolveProblem() int {
	IP = 0

	ExecuteProgram()

	return Program[0]
}

func SolvePart1() int {
	input := common.GetInput("day05/input")
	LoadProgram(input)

	IP = 0

	STD_INPUT = 1
	STD_OUTPUT = os.Stdout

	return 0
}
