package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(fn string) []string {
	file, err := os.Open(fn)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputList []string

	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	return inputList
}

func parseInput(lines []string) {
	Samples = make([]*Sample, 0)

	for i := 0; i < len(lines); {
		s := NewSample(lines[i : i+3])

		Samples = append(Samples, s)

		i += 4
	}
}

var Registers []int
var Opcode map[string]func(int, int, int)
var OpcodeNumber map[int]string

var Samples []*Sample

func setup() {
	Registers = make([]int, 4)

	for i, _ := range Registers {
		Registers[i] = 0
	}

	Opcode = make(map[string]func(int, int, int))
	OpcodeNumber = make(map[int]string)

	Opcode["addr"] = func(a, b, c int) {
		Registers[c] = Registers[a] + Registers[b]
	}
	Opcode["addi"] = func(a, b, c int) {
		Registers[c] = Registers[a] + b
	}
	Opcode["mulr"] = func(a, b, c int) {
		Registers[c] = Registers[a] * Registers[b]
	}
	Opcode["muli"] = func(a, b, c int) {
		Registers[c] = Registers[a] * b
	}
	Opcode["banr"] = func(a, b, c int) {
		Registers[c] = Registers[a] & Registers[b]
	}
	Opcode["bani"] = func(a, b, c int) {
		Registers[c] = Registers[a] & b
	}
	Opcode["borr"] = func(a, b, c int) {
		Registers[c] = Registers[a] | Registers[b]
	}
	Opcode["bori"] = func(a, b, c int) {
		Registers[c] = Registers[a] | b
	}
	Opcode["setr"] = func(a, b, c int) {
		Registers[c] = Registers[a]
	}
	Opcode["seti"] = func(a, b, c int) {
		Registers[c] = a
	}
	Opcode["gtir"] = func(a, b, c int) {
		if a > Registers[b] {
			Registers[c] = 1
		} else {
			Registers[c] = 0
		}
	}
	Opcode["gtri"] = func(a, b, c int) {
		if Registers[a] > b {
			Registers[c] = 1
		} else {
			Registers[c] = 0
		}
	}
	Opcode["gtrr"] = func(a, b, c int) {
		if Registers[a] > Registers[b] {
			Registers[c] = 1
		} else {
			Registers[c] = 0
		}
	}
	Opcode["eqir"] = func(a, b, c int) {
		if a == Registers[b] {
			Registers[c] = 1
		} else {
			Registers[c] = 0
		}
	}
	Opcode["eqri"] = func(a, b, c int) {
		if Registers[a] == b {
			Registers[c] = 1
		} else {
			Registers[c] = 0
		}
	}
	Opcode["eqrr"] = func(a, b, c int) {
		if Registers[a] == Registers[b] {
			Registers[c] = 1
		} else {
			Registers[c] = 0
		}
	}
}

type Sample struct {
	BeforeRegisters []int
	AfterRegisters  []int

	A int
	B int
	C int

	Op int
}

func NewSample(lines []string) *Sample {
	sample := new(Sample)
	sample.BeforeRegisters = make([]int, 4)
	sample.AfterRegisters = make([]int, 4)

	beforeFields := strings.Fields(lines[0])
	afterFields := strings.Fields(lines[2])

	for i, _ := range sample.BeforeRegisters {
		s := beforeFields[i+1]

		s = strings.TrimLeft(s, "[")
		s = strings.TrimRight(s, "]")
		s = strings.TrimRight(s, ",")

		sample.BeforeRegisters[i], _ = strconv.Atoi(s)
	}

	for i, _ := range sample.AfterRegisters {
		s := afterFields[i+1]

		s = strings.TrimLeft(s, "[")
		s = strings.TrimRight(s, "]")
		s = strings.TrimRight(s, ",")

		sample.AfterRegisters[i], _ = strconv.Atoi(s)
	}

	opFields := strings.Fields(lines[1])

	sample.Op, _ = strconv.Atoi(opFields[0])
	sample.A, _ = strconv.Atoi(opFields[1])
	sample.B, _ = strconv.Atoi(opFields[2])
	sample.C, _ = strconv.Atoi(opFields[3])

	return sample
}

func printRegisters() {
	fmt.Printf("Registers: [%v, %v, %v, %v]\n",
		Registers[0], Registers[1], Registers[2], Registers[3])
}

func registersAreEqual(r1, r2 []int) bool {
	for i, _ := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}

func setRegisters(r []int) {
	for i, _ := range Registers {
		Registers[i] = r[i]
	}
}

func main() {
	setup()
	parseInput(getInput("input_states"))

	totalMatches := 0

	for _, sample := range Samples {

		thisMatches := 0

		//var matchedOpcodeName string
		//var matchedOpcodeNumber int

		for _, execFunc := range Opcode {
			setRegisters(sample.BeforeRegisters)

			execFunc(sample.A, sample.B, sample.C)

			if registersAreEqual(Registers, sample.AfterRegisters) {
				thisMatches++
				//matchedOpcodeName = key
				//matchedOpcodeNumber = sample.Op
			}
		}

		if thisMatches >= 3 {
			totalMatches++
		}
	}

	fmt.Printf("Total at least three: %v", totalMatches)
}
