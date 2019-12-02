package main

import (
	"advent/common"
	"advent/day02"

	"fmt"
)

func main() {
	input := common.GetInput("day02/input")
	day02.LoadProgram(input)
	day02.SetProgramAlarm()
	part1 := day02.SolveProblem()
	fmt.Printf("Part 1: %v\n", part1)

	part2 := day02.FindNounVerb()

	fmt.Printf("NounVerb: %v\n", part2)
}
