package day04

import (
	"advent/common"

	"fmt"
	"strconv"
	"strings"
)

func ContainsField(fields []string, key string) bool {
	for _, field := range fields {
		fieldSplit := strings.Split(field, ":")
		fieldKey := fieldSplit[0]

		if fieldKey == key {
			return true
		}
	}

	return false
}

func GetFieldValue(fields []string, key string) string {
	for _, field := range fields {
		fieldSplit := strings.Split(field, ":")
		fieldKey := fieldSplit[0]
		fieldValue := fieldSplit[1]

		if fieldKey == key {
			return fieldValue
		}
	}

	panic("Could not find field value")
}

func PassportIsValid(passportStr string) bool {
	fields := strings.Fields(passportStr)

	return ContainsField(fields, "byr") &&
		ContainsField(fields, "iyr") &&
		ContainsField(fields, "eyr") &&
		ContainsField(fields, "hgt") &&
		ContainsField(fields, "hcl") &&
		ContainsField(fields, "ecl") &&
		ContainsField(fields, "pid")
}

func PassportValuesAreValid(passportStr string) bool {
	fields := strings.Fields(passportStr)

	birthYearStr := GetFieldValue(fields, "byr")

	birthYear, err := strconv.Atoi(birthYearStr)

	if err != nil {
		return false
	}

	if birthYear < 1920 || birthYear > 2002 {
		return false
	}

	issueYearStr := GetFieldValue(fields, "iyr")

	issueYear, err := strconv.Atoi(issueYearStr)

	if err != nil {
		return false
	}

	if issueYear < 2010 || issueYear > 2020 {
		return false
	}

	expYearStr := GetFieldValue(fields, "eyr")

	expYear, err := strconv.Atoi(expYearStr)

	if err != nil {
		return false
	}

	if expYear < 2020 || expYear > 2030 {
		return false
	}

	heightStr := GetFieldValue(fields, "hgt")

	heightUnit := heightStr[len(heightStr)-2:]

	height, err := strconv.Atoi(heightStr[:len(heightStr)-2])

	if err != nil {
		return false
	}

	if heightUnit == "in" {
		if height < 59 || height > 76 {
			return false
		}
	} else if heightUnit == "cm" {
		if height < 150 || height > 193 {
			return false
		}
	} else {
		return false
	}

	hairStr := GetFieldValue(fields, "hcl")

	if string(hairStr[0]) != "#" {
		return false
	}

	hairHex := hairStr[1:]

	if len(hairHex) != 6 {
		return false
	}

	validChars := "0123456789abcdefg"

	for _, c := range hairHex {
		if strings.Index(validChars, string(c)) == -1 {
			return false
		}
	}

	eyeColor := GetFieldValue(fields, "ecl")

	switch eyeColor {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	passportID := GetFieldValue(fields, "pid")

	if len(passportID) != 9 {
		return false
	}

	validChars = "0123456789"

	for _, c := range passportID {
		if strings.Index(validChars, string(c)) == -1 {
			return false
		}
	}

	return true
}

func Part1() {
	input := common.GetInput("day04/input")

	passportStrs := make([]string, 0)

	passport := ""
	for _, line := range input {
		if line == "" {
			passportStrs = append(passportStrs, passport)
			passport = ""
		}

		passport = passport + " " + line
	}

	if passport != "" {
		passportStrs = append(passportStrs, passport)
	}

	numValid := 0

	for _, p := range passportStrs {
		if PassportIsValid(p) {
			numValid++
		}
	}

	fmt.Printf("Num Valid = %v\n", numValid)
}

func Part2() {
	input := common.GetInput("day04/input")

	passportStrs := make([]string, 0)

	passport := ""
	for _, line := range input {
		if line == "" {
			passportStrs = append(passportStrs, passport)
			passport = ""
		}

		passport = passport + " " + line
	}

	if passport != "" {
		passportStrs = append(passportStrs, passport)
	}

	numValid := 0

	for _, p := range passportStrs {
		if PassportIsValid(p) == false {
			continue
		}

		if PassportValuesAreValid(p) == false {
			continue
		}

		numValid++
	}

	fmt.Printf("Num Valid = %v\n", numValid)

}
