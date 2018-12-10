package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func parseInputTuple(tupleString string) (x int, y int) {
	positionSplit := strings.Split(tupleString, "=")
	positionCore := positionSplit[1]
	positionCore = strings.TrimLeft(positionCore, "<")
	positionCore = strings.TrimRight(positionCore, ">")

	positionsStrings := strings.Split(positionCore, ",")

	x, err := strconv.Atoi(strings.TrimSpace(positionsStrings[0]))

	if err != nil {
		log.Fatal(err)
	}

	y, err = strconv.Atoi(strings.TrimSpace(positionsStrings[1]))

	if err != nil {
		log.Fatal(err)
	}

	return
}

func parseLine(line string) (xPos int, yPos int, xVel int, yVel int) {
	l := strings.TrimRight(line, ">")
	l2 := strings.Split(l, ">")

	positionCore := l2[0]
	velocityCore := strings.TrimSpace(l2[1])

	xPos, yPos = parseInputTuple(positionCore)
	xVel, yVel = parseInputTuple(velocityCore)

	return
}

type Point struct {
	X         int
	Y         int
	xVelocity int
	yVelocity int
}

func NewPoint(xPos int, yPos int, xVel int, yVel int) *Point {
	p := new(Point)

	p.X = xPos
	p.Y = yPos
	p.xVelocity = xVel
	p.yVelocity = yVel

	return p
}

func (p *Point) tick() {
	p.X += p.xVelocity
	p.Y += p.yVelocity
}

var points map[int]*Point

var DisplayWidth int = 100
var DisplayHeight int = 11

func DisplayPoints(xOffset int, yOffset int) {
	for j := 0; j < DisplayHeight; j++ {
		for i := 0; i < DisplayWidth; i++ {

			pointInSpace := false

			for _, p := range points {
				if i == (p.X-xOffset) && j == (p.Y-yOffset) {
					pointInSpace = true
					break
				}
			}

			if pointInSpace == true {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func AllPointsInDisplay() (inDisplay bool, xOffset int, yOffset int) {
	minX := points[0].X
	maxX := points[0].X
	minY := points[0].Y
	maxY := points[0].Y

	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}

		if p.X > maxX {
			maxX = p.X
		}

		if p.Y < minY {
			minY = p.Y
		}

		if p.Y > maxY {
			maxY = p.Y
		}
	}

	inDisplay = ((maxX - minX) <= DisplayWidth) && ((maxY - minY) <= DisplayHeight)

	xOffset = minX
	yOffset = minY

	return
}

func main() {
	input := getInput()

	points = make(map[int]*Point)

	for i, line := range input {
		p := NewPoint(parseLine(line))

		points[i] = p
	}

	for i := 0; i < 100000; i++ {
		for _, p := range points {
			p.tick()
		}

		inDisplay, x, y := AllPointsInDisplay()

		if inDisplay == false {
			continue
		}

		DisplayPoints(x, y)

		fmt.Printf("%v\n", i+1)
	}
}
