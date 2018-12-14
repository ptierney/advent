package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() []string {
	//file, err := os.Open("input")
	file, err := os.Open("test_input")

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

var state string
var rules []*Rule

type Rule struct {
	Pattern string

	Result rune
}

func NewRule(pattern string, r string) *Rule {
	rule := new(Rule)

	rule.Pattern = pattern

	for _, rune := range r {
		rule.Result = rune
		break
	}

	return rule
}

func parseInput(input []string) {
	stateLine := input[0]
	stateFields := strings.Fields(stateLine)
	stateString := stateFields[2]

	parseInitialState(stateString)

	parseRules(input[2:])
}

var plantOffset int = 100

func parseInitialState(str string) {
	state = ""

	for i := 0; i < plantOffset; i++ {
		state += "."
	}

	state += str

	for i := 0; i < plantOffset; i++ {
		state += "."
	}
}

func parseRules(rulesStrings []string) {
	for _, rs := range rulesStrings {
		rsFields := strings.Fields(rs)

		rules = append(rules, NewRule(rsFields[0], rsFields[2]))
	}
}

func matches(s1 []string, s2 []string) bool {
	for i := 0; i < 5; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func tickPlants() {
	stateCopy := state

	for i := 0; i < len(state)-5; i++ {
		rule := findMatchingRule(state[i:(i + 5)])

		var newValue rune = '.'
		if rule != nil {
			newValue = rule.Result
		}

		stateCopy = stateCopy[:
		
		stateCopy[i+2] = newValue
	}

	copy(state, stateCopy)
}

func findMatchingRule(substring string) *Rule {
	for _, rule := range rules {
		if substring == rule.Pattern {
			return rule
		}
	}

	return nil
}

func sumPlants() int {
	sum := 0

	for i, val := range state {
		if val == "." {
			continue
		}

		sum += (i - plantOffset)
	}

	return sum
}

func main() {
	parseInput(getInput())

	var i uint64

	//for i = 0; i < 50000000000; i++ {
	for i = 0; i < 20; i++ {

		// if i%1000000000 == 0 {
		// 	fmt.Printf("%v\n", i)
		// }

		//tickPlants()
	}

	fmt.Printf("Plants Sum: %v\n\n", sumPlants())

	for _, s := range state {
		fmt.Printf("%v", s)
	}

	fmt.Printf("\n")
}
