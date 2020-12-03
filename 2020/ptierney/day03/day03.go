package day03

import (
	"advent/common"

	"fmt"
)

type TileType int

const (
	Empty TileType = 0
	Tree  TileType = 1
)

type Tile struct {
	Type TileType
}

func TileFromChar(c string) *Tile {
	t := new(Tile)

	if c == "." {
		t.Type = Empty
	} else if c == "#" {
		t.Type = Tree
	} else {
		panic("Unknown type")
	}

	return t
}

func CreateTileGrid(input []string) [][]*Tile {
	gridHeight := len(input)

	inputWidth := len(input[0])
	replication := 100

	gridWidth := inputWidth * replication

	grid := make([][]*Tile, gridWidth)

	for x := 0; x < gridWidth; x++ {
		grid[x] = make([]*Tile, gridHeight)
		for y := 0; y < gridHeight; y++ {
			line := input[y]
			lineChar := string(line[x%inputWidth])

			grid[x][y] = TileFromChar(lineChar)
		}
	}

	return grid
}

func Part1() {
	input := common.GetInput("day03/input")

	grid := CreateTileGrid(input)

	numTrees := 0

	xPos := 0
	yPos := 0

	for yPos < len(input) {
		if grid[xPos][yPos].Type == Tree {
			numTrees++
		}

		xPos += 3
		yPos += 1
	}

	fmt.Printf("Num Trees = %v\n", numTrees)
}

func CountTrees(grid [][]*Tile, xSlope, ySlope int) int {
	numTrees := 0

	xPos := 0
	yPos := 0

	for yPos < len(grid[0]) {
		if grid[xPos][yPos].Type == Tree {
			numTrees++
		}

		xPos += xSlope
		yPos += ySlope
	}

	return numTrees
}

func Part2() {
	input := common.GetInput("day03/input")

	grid := CreateTileGrid(input)

	trees1 := CountTrees(grid, 1, 1)
	trees2 := CountTrees(grid, 3, 1)
	trees3 := CountTrees(grid, 5, 1)
	trees4 := CountTrees(grid, 7, 1)
	trees5 := CountTrees(grid, 1, 2)

	fmt.Printf("Num Trees = %v\n", trees1*trees2*trees3*trees4*trees5)
}
