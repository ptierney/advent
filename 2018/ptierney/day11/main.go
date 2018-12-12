package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func getInput() []string {
	file, err := os.Open("input")
	//file, err := os.Open("test_input")

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

var grid [][]int

func PowerLevelAtPoint(x int, y int) int {
	rackID := x + 10

	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID

	powerLevelString := strconv.Itoa(powerLevel)

	numDigits := utf8.RuneCountInString(powerLevelString)

	if numDigits < 3 {
		powerLevel = 0
	} else {
		powerLevelDigit := string(powerLevelString[numDigits-3])

		num, err := strconv.Atoi(powerLevelDigit)

		if err != nil {
			log.Fatal(err)
		}

		powerLevel = num
	}

	powerLevel -= 5

	return powerLevel
}

func GetSquareSum(x int, y int, size int) int {
	sum := 0

	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			sum += grid[i][j]
		}
	}

	return sum
}

var serialNumber int = 2568

var sideDim int = 300

func main() {

	grid = make([][]int, sideDim)

	for i := 0; i < sideDim; i++ {
		grid[i] = make([]int, sideDim)
	}

	// set grid values
	for x := 0; x < sideDim; x++ {
		for y := 0; y < sideDim; y++ {
			grid[x][y] = PowerLevelAtPoint(x+1, y+1)
		}
	}

	squareSums := make(map[string]int)

	for squareSize := 1; squareSize <= sideDim; squareSize++ {
		for x := 0; x < sideDim-(squareSize-1); x++ {
			for y := 0; y < sideDim-(squareSize-1); y++ {
				key := fmt.Sprintf("%v,%v,%v", x+1, y+1, squareSize)

				squareSums[key] = GetSquareSum(x, y, squareSize)
			}
		}
	}

	maxKey := "1,1,1"
	maxSum := squareSums[maxKey]

	for key, val := range squareSums {
		if val > maxSum {
			maxKey = key
			maxSum = val
		}
	}

	fmt.Printf("Max Index: %v\n", maxKey)

}
