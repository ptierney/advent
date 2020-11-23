package common

import (
	"bufio"
	"log"
	"os"
)

var rootDir string = "/home/patrick/dev/advent/2020/ptierney"

func GetInput(subPath string) []string {
	file, err := os.Open(rootDir + "/" + subPath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputList := make([]string, 0)

	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	return inputList
}

// From https://github.com/kindermoumoute/adventofcode/blob/master/pkg/math.go

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type nullableInt struct {
	value int
}

func Max(values ...int) int {
	if len(values) == 0 {
		panic("no value in max function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || max.value < value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func Min(values ...int) int {
	if len(values) == 0 {
		panic("no value in min function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || max.value > value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func Sum(values ...int) int {
	if len(values) == 0 {
		panic("no value in sum function")
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func StringArrayFromString(s string) []string {
	arr := make([]string, 0)

	for _, rune := range s {
		arr = append(arr, string(rune))
	}

	return arr
}
