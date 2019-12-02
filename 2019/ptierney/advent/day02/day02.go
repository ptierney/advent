package day02

import (
	"fmt"
	"strconv"
	"strings"
)

var IP int

var Program []int
var StartingProgram []int

var ErrorReg bool

var StartingInput string

// executeProgram returns true if it should continue running
func executeInstruction() bool {
	opcode := Program[IP]

	if opcode == 99 {
		return false
	}

	aRegLocation := Program[IP+1]
	bRegLocation := Program[IP+2]

	aReg := Program[aRegLocation]
	bReg := Program[bRegLocation]

	var result int

	if opcode == 1 {
		result = aReg + bReg
	} else if opcode == 2 {
		result = aReg * bReg
	} else {
		ErrorReg = true
		//panic(fmt.Errorf("Received unknown opcode"))
		return false
	}

	resultLocation := Program[IP+3]

	Program[resultLocation] = result

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

func SolveProblem() int {
	IP = 0

	for {
		if executeInstruction() == false {
			break
		}

		IP += 4
	}

	return Program[0]
}

func FindNounVerb() int {
	goal := 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			ResetProgram()
			SetProgramNounVerb(noun, verb)

			result := SolveProblem()

			// this was an error program
			if ErrorReg == true {
				continue
			}

			if result == goal {
				return 100*noun + verb
			}
		}
	}

	panic(fmt.Errorf("Could not find noun verb"))

	return 0
}
