package day13

import (
	"advent/common"

	"bufio"
	"fmt"
	"os"
	"os/exec"
	//"time"
)

type Tile struct {
	Type TileType
}

type TileType int

const (
	Empty   TileType = 0
	Wall    TileType = 1
	Block   TileType = 2
	HPaddle TileType = 3
	Ball    TileType = 4
)

func NewTile() *Tile {
	t := new(Tile)
	t.Type = Empty
	return t
}

func (t *Tile) SetType(tileType int) {
	var tt TileType

	switch tileType {
	case 0:
		tt = Empty
	case 1:
		tt = Wall
	case 2:
		tt = Block
	case 3:
		tt = HPaddle
	case 4:
		tt = Ball
	}

	t.Type = tt
}

func (t *Tile) ToString() string {
	switch t.Type {
	case Empty:
		return "."
	case Wall:
		return "W"
	case Block:
		return "B"
	case HPaddle:
		return "P"
	case Ball:
		return "*"
	}

	panic("Unknown type")
}

var Tiles [][]*Tile

var GameComputer *Computer

var WorldHeight int = 25
var WorldWidth int = 50

var Score int = 0

var LastBallX int = -1
var LastBallY int = -1

func Part1() {
	input := common.GetInput("day13/input")

	GameComputer = NewComputer()

	GameComputer.LoadProgramFromInput(input)

	InitTiles()

	for {
		halt, xPos, yPos, tileType := ExecuteTile()

		if halt == true {
			break
		}

		Tiles[xPos][yPos].SetType(tileType)
	}

	fmt.Printf("Block tiles: %v\n", CountBlocks())

	PrintTiles()
}

func Part2() {
	input := common.GetInput("day13/input")

	GameComputer = NewComputer()

	GameComputer.LoadProgramFromInput(input)

	GameComputer.SetMemoryAtAddress(0, 2)

	InitTiles()

	SetJoystickNeutral()

	setupComplete := false

	for {
		halt, xPos, yPos, tileType := ExecuteTile()

		if halt == true {
			break
		}

		if xPos == -1 && yPos == 0 {
			Score = tileType
		} else {
			Tiles[xPos][yPos].SetType(tileType)

			if TileType(tileType) == HPaddle {
				setupComplete = true
			}
		}

		if setupComplete == false {
			continue
		}

		AutoMovePaddle()

		ClearScreen()
		PrintTiles()

		//gtime.Sleep(150 * time.Millisecond)
	}
}

func ExecuteTile() (bool, int, int, int) {
	xPos, didHalt, err := GameComputer.ExecuteUntilOutput()

	if didHalt == true {
		return true, 0, 0, 0
	}

	if err != nil {
		panic(err)
	}

	GameComputer.ClearOutput()

	yPos, didHalt, err := GameComputer.ExecuteUntilOutput()

	if didHalt == true {
		return true, 0, 0, 0
	}

	if err != nil {
		panic(err)
	}

	GameComputer.ClearOutput()

	tileType, didHalt, err := GameComputer.ExecuteUntilOutput()

	if didHalt == true {
		return true, 0, 0, 0
	}

	if err != nil {
		panic(err)
	}

	GameComputer.ClearOutput()

	return false, int(xPos), int(yPos), int(tileType)
}

func InitTiles() {
	Tiles = make([][]*Tile, WorldWidth)

	for i := 0; i < WorldWidth; i++ {
		Tiles[i] = make([]*Tile, WorldHeight)

		for j := 0; j < WorldHeight; j++ {
			Tiles[i][j] = NewTile()
		}
	}
}

func PrintTiles() {
	fmt.Printf("Score: %v\n", Score)

	for y := 0; y < WorldHeight; y++ {
		for x := 0; x < WorldWidth; x++ {
			fmt.Printf("%v", Tiles[x][y].ToString())
		}

		fmt.Printf("\n")
	}

	f := bufio.NewWriter(os.Stdout)
	f.Flush()
}

func CountBlocks() int {
	var sum int = 0

	for _, tl := range Tiles {
		for _, t := range tl {
			if t.Type == Block {
				sum++
			}
		}
	}

	return sum
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func AutoMovePaddle() {
	lastX := LastBallX
	lastY := LastBallY

	ballX, ballY, err := BallPosition()

	if err != nil {
		return
	}

	LastBallX = ballX
	LastBallY = ballY

	if lastX < 0 {
		return
	}

	paddleX, _, err := PaddlePosition()

	if err != nil {
		return
	}

	movingUp := ballY > lastY
	movingRight := ballX > lastX

	var targetX int

	// if the ball is moving up, position the paddle so it is under the ball
	if movingUp == true {
		targetX = ballX
	} else { // ball is moving down, position it ahead of the moving ball
		if movingRight == true {
			targetX = ballX + 1
		} else {
			targetX = ballX - 1
		}
	}

	if targetX == paddleX {
		SetJoystickNeutral()
	} else if targetX > paddleX {
		SetJoystickRight()
	} else if targetX < paddleX {
		SetJoystickLeft()
	}
}

func SetJoystickNeutral() {
	GameComputer.SetInput(0)
}

func SetJoystickLeft() {
	GameComputer.SetInput(-1)
}

func SetJoystickRight() {
	GameComputer.SetInput(1)
}

func PaddlePosition() (int, int, error) {
	return TilePosition(HPaddle)
}

func BallPosition() (int, int, error) {
	return TilePosition(Ball)
}

func TilePosition(tt TileType) (int, int, error) {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if Tiles[x][y].Type == tt {
				return x, y, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("Could not find %v in Tiles", tt)
}
