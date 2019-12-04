package main

import (
	"advent/common"
	"advent/day03"

	"fmt"
)

func main() {
	input := common.GetInput("day03/input")
	part1 := day03.SolveProblem(input)
	fmt.Printf("Part 1: %v\n", part1)

	part2 := day03.SolvePart2(input)
	fmt.Printf("Part 2: %v\n", part2)
}
