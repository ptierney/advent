package day09

import (
	"advent/common"

	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var STD_INPUT int64
var STD_OUTPUT io.Writer
var STD_OUTPUT_INT64 []int64

var IP int64
var RelBase int64

var Program []int64

var ErrorReg bool

var StartingInput string

type ParameterMode int64

var MemorySize int = 10000000

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

// ParseInstruction parses the instruction starting at the IP, and returns
// the amount that the IP should be incremented by
func ParseInstruction() (*Instruction, int64) {
	splitOpcode := splitNum(Program[IP])

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
		ErrorReg = true
		return nil, 0
	}

	for i := int64(1); i <= numParams; i++ {
		paramValue := Program[IP+i]
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

func trueValueOfParameter(param *Parameter) int64 {
	if param.Mode == Immediate {
		return param.Value
	} else if param.Mode == Position {
		return Program[param.Value]
	} else if param.Mode == Relative {
		return Program[param.Value+RelBase]
	} else {
		panic("Unknown parameter mode")
	}

	return 0
}

func writeValueToParameter(value int64, param *Parameter) {
	if param.Mode == Relative {
		Program[param.Value+RelBase] = value
	} else if param.Mode == Position {
		Program[param.Value] = value
	} else {
		panic("Trying to write to immediate mode")
	}
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

		writeValueToParameter(result, cParam)
	} else if opcode == Input {
		aParam := instruction.Parameters[0]
		writeValueToParameter(STD_INPUT, aParam)
	} else if opcode == Output {
		aParam := instruction.Parameters[0]
		value := trueValueOfParameter(aParam)
		fmt.Fprintf(STD_OUTPUT, "%v\n", value)
		STD_OUTPUT_INT64 = append(STD_OUTPUT_INT64, value)
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
	} else if opcode == AdjRelBase {
		aParam := instruction.Parameters[0]
		aValue := trueValueOfParameter(aParam)

		RelBase += aValue
	} else {
		ErrorReg = true
		return false
	}

	IP += offset

	return true
}

func SetProgram(prog string) {
	StartingInput = prog
}

func LoadProgram(input []string) {
	SetProgram(input[0])

	loadProgramString(StartingInput)
}

func loadProgramString(startInput string) {
	programStr := strings.Split(startInput, ",")

	Program = make([]int64, MemorySize)

	for i, str := range programStr {
		result, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}

		Program[i] = int64(result)
	}
}

func ResetProgram() {
	loadProgramString(StartingInput)
	IP = 0
	ErrorReg = false
	STD_INPUT = 0
	STD_OUTPUT = os.Stdout
	STD_OUTPUT_INT64 = make([]int64, 0)
	RelBase = 0
}

func SetProgramAlarm() {
	SetProgramNounVerb(12, 2)
}

func SetProgramNounVerb(noun, verb int64) {
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

func SolveProblem() int64 {
	IP = 0

	ExecuteProgram()

	return Program[0]
}

func PrintInts(ints []int64) {
	for _, i := range ints {
		fmt.Printf("%v\n", i)
	}
}

func GetBOOSTKeycode() {
	input := common.GetInput("day09/input")

	SetProgram(input[0])
	ResetProgram()

	STD_INPUT = 1

	fmt.Printf("Starting\n")

	ExecuteProgram()

	PrintInts(STD_OUTPUT_INT64)
}

func GetCoords() {
	input := common.GetInput("day09/input")

	SetProgram(input[0])
	ResetProgram()

	STD_INPUT = 2

	fmt.Printf("Starting\n")

	ExecuteProgram()

	PrintInts(STD_OUTPUT_INT64)
}
