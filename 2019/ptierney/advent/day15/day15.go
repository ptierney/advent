package day15

import (
	"advent/common"

	"fmt"
	"math/rand"
)

type TileType int

const (
	Unknown TileType = iota
	Empty   TileType = iota
	Wall    TileType = iota
	Oxygen  TileType = iota
)

type Direction int

const (
	North Direction = 1
	South Direction = 2
	West  Direction = 3
	East  Direction = 4
)

type Tile struct {
	Type TileType

	Visited bool

	SearchVisited bool
	CorrectPath   bool
}

func NewTile() *Tile {
	t := new(Tile)
	t.Type = Unknown
	t.Visited = false
	t.SearchVisited = false
	t.CorrectPath = false
	return t
}

func (t *Tile) ToString() string {
	if t.CorrectPath == true {
		return "P"
	}

	switch t.Type {
	case Empty:
		return "."
	case Wall:
		return "#"
	case Oxygen:
		return "O"
	case Unknown:
		return "-"
	}

	panic("Unknown Tile Type in To String")
}

var Maze [][]*Tile

var WorldDim int = 500

var RobotComputer *Computer

var RobotX int
var RobotY int

var RobotStartX int
var RobotStartY int

var RobotEndX int
var RobotEndY int

var OxygenX int
var OxygenY int

var UnknownTiles map[*Tile]*Tile

func Part1() {
	input := common.GetInput("day15/input")

	RobotStartX = WorldDim / 2
	RobotStartY = WorldDim / 2

	RobotX = RobotStartX
	RobotY = RobotStartY

	RobotComputer = NewComputerFromInput(input)

	CreateEmptyMaze()

	MapMaze()

	_ = RecursiveSolve(RobotStartX, RobotStartY)

	var sum int = 0

	for _, mr := range Maze {
		for _, t := range mr {
			if t.CorrectPath {
				sum++
			}
		}
	}

	fmt.Printf("Steps: %v", sum)

	//PrintMazeAroundRobot()
}

func Part2() {
	input := common.GetInput("day15/input")

	CreateEmptyMaze()

	MapMaze()

	SetOxygenXY()

	// Modify the recursive function to find the longest path

}

func SetOxygenXY() {
	for x, mr := range Maze {
		for y, t := range mr {
			if t.Type == Oxygen {
				OxygenX = x
				OxygenY = y
			}
		}
	}
}

func MoveRobot(d Direction) TileType {
	RobotComputer.SetInput(int64(d))

	out, halt, err := RobotComputer.ExecuteUntilOutput()

	if err != nil {
		panic(err)
	}

	if halt == true {
		panic("Should not halt")
	}

	RobotComputer.ClearOutput()

	switch out {
	case 0:
		return Wall
	case 1:
		return Empty
	case 2:
		return Oxygen
	}

	panic("Unknown output code")
}

func MapMaze() {
	s := GetSurroundingTiles()

	for _, t := range s {
		UnknownTiles[t] = t
	}

	for len(UnknownTiles) > 0 {
		MoveRandom()
	}
}

func MoveRandom() {
	dirInt := rand.Int()%4 + 1
	dir := Direction(dirInt)

	tt := MoveRobot(dir)

	var targetX int
	var targetY int

	switch dir {
	case North:
		targetX = RobotX
		targetY = RobotY - 1
	case South:
		targetX = RobotX
		targetY = RobotY + 1
	case West:
		targetX = RobotX - 1
		targetY = RobotY
	case East:
		targetX = RobotX + 1
		targetY = RobotY
	}

	targetTile := Maze[targetX][targetY]

	if targetTile.Type == Unknown {
		targetTile.Visited = true
		targetTile.Type = tt

		if tt == Oxygen {
			RobotEndX = targetX
			RobotEndY = targetY
		}

		delete(UnknownTiles, targetTile)
	} else {
		if targetTile.Type != tt {
			panic("Maze and intcomputer mismatch")
		}
	}

	if tt != Wall {
		RobotX = targetX
		RobotY = targetY

		s := GetSurroundingTiles()

		for _, t := range s {
			if t.Type == Unknown {
				UnknownTiles[t] = t
			}
		}
	}
}

func GetSurroundingTiles() []*Tile {
	tiles := make([]*Tile, 4)

	tiles[0] = Maze[RobotX+1][RobotY]
	tiles[1] = Maze[RobotX-1][RobotY]
	tiles[2] = Maze[RobotX][RobotY+1]
	tiles[3] = Maze[RobotX][RobotY-1]

	return tiles
}

func PrintMazeAroundRobot() {
	width := 100
	height := 75

	startX := common.Max(0, RobotX-width/2)
	startY := common.Max(0, RobotY-height/2)

	endX := common.Min(RobotX+width/2, WorldDim)
	endY := common.Min(RobotY+height/2, WorldDim)

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {

			if x == RobotStartX && y == RobotStartY {
				fmt.Printf("X")
			} else {
				fmt.Printf("%v", Maze[x][y].ToString())
			}
		}
		fmt.Printf("\n")
	}
}

func CreateEmptyMaze() {
	Maze = make([][]*Tile, WorldDim)
	for i := 0; i < WorldDim; i++ {
		Maze[i] = make([]*Tile, WorldDim)
		for j := 0; j < WorldDim; j++ {
			Maze[i][j] = NewTile()
		}
	}

	UnknownTiles = make(map[*Tile]*Tile)
}

// From https://en.wikipedia.org/wiki/Maze_solving_algorithm
func RecursiveSolve(x, y int) bool {
	if x == RobotEndX && y == RobotEndY {
		return true
	}

	if Maze[x][y].Type == Wall {
		return false
	}

	if Maze[x][y].SearchVisited == true {
		return false
	}

	Maze[x][y].SearchVisited = true

	if x != 0 {
		if RecursiveSolve(x-1, y) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	if x != WorldDim-1 {
		if RecursiveSolve(x+1, y) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	if y != 0 {
		if RecursiveSolve(x, y-1) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	if y != WorldDim-1 {
		if RecursiveSolve(x, y+1) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	return false
}

// TODO: Update this
func RecursiveTraverse(x, y int) bool {
	if x == RobotEndX && y == RobotEndY {
		return true
	}

	if Maze[x][y].Type == Wall {
		return false
	}

	if Maze[x][y].SearchVisited == true {
		return false
	}

	Maze[x][y].SearchVisited = true

	if x != 0 {
		if RecursiveSolve(x-1, y) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	if x != WorldDim-1 {
		if RecursiveSolve(x+1, y) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	if y != 0 {
		if RecursiveSolve(x, y-1) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	if y != WorldDim-1 {
		if RecursiveSolve(x, y+1) == true {
			Maze[x][y].CorrectPath = true
			return true
		}
	}

	return false
}
