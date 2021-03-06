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

func parseLine(line string) (prereq string, step string) {
	fields := strings.Fields(line)

	return fields[1], fields[7]
}

func parseInput(lines []string) {
	stepsMap := make(map[string]*Step)

	stepsAll := make([]string, 0)

	// create all the steps
	for _, line := range lines {
		prereq, step := parseLine(line)

		stepsAll = append(stepsAll, prereq)
		stepsAll = append(stepsAll, step)
	}

	for _, step := range stepsAll {
		_, containsStep := stepsMap[step]

		if containsStep == true {
			continue
		}

		s := NewStep(step)
		stepsMap[step] = s
	}

	// set all the Prerequisites and Children
	for _, line := range lines {
		prereq, step := parseLine(line)

		prereqStep := stepsMap[prereq]
		stepStep := stepsMap[step]

		stepStep.AddPrerequisite(prereqStep)
		prereqStep.AddChild(stepStep)
	}

	for _, value := range stepsMap {
		Steps = append(Steps, value)
	}
}

type ByStep []*Step

func (a ByStep) Len() int           { return len(a) }
func (a ByStep) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStep) Less(i, j int) bool { return a[i].ID[0] < a[j].ID[0] }

type Step struct {
	ID            string
	Children      []*Step
	Prerequisites []*Step

	Done       bool
	InProgress bool
}

func NewStep(id string) *Step {
	s := new(Step)
	s.ID = id
	s.Done = false
	s.InProgress = false

	s.Children = make([]*Step, 0)
	s.Prerequisites = make([]*Step, 0)

	return s
}

func (s *Step) AddPrerequisite(prereq *Step) {
	s.Prerequisites = append(s.Prerequisites, prereq)
}

func (s *Step) AddChild(c *Step) {
	s.Children = append(s.Children, c)
}

func (s *Step) PrerequisitesAreDone() bool {
	for _, p := range s.Prerequisites {
		if p.Done == false {
			return false
		}
	}

	return true
}

func (s *Step) IsDoAble() bool {
	return s.Done == false && s.InProgress == false
}

func (s *Step) GetWorkTime() int {
	// ASCII value normalized to 1, plus 60 seconds
	//return int(s.ID[0]) - 64 + 60

	return int(s.ID[0]) - 64 + 60
}

var Steps []*Step
var FinalOrder []*Step
var ElfPool []*Elf

// Elfs do work
type Elf struct {
	CurrentStep *Step

	RemainingTime int

	ID int
}

func NewElf(id int) *Elf {
	e := new(Elf)

	e.CurrentStep = nil

	e.ID = id

	return e
}

func (e *Elf) AssignTask(s *Step) {
	e.CurrentStep = s

	e.RemainingTime = s.GetWorkTime()

	s.InProgress = true
}

func (e *Elf) IsOccupied() bool {
	return e.CurrentStep != nil
}

func (e *Elf) TaskGlyph() string {
	if e.IsOccupied() {
		return e.CurrentStep.ID
	} else {
		return "."
	}

}

func (e *Elf) Tick() {
	//fmt.Printf("Elf %v is working on %v\n", e.ID, e.TaskGlyph())

	if e.CurrentStep == nil {
		return
	}

	e.RemainingTime -= 1

	if e.RemainingTime > 0 {
		return
	}

	FinalOrder = append(FinalOrder, e.CurrentStep)

	e.CurrentStep.Done = true
	e.CurrentStep.InProgress = false

	e.CurrentStep = nil
}

func CreateElfPool() {
	for i := 0; i < 5; i++ {
		ElfPool = append(ElfPool, NewElf(i))
	}
}

func printSteps() {
	printStepList(Steps)
}

func printFinalOrder() {
	fmt.Printf("Final Order: ")
	printStepList(FinalOrder)
}

func FinalOrderString() string {
	if len(FinalOrder) == 0 {
		return " "
	}

	var s string = ""

	for _, elem := range FinalOrder {
		s = s + elem.ID
	}

	return s
}

func printStepList(sl []*Step) {
	for _, s := range sl {
		fmt.Printf("%v", s.ID)
	}

	fmt.Printf("\n")
}

func getDoAbleRoots() []*Step {
	roots := make([]*Step, 0)

	for _, s := range Steps {
		if len(s.Prerequisites) == 0 && s.IsDoAble() {
			roots = append(roots, s)
		}
	}

	return roots
}

func getRoot() *Step {
	possibleRoots := getDoAbleRoots()

	sort.Sort(ByStep(possibleRoots))

	if len(possibleRoots) == 0 {
		return nil
	} else {
		return possibleRoots[0]
	}
}

func doStep(s *Step) {
	FinalOrder = append(FinalOrder, s)

	s.Done = true
}

func calculateNextStep() *Step {
	doneSteps := make([]*Step, 0)

	for _, s := range Steps {
		if s.Done == true {
			doneSteps = append(doneSteps, s)
		}
	}

	possibleSteps := make(map[string]*Step)

	for _, s1 := range doneSteps {
		for _, s2 := range s1.Children {
			if s2.PrerequisitesAreDone() == true && s2.IsDoAble() {
				possibleSteps[s2.ID] = s2
			}
		}
	}

	// the graph can have multiple "roots"
	roots := getDoAbleRoots()

	for _, r := range roots {
		possibleSteps[r.ID] = r
	}

	possibleStepsList := make([]*Step, 0)

	for _, s := range possibleSteps {
		possibleStepsList = append(possibleStepsList, s)
	}

	// there are no possible next steps
	if len(possibleStepsList) == 0 {
		return nil
	}

	sort.Sort(ByStep(possibleStepsList))

	return possibleStepsList[0]
}

func AllStepsAreDone() bool {
	for _, s := range Steps {
		if s.Done == false {
			return false
		}
	}

	return true
}

func AllElfsAreOccupied() bool {
	for _, e := range ElfPool {
		if e.IsOccupied() == false {
			return false
		}
	}

	return true
}

func AssignStepToElfPool(step *Step) {
	if AllElfsAreOccupied() {
		return
	}

	var workElf *Elf = nil

	for _, e := range ElfPool {
		if e.IsOccupied() == false {
			workElf = e
			break
		}
	}

	workElf.AssignTask(step)
}

func TickElfs() {
	for _, e := range ElfPool {
		e.Tick()
	}
}

func main() {
	parseInput(getInput())

	CreateElfPool()

	var seconds int = 0

	for AllStepsAreDone() == false {

		for AllElfsAreOccupied() == false {
			s := calculateNextStep()

			if s == nil {
				break
			}

			AssignStepToElfPool(s)
		}

		// fmt.Printf("%v  %v  %v  %v\n", seconds,
		// 	ElfPool[0].TaskGlyph(), ElfPool[1].TaskGlyph(),
		// 	FinalOrderString())

		TickElfs()

		seconds++
	}

	printFinalOrder()

	fmt.Printf("Total Time: %v\n", seconds)
}
