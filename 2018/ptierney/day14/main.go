package main

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
)

var startingRecipes string = "37"

var inputString string = "074501"

//var inputString string = "59414"
var inputNumber int = 74501

var recipes *list.List

var targetRecipes *list.List

var elf1 *list.Element
var elf2 *list.Element

func parseInput() {
	parseStringToList(startingRecipes, recipes)
}

func parseTargetRecipes() {
	parseStringToList(inputString, targetRecipes)
}

func parseStringToList(s string, l *list.List) {
	for _, r := range s {
		val, err := strconv.Atoi(string(r))

		if err != nil {
			log.Fatal(err)
		}

		l.PushBack(val)
	}
}

func printRecipes() {
	for e := recipes.Front(); e != nil; e = e.Next() {
		if e == elf1 {
			fmt.Printf("(%v) ", e.Value)
		} else if e == elf2 {
			fmt.Printf("[%v] ", e.Value)
		} else {
			fmt.Printf("%v ", e.Value)
		}
	}

	fmt.Printf("\n")
}

// This method doesn't work, because it doesn't modify the input list element
// I think I would need to pass in a **lsit.Element
func tickElf(e *list.Element) {
	numTicks := e.Value.(int) + 1

	for i := 0; i < numTicks; i++ {
		if e.Next() == nil {
			e = recipes.Front()
		} else {
			e = e.Next()
		}
	}
}

func tickRecipes() {
	combined := elf1.Value.(int) + elf2.Value.(int)

	if combined < 10 {
		recipes.PushBack(combined)
	} else {
		recipes.PushBack(int(combined / 10))
		recipes.PushBack(combined % 10)
	}

	// tick elf 1
	numTicks := elf1.Value.(int) + 1

	for i := 0; i < numTicks; i++ {
		if elf1.Next() == nil {
			elf1 = recipes.Front()
		} else {
			elf1 = elf1.Next()
		}
	}

	numTicks = elf2.Value.(int) + 1

	for i := 0; i < numTicks; i++ {
		if elf2.Next() == nil {
			elf2 = recipes.Front()
		} else {
			elf2 = elf2.Next()
		}
	}

	// this is how it should be, just need to figure out how to
	// correctly modify a pointer
	//tickElf(elf1)
	//tickElf(elf2)
}

func subListMatchesTarget(l1 *list.Element) bool {
	l2 := targetRecipes.Front()

	for {
		if l1.Value.(int) != l2.Value.(int) {
			return false
		}

		l1 = l1.Next()
		l2 = l2.Next()

		// we've reached the end of the target recipes
		if l2 == nil {
			return true
		}

		// l1 is shorter than the target recipes
		if l1 == nil {
			return false
		}
	}

	log.Fatal("Error matching lists")

	return false
}

func checkForTargetMatch() (bool, *list.Element) {
	// we've either added one or two elements on the end of the

	targetLenth := targetRecipes.Len()

	// seek back target length + 1
	farBackElement := recipes.Back()

	for i := 0; i < targetLenth+1; i++ {
		farBackElement = farBackElement.Prev()

		if farBackElement == nil {
			return false, nil
		}
	}

	if subListMatchesTarget(farBackElement) == true {
		return true, farBackElement
	}

	farBackElement = farBackElement.Next()

	if subListMatchesTarget(farBackElement) == true {
		return true, farBackElement
	}

	return false, nil
}

func main() {
	recipes = list.New()
	targetRecipes = list.New()

	parseInput()
	parseTargetRecipes()

	elf1 = recipes.Front()
	elf2 = elf1.Next()

	// printRecipes()

	// desiredRecipes := inputNumber + 10

	// for recipes.Len() <= desiredRecipes {
	// 	tickRecipes()
	// }

	// elem := recipes.Front()

	// for i := 0; i < inputNumber; i++ {
	// 	elem = elem.Next()
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%v", elem.Value)
	// 	elem = elem.Next()
	// }

	// fmt.Printf("\n")

	// for i := 0; i < 10; i++ {
	// 	tickRecipes()
	// 	printRecipes()
	// }

	var matchedStart *list.Element = nil

	for {
		tickRecipes()

		doesMatch, elem := checkForTargetMatch()

		if doesMatch == true {
			matchedStart = elem
			break
		}
	}

	// count the number to the left

	var leftCount int = 0

	for e := recipes.Front(); e != nil; e = e.Next() {
		if matchedStart == e {
			break
		}

		leftCount++
	}

	fmt.Printf("Num to left: %v\n", leftCount)
}
