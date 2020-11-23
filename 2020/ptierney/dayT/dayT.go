package dayT

import (
	"advent/common"

	"fmt"
	"strings"
)

func Part1() {
	input := common.GetInput("dayT/input")

	turnsString := input[0]

	turnsList := strings.Split(turnsString, ", ")

	fmt.Printf("%v", turnsList)
}

func Part2() {

}
