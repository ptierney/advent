package day01

import (
	"advent/common"

	"fmt"
	"strconv"
)

func Part1() {
	input := common.GetInput("day01/input")

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			iInt, _ := strconv.Atoi(input[i])
			jInt, _ := strconv.Atoi(input[j])

			if iInt+jInt == 2020 {
				fmt.Printf("Value = %v\n", iInt*jInt)
			}
		}
	}
}

func Part2() {
	input := common.GetInput("day01/input")

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				iInt, _ := strconv.Atoi(input[i])
				jInt, _ := strconv.Atoi(input[j])
				kInt, _ := strconv.Atoi(input[k])

				if iInt+jInt+kInt == 2020 {
					fmt.Printf("Value = %v\n", iInt*jInt*kInt)
				}
			}
		}
	}
}
