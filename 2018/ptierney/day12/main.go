package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

var state []string
var rules []*Rule

var stateLookup map[string][]string

var steadyState bool

var plantOffset uint64 = 150

type Rule struct {
	Pattern []string

	Result string
}

func NewRule(pattern string, r string) *Rule {
	rule := new(Rule)

	rule.Pattern = make([]string, 0)

	for _, s := range pattern {
		rule.Pattern = append(rule.Pattern, string(s))
	}

	rule.Result = r

	return rule
}

func parseInput(input []string) {
	stateLine := input[0]
	stateFields := strings.Fields(stateLine)
	stateString := stateFields[2]

	parseInitialState(stateString)

	parseRules(input[2:])
}

func getKey(input []string) string {
	key := ""

	for _, s := range input {
		key += s
	}

	return key
}

func parseInitialState(str string) {
	state = make([]string, plantOffset)

	var i uint64
	for i = 0; i < plantOffset; i++ {
		state[i] = "."
	}

	for _, s := range str {
		state = append(state, string(s))
	}

	for i = 0; i < plantOffset; i++ {
		state = append(state, ".")
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

func statesMatch(s1 []string, s2 []string) bool {
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func tickPlants() {
	key := getKey(state)

	value, found := stateLookup[key]

	if found == true {
		if statesMatch(state, value) {
			steadyState = true
			return
		}

		copy(state, value)
		return
	}

	stateCopy := make([]string, len(state))

	copy(stateCopy, state)

	for i := 0; i < len(state)-5; i++ {
		rule := findMatchingRule(state[i:(i + 5)])

		newValue := "."
		if rule != nil {
			newValue = rule.Result
		}

		stateCopy[i+2] = newValue
	}

	copy(state, stateCopy)

	stateLookup[key] = state
}

func findMatchingRule(subSlice []string) *Rule {
	for _, rule := range rules {
		if matches(subSlice, rule.Pattern) {
			return rule
		}
	}

	return nil
}

func sumPlants() uint64 {
	var sum uint64 = 0

	var i uint64 = 0

	for i = 0; i < uint64(len(state)); i++ {
		if state[i] == "." {
			continue
		}

		sum += (i - plantOffset)
	}

	return sum
}

func printState() {
	for _, s := range state {
		fmt.Printf("%v", s)
	}

	fmt.Printf("  << Sum: %v", sumPlants())

	fmt.Printf("\n")
}

func main() {
	stateLookup = make(map[string][]string)
	steadyState = false

	parseInput(getInput())

	var i uint64

	for i = 0; i < 50000000000 && steadyState == false; i++ {
		tickPlants()

		fmt.Printf("%v: ", i+1)
		printState()
	}

	//fmt.Printf("Plants Sum: %v\n\n", sumPlants())

	//printState()

	//                         999999999414
	//                        3450000002337
	var answer uint64 = (uint64(50000000000)-uint64(102))*uint64(69) + uint64(9306)

	fmt.Printf("Answer: %v\n", answer)

}
