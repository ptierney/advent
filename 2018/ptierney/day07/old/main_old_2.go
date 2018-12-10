package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getInput() []string {
	//file, err := os.Open("input")
	file, err := os.Open("test_input")
	//file, err := os.Open("test_input_2")

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

func parseLine(line string) (step string, prereq string) {
	fields := strings.Fields(line)

	return fields[7], fields[1]
}

type Step struct {
	ID       string
	Children []*Step
}

func NewStep(id string) *Step {
	s := new(Step)
	s.ID = id
	return s
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		fmt.Printf("Visit all: %v\n", items)

		for _, item := range items {
			fmt.Printf("for item: %v\n", item)

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				fmt.Printf("Appending: %v\n", item)
				order = append(order, item)
			}
		}
	}

	var keys []string

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	input := getInput()

	steps := make(map[string][]string)

	stepPtrs := make(map[string]*Step)

	for _, line := range input {
		step, prereq := parseLine(line)

		steps[step] = append(steps[step], prereq)
	}

	var keys []string
	for key := range steps {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%v: %v\n", key, steps[key])
	}

	//fmt.Printf("%v\n", steps)

	order := topoSort(steps)

	for _, s := range order {
		fmt.Printf("%v", s)
	}

	fmt.Printf("\n")
}
