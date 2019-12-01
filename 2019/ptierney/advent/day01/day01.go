package day01

import (
	"advent/common"

	"fmt"
	"strconv"
)

func SolveProblem() {
	input := common.GetInput("day01")

	var sum int = 0

	for _, iStr := range input {
		moduleMass, err := strconv.Atoi(iStr)

		if err != nil {
			panic(err)
		}

		moduleFuel := moduleMass/3 - 2

		sum += moduleFuel

		for {
			fuelFuel := moduleFuel/3 - 2

			if fuelFuel < 1 {
				break
			}

			sum += fuelFuel
			moduleFuel = fuelFuel
		}
	}

	fmt.Printf("Sum: %v\n", sum)
}
