package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"unicode"
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

func compactPolymer(polymer *list.List) {
	finished := false

	for finished == false {

		for unit := polymer.Front(); unit != nil; unit = unit.Next() {
			// check for the end condition
			if unit.Next() == nil {
				finished = true
				break
			}

			// check if there's a match

			thisUnitRune, _ := unit.Value.(rune)
			nextUnitRune, _ := unit.Next().Value.(rune)

			if unicode.ToLower(thisUnitRune) == unicode.ToLower(nextUnitRune) {
				if (unicode.IsUpper(thisUnitRune) && unicode.IsLower(nextUnitRune)) ||
					(unicode.IsLower(thisUnitRune) && unicode.IsUpper(nextUnitRune)) {

					// it's a match
					polymer.Remove(unit.Next())
					polymer.Remove(unit)
					break
				}
			}
		}
	}
}

func removeLetterFromPolymer(polymer *list.List, character rune) {
	for unit := polymer.Front(); unit != nil; {
		thisUnitRune, _ := unit.Value.(rune)

		if unicode.ToLower(thisUnitRune) != character {
			unit = unit.Next()
			continue
		}

		next := unit.Next()

		polymer.Remove(unit)

		unit = next
	}
}

func copyList(l *list.List) *list.List {
	newList := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		newList.PushBack(e.Value)
	}

	return newList
}

var letterLength map[rune]chan int

func computePolymer(polymer *list.List, letter rune, length chan int) {
	polymerCopy := copyList(polymer)

	removeLetterFromPolymer(polymerCopy, letter)

	compactPolymer(polymerCopy)

	length <- polymerCopy.Len()
}

func main() {
	inputSequence := getInput()

	polymer := list.New()

	for _, char := range inputSequence[0] {
		polymer.PushBack(char)
	}

	compactPolymer(polymer)

	fmt.Printf("Shorted Polymer Length: %v\n", polymer.Len())

	// Part 2

	letterLength = make(map[rune]chan int)

	for i := 97; i < 123; i++ {
		letterLength[rune(i)] = make(chan int)
	}

	for letter, lengthChan := range letterLength {
		go computePolymer(polymer, letter, lengthChan)
	}

	shortestLen := -1

	for _, length := range letterLength {
		var lengthValue int = <-length

		if shortestLen < 0 {
			shortestLen = lengthValue
			continue
		}

		if lengthValue < shortestLen {
			shortestLen = lengthValue
		}
	}

	fmt.Printf("Smallest Length: %v\n", shortestLen)
}
