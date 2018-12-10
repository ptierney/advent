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

func parseLine(line string) (base rune, before rune) {
	fields := strings.Fields(line)

	var b rune
	var n rune

	for _, r := range fields[1] {
		b = r
	}

	for _, r := range fields[7] {
		n = r
	}

	return b, n
}

type Step struct {
	ID       rune
	Children []*Step
}

func newStep(id rune) *Step {
	s := new(Step)
	s.ID = id
	s.Children = make([]*Step, 0)

	return s
}

var input []string

func findChildren(id rune) []rune {
	children := make([]rune, 0)

	for _, line := range input {
		this, next := parseLine(line)

		if this != id {
			continue
		}

		children = append(children, next)
	}

	return children
}

func setChildren(step *Step, children []rune) {
	if len(children) == 0 {
		return
	}

	for _, childID := range children {
		childStep := newStep(childID)

		step.Children = append(step.Children, childStep)

		childChildren := findChildren(childID)

		setChildren(childStep, childChildren)
	}
}

// Sorting Code
type ByStep []*Step

func (a ByStep) Len() int           { return len(a) }
func (a ByStep) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStep) Less(i, j int) bool { return a[i].ID < a[j].ID }

var finalOrder []rune

func findOrder(step *Step) {
	finalOrder = append(finalOrder, step.ID)

	orderedChildren := make([]rune, len(step.Children))
	copy(orderedChildren, step.Children)

	sort.Sort(ByStep(orderedChildren))

	for _, c := range orderedChildren {
		findOrder(c)
	}

}

func main() {
	input = getInput()

	// find base
	baseID, _ := parseLine(input[0])

	for {
		for _, line := range input {
			this, next := parseLine(line)

			if baseID == next {
				baseID = this
				break
			}
		}

		break
	}

	baseStep := newStep(baseID)

	setChildren(baseStep, findChildren(baseStep.ID))

	// parse the graph

	finalOrder = make([]rune, 0)

	findOrder(baseStep)

	fmt.Printf("Base is: %c\n", baseStep.ID)
}
