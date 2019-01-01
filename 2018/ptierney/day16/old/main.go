package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

}

var Registers []int

var Opcodes map[string]func(int, int, int)

func setup() {
	Registers = make([]int, 4)

	for i, _ := range Registers {
		Registers[i] = 0
	}

	Opcodes
}

type Opcode interface {
	Execute(a, b, c int)
}

// type BaseOpcode struct {
// }

type ADDR struct{}

func (op *ADDR) Execute(a, b, c int) {
	Registers[c] = Registers[a] + Registers[b]
}

type ADDI struct{}

func (op *ADDI) Execute(a, b, c int) {
	Registers[c] = Registers[a] + b
}

type MULR struct{}

func (op *MULR) Execute(a, b, c int) {
	Registers[c] = Registers[a] * Registers[b]
}

type MULI struct{}

func (op *MULI) Execute(a, b, c int) {
	Registers[c] = Registers[a] * b
}

type BANR struct{}

func (op *BANR) Execute(a, b, c int) {
	Registers[c] = Registers[a] & Registers[b]
}

type BANI struct{}

func (op *BANI) Execute(a, b, c int) {
	Registers[c] = Registers[a] & b
}

type BORR struct{}

func (op *BORR) Execute(a, b, c int) {
	Registers[c] = Registers[a] | Registers[b]
}

type BORI struct{}

func (op *BORI) Execute(a, b, c int) {
	Registers[c] = Registers[a] | b
}

type SETR struct{}

func (op *SETR) Execute(a, b, c int) {
	Registers[c] = Registers[a]
}

type SETI struct{}

func (op *SETI) Execute(a, b, c int) {
	Registers[c] = a
}

type GTIR struct{}

func (op *SETI) Execute(a, b, c int) {

}

func main() {
	setup()
	parseInput(getInput("input_states"))

	fmt.Printf("Hello world\n")

}
