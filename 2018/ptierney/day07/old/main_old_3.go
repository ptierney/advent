package main

import (
	"bufio"
	"container/list"
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

func parseLine(line string) (prereq string, step string) {
	fields := strings.Fields(line)

	return fields[1], fields[7]
}

type Step struct {
	ID            string
	Children      []*Step
	Prerequisites []*Step
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

var input []string

func getChildren(stepName string) *list.List {
	children := list.New()

	for _, line := range input {
		step, prereq := parseLine(line)

		if prereq != stepName {
			continue
		}

		children.PushBack(step)
	}

	return children
}

func getPrereqs(stepName string) *list.List {
	prereqs := list.New()

	for _, line := range input {
		step, prereq := parseLine(line)

		if step != stepName {
			continue
		}

		prereqs.PushBack(prereq)
	}

	return prereqs
}

func getRoot() string {
	_, rootID := parseLine(input[0])

	for {
		for _, line := range input {
			step, prereq := parseLine(line)

			if rootID == step {
				rootID = prereq
				break
			}
		}

		break
	}

	return rootID
}

func getNextAvailable() (string, *list.Element) {
	a := make([]string, 0)

	for e := availableQueue.Front(); e != nil; e = e.Next() {
		a = append(a, e.Value.(string))
	}

	sort.Strings(a)

	next := a[0]

	var removed *list.Element

	for e := availableQueue.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == next {
			availableQueue.Remove(e)
			removed = e
			break
		}
	}

	return next, removed
}

func addChildren(children *list.List) {

	for e := children.Front(); e != nil; e = e.Next() {

		id := e.Value.(string)

		inQueue := false

		inOutput := false
		// check if it's in the output
		for e3 := output.Front(); e3 != nil; e3 = e3.Next() {
			if e3.Value.(string) == id {
				inOutput = true
				break
			}
		}

		if inOutput == true {
			continue
		}

		// check to make sure it's not already in the queue
		for e2 := availableQueue.Front(); e2 != nil; e2 = e2.Next() {
			id2 := e2.Value.(string)

			if id == id2 {
				inQueue = true
				break
			}
		}

		// check to make sure all of the rereqs are satisfied
		prereqs := getPrereqs(id)
		prereqsSatisfied := make(map[string]bool)

		for p := prereqs.Front(); p != nil; p = p.Next() {
			prereqsSatisfied[p.Value.(string)] = false
		}

		for p := prereqs.Front(); p != nil; p = p.Next() {
			for e2 := output.Front(); e2 != nil; e2 = e2.Next() {
				if p.Value.(string) == e2.Value.(string) {
					prereqsSatisfied[p.Value.(string)] = true
					break
				}
			}
		}

		allSatisfied := true

		for _, value := range prereqsSatisfied {
			if value == false {
				allSatisfied = false
				break
			}
		}

		if inQueue == false {
			if allSatisfied == true {
				availableQueue.PushBack(id)
				bufferQueue.Remove(e)
			} else {
				bufferQueue.PushBack(id)
			}
		}
	}
}

func checkBufferQueue() {
	for e := bufferQueue.Front(); e != nil; {
		id := e.Value.(string)

		inOutput := false
		// check if it's in the output
		for e3 := output.Front(); e3 != nil; e3 = e3.Next() {
			if e3.Value.(string) == id {
				inOutput = true
				break
			}
		}

		if inOutput == true {
			continue
		}

		prereqs := getPrereqs(id)
		prereqsSatisfied := make(map[string]bool)

		for p := prereqs.Front(); p != nil; p = p.Next() {
			prereqsSatisfied[p.Value.(string)] = false
		}

		for p := prereqs.Front(); p != nil; p = p.Next() {
			for e2 := output.Front(); e2 != nil; e2 = e2.Next() {
				if p.Value.(string) == e2.Value.(string) {
					prereqsSatisfied[p.Value.(string)] = true
					break
				}
			}
		}

		allSatisfied := true

		for _, value := range prereqsSatisfied {
			if value == false {
				allSatisfied = false
				break
			}
		}

		if allSatisfied == true {
			availableQueue.PushBack(e.Value.(string))

			next := e.Next()
			bufferQueue.Remove(e)
			e = next
		} else {
			e = e.Next()
		}
	}
}

var output *list.List
var availableQueue *list.List
var bufferQueue *list.List

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(string))
	}

	fmt.Printf("\n")
}

func main() {
	input = getInput()

	output = list.New()
	availableQueue = list.New()
	bufferQueue = list.New()

	rootID := getRoot()
	output.PushBack(rootID)

	availableQueue.PushBackList(getChildren(rootID))

	printList(availableQueue)

	for availableQueue.Len() > 0 {
		next, _ := getNextAvailable()

		output.PushBack(next)

		// parse buffer queue
		checkBufferQueue()

		children := getChildren(next)

		addChildren(children)
	}

	for e := output.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value.(string))
	}

	fmt.Printf("\n")
}
