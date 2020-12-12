package day06

import (
	"advent/common"

	"fmt"
	"strings"
)

type QuestionMap map[string]bool
type QuestionList []string

func Part1() {
	input := common.GetInput("day06/input")

	groupQuestions := make([]*QuestionMap, 1)

	currentGroup := new(QuestionMap)
	*currentGroup = make(QuestionMap)
	groupQuestions[0] = currentGroup

	for _, line := range input {
		if line == "" {
			currentGroup = new(QuestionMap)
			*currentGroup = make(QuestionMap)
			groupQuestions = append(groupQuestions, currentGroup)
		}

		for _, c := range line {
			q := string(c)

			(*currentGroup)[q] = true
		}
	}

	sum := 0

	for _, group := range groupQuestions {
		for _, _ = range *group {
			sum++
		}
	}

	fmt.Printf("Sum = %v\n", sum)
}

func Part2() {
	input := common.GetInput("day06/input")

	groupQuestions := make([]*QuestionList, 1)

	currentGroup := new(QuestionList)
	*currentGroup = make(QuestionList, 0)
	groupQuestions[0] = currentGroup

	atFirstPersonInGroup := true

	for _, line := range input {
		if line == "" {
			currentGroup = new(QuestionList)
			*currentGroup = make(QuestionList, 0)
			groupQuestions = append(groupQuestions, currentGroup)

			atFirstPersonInGroup = true

			continue
		}

		if atFirstPersonInGroup == true {
			questions := common.StringArrayFromString(line)

			for _, q := range questions {
				(*currentGroup) = append(*currentGroup, q)
			}

			atFirstPersonInGroup = false

			continue
		}

		for i := 0; i < len(*currentGroup); {
			q := (*currentGroup)[i]

			// if this question is contained in the line
			if strings.Index(line, q) != -1 {
				i++
				continue
			}

			// remove the element at this index
			(*currentGroup) = append((*currentGroup)[:i], (*currentGroup)[i+1:]...)
			i = 0
		}

		atFirstPersonInGroup = false
	}

	sum := 0

	for _, group := range groupQuestions {
		// from the puzzle:
		//Duplicate answers to the same question don't count extra; each question counts at most once.
		unique := make(map[string]bool)

		for _, q := range *group {
			unique[q] = true
		}

		sum += len(unique)

		//sum += len(*group)
	}

	fmt.Printf("Sum = %v\n", sum)
}
