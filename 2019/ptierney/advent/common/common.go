package common

import (
	"bufio"
	"log"
	"os"
)

var rootDir string = "/home/patrick/dev/advent/2019/ptierney/advent"

func GetInput(dir string) []string {
	file, err := os.Open(rootDir + "/" + dir + "/input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputList := make([]string, 0)

	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	return inputList
}
