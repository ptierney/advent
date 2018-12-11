package main

import (
	"container/list"
	"fmt"
)

var circle *list.List

func GetOneClockwiseElement(currentMarble *list.Element) *list.Element {
	if currentMarble.Next() == nil {
		return circle.Front()
	}

	return currentMarble.Next()
}

func GetOneCounterClockwiseElement(currentMarble *list.Element) *list.Element {
	if currentMarble.Prev() == nil {
		return circle.Back()
	}

	return currentMarble.Prev()
}

func AddMarble(currentMarble *list.Element, index int, playerID int) *list.Element {
	cw := GetOneClockwiseElement(currentMarble)

	if index%23 != 0 {
		return circle.InsertAfter(index, cw)

	}

	scoreTotal := index

	sevenCC := currentMarble

	for i := 0; i < 7; i++ {
		sevenCC = GetOneCounterClockwiseElement(sevenCC)
	}

	scoreTotal += sevenCC.Value.(int)

	currentMarble = GetOneClockwiseElement(sevenCC)

	circle.Remove(sevenCC)

	playerScores[playerID] += scoreTotal

	return currentMarble
}

var playerScores map[int]int

func main() {
	circle = list.New()
	playerScores = make(map[int]int)

	// settings
	numPlayers := 455
	maxMarble := 71223 * 100

	currentPlayer := 1

	currentMarble := circle.PushBack(0)

	for i := 1; i <= maxMarble; i++ {
		currentMarble = AddMarble(currentMarble, i, currentPlayer)

		currentPlayer += 1
		if currentPlayer > numPlayers {
			currentPlayer = 1
		}
	}

	maxScore := playerScores[1]

	for _, score := range playerScores {
		if score > maxScore {
			maxScore = score
		}
	}

	fmt.Printf("Max Score: %v\n", maxScore)
}
