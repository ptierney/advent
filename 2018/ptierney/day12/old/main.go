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

var plantOffset int = 10000

func parseInitialState(str string) {
	state = make([]string, plantOffset)

	for i := 0; i < plantOffset; i++ {
		state[i] = "."
	}

	for _, s := range str {
		state = append(state, string(s))
	}

	for i := 0; i < plantOffset; i++ {
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

func tickPlants() {
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
}

func findMatchingRule(subSlice []string) *Rule {
	for _, rule := range rules {
		if matches(subSlice, rule.Pattern) {
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

	for i = 0; i < 50000000000; i++ {
		if i%1000000000 == 0 {
			fmt.Printf("%v\n", i)
		}

		tickPlants()
	}

	fmt.Printf("Plants Sum: %v\n\n", sumPlants())

	for _, s := range state {
		fmt.Printf("%v", s)
	}

	fmt.Printf("\n")
}
