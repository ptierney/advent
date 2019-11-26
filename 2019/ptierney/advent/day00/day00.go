package day00

import (
	"advent/common"

	"log"
	"strconv"
)

func SumNumbers() int {
	input := common.GetInput("day00")

	sum := 0

	for _, line := range input {
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		sum += i
	}

	return sum
}
