package day04

import (
	"fmt"
)

func splitNum(num int) []int {
	splitted := make([]int, 6)

	splitted[5] = num % 10
	splitted[4] = (num / 10) % 10
	splitted[3] = (num / 100) % 10
	splitted[2] = (num / 1000) % 10
	splitted[1] = (num / 10000) % 10
	splitted[0] = (num / 100000) % 10

	return splitted
}

func neverDecrease(num int) bool {
	s := splitNum(num)

	for i := 0; i < 5; i++ {
		if s[i+1] < s[i] {
			return false
		}
	}

	return true
}

func twoAreTheSame(num int) bool {
	s := splitNum(num)

	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func hasSingleDouble(num int) bool {
	s := splitNum(num)

	for i := 0; i < 5; i++ {
		if s[i] != s[i+1] {
			continue
		}

		if i > 0 && s[i-1] == s[i] {
			continue
		}

		if i < 4 && s[i+2] == s[i] {
			continue
		}

		return true
	}

	return false
}

func valueCorrect(num int) bool {
	return neverDecrease(num) && twoAreTheSame(num)
}

func valueCorrectPart2(num int) bool {
	return neverDecrease(num) && hasSingleDouble(num)
}

func SolvePart1() {
	minValue := 264793
	maxValue := 803935

	var totalNum int = 0

	for i := minValue; i <= maxValue; i++ {
		if valueCorrect(i) == true {
			totalNum++
		}
	}

	fmt.Printf("%v\n", totalNum)
}

func SolvePart2() {
	minValue := 264793
	maxValue := 803935

	var totalNum int = 0

	for i := minValue; i <= maxValue; i++ {
		if valueCorrectPart2(i) == true {
			totalNum++
		}
	}

	fmt.Printf("%v\n", totalNum)
}
