package day02

import (
	"advent/common"

	"fmt"
	"strconv"
	"strings"
)

type PasswordEntry struct {
	PolicyChar string
	PolicyMin  int
	PolicyMax  int

	Password string
}

func ParseLines(lines []string) []*PasswordEntry {
	entries := make([]*PasswordEntry, 0)

	for _, line := range lines {
		entry := ParseLine(line)

		entries = append(entries, entry)
	}

	return entries
}

func ParseLine(line string) *PasswordEntry {
	passwordEntry := new(PasswordEntry)

	split1 := strings.Split(line, " ")

	policyNumStr := split1[0]
	policyCharStr := split1[1]

	passwordEntry.Password = split1[2]

	split2 := strings.Split(policyNumStr, "-")

	passwordEntry.PolicyMin, _ = strconv.Atoi(split2[0])
	passwordEntry.PolicyMax, _ = strconv.Atoi(split2[1])

	split3 := strings.Split(policyCharStr, ":")

	passwordEntry.PolicyChar = split3[0]

	return passwordEntry
}

func EntryIsValid(entry *PasswordEntry) bool {
	charCount := strings.Count(entry.Password, entry.PolicyChar)

	return charCount >= entry.PolicyMin && charCount <= entry.PolicyMax
}

type PasswordEntry2 struct {
	PolicyChar string
	Position1  int
	Position2  int

	Password string
}

func ParseLines2(lines []string) []*PasswordEntry2 {
	entries := make([]*PasswordEntry2, 0)

	for _, line := range lines {
		entry := ParseLine2(line)

		entries = append(entries, entry)
	}

	return entries
}

func ParseLine2(line string) *PasswordEntry2 {
	passwordEntry := new(PasswordEntry2)

	split1 := strings.Split(line, " ")

	policyNumStr := split1[0]
	policyCharStr := split1[1]

	passwordEntry.Password = split1[2]

	split2 := strings.Split(policyNumStr, "-")

	passwordEntry.Position1, _ = strconv.Atoi(split2[0])
	passwordEntry.Position2, _ = strconv.Atoi(split2[1])

	split3 := strings.Split(policyCharStr, ":")

	passwordEntry.PolicyChar = split3[0]

	return passwordEntry
}

func EntryIsValid2(entry *PasswordEntry2) bool {
	charPos1 := string(entry.Password[entry.Position1-1])
	charPos2 := string(entry.Password[entry.Position2-1])

	pos1Valid := charPos1 == entry.PolicyChar
	pos2Valid := charPos2 == entry.PolicyChar

	if pos1Valid && pos2Valid {
		return false
	}

	return pos1Valid || pos2Valid
}

func Part1() {
	input := common.GetInput("day02/input")

	entries := ParseLines(input)

	numValid := 0

	for _, entry := range entries {
		if EntryIsValid(entry) {
			numValid += 1
		}
	}

	fmt.Printf("Valid = %v\n", numValid)
}

func Part2() {
	input := common.GetInput("day02/input")

	entries := ParseLines2(input)

	numValid := 0

	for _, entry := range entries {
		if EntryIsValid2(entry) {
			numValid += 1
		}
	}

	fmt.Printf("Valid = %v\n", numValid)

}
