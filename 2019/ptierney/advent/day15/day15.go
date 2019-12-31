package day15

import (
	"advent/common"

	"fmt"
)

type TileType int

const (
	Empty  TileType = iota
	Wall   TileType = iota
	Oxygen TileType = iota
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
}

func NewTile() *Tile {
	t := new(Tile)
	t.Type = Empty
	return t
}

func (t *Tile) ToString() string {
	switch t.TileType {
	case Empty:
		return "."
	case Wall:
		return "#"
	case Oxygen:
		return "O"
	}

	panic("Unknown Tile Type in To String")
}

var Maze [][]*Tile

var WorldDim int = 10000

var RobotComputer *Computer

var RobotX int
var RobotY int

func Part1() {
	input := common.GetInput("day15/input")

	RobotX = WorldDim / 2
	RobotY = WorldDim / 2

	RobotComputer = NewComputerFromInput(input)

	fmt.Printf("%v", input)
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
	// head north until he hits a wall

	// left hand rule maze

}

func MoveRandom() {

}

func PrintMazeAroundRobot() {
	width := 50
	height := 25

	startX := common.Max(0, RobotX-width/2)
	startY := common.Max(0, RobotY-height/2)

	endX := common.Min(RobotX+width/2, WorldDim)
	endY := common.Min(RobotY+height/2, WorldDim)

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			fmt.Printf("%v", Maze[x][y].ToString())
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
}
