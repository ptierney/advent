package day05

import (
	"advent/common"

	"fmt"
	"math"
)

func BinaryToRC(binarySeat string) (int, int) {
	bin := common.StringArrayFromString(binarySeat)

	row := 0

	for i := 0; i < 7; i++ {
		exponent := 6 - i

		if bin[i] == "B" {
			row += int(math.Pow(2, float64(exponent)))
		}
	}

	column := 0

	for i := 7; i < 10; i++ {
		exponent := 2 - (i - 7)

		if bin[i] == "R" {
			column += int(math.Pow(2, float64(exponent)))
		}
	}

	return row, column
}

func RCToSeatID(row, column int) int {
	return row*8 + column
}

func Part1() {
	input := common.GetInput("day05/input")

	highestID := 0

	for _, seat := range input {
		row, column := BinaryToRC(seat)

		seatID := RCToSeatID(row, column)

		if seatID > highestID {
			highestID = seatID
		}
	}

	fmt.Printf("Highest = %v\n", highestID)
}

func RCKey(r, c int) string {
	return fmt.Sprintf("R:%v-C:%v", r, c)
}

func Part2() {
	input := common.GetInput("day05/input")

	seatsTaken := make(map[string]bool)

	for r := 0; r < 128; r++ {
		for c := 0; c < 8; c++ {
			seatsTaken[RCKey(r, c)] = false
		}
	}

	for _, seat := range input {
		row, column := BinaryToRC(seat)

		seatsTaken[RCKey(row, column)] = true
	}

	for k, v := range seatsTaken {
		if v == false {
			fmt.Printf("%v\n", k)
		}
	}
}
