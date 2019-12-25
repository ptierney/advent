package day11

import (
	"advent/common"

	"fmt"
)

type Panel struct {
	Color int

	HasBeenVisited bool
}

func NewPanel() *Panel {
	p := new(Panel)

	p.Color = 0
	p.HasBeenVisited = false

	return p
}

func (p *Panel) SetColor(c int) {
	if c != 1 && c != 0 {
		panic(fmt.Sprintf("Unknown color: %v", c))
	}

	p.Color = c
	p.HasBeenVisited = true
}

var Panels [][]*Panel

var WorldDim int = 5000

var RobotX int
var RobotY int

var RobotDirection Direction

var RobotComputer *Computer

type Direction int

const (
	Up    Direction = 0
	Left  Direction = 1
	Down  Direction = 2
	Right Direction = 3
)

func InitPanels() {
	Panels = make([][]*Panel, WorldDim)

	for i := 0; i < WorldDim; i++ {
		Panels[i] = make([]*Panel, WorldDim)

		for j := 0; j < WorldDim; j++ {
			Panels[i][j] = NewPanel()
		}
	}
}

func TurnRobot(dir int) {
	if dir == 0 {
		TurnLeft()
	} else if dir == 1 {
		TurnRight()
	} else {
		panic(fmt.Sprintf("Unknown turning direction: %v", dir))
	}
}

func MoveRobot() {
	switch RobotDirection {
	case Up:
		RobotY--
	case Down:
		RobotY++
	case Left:
		RobotX--
	case Right:
		RobotX++
	}
}

func TurnLeft() {
	if RobotDirection == Right {
		RobotDirection = Up
	} else {
		RobotDirection++
	}
}

func TurnRight() {
	if RobotDirection == Up {
		RobotDirection = Right
	} else {
		RobotDirection--
	}
}

func Part1() {
	input := common.GetInput("day11/input")

	InitPanels()

	RobotX = WorldDim / 2
	RobotY = WorldDim / 2

	RobotDirection = Up

	RobotComputer = NewComputer()

	RobotComputer.LoadProgramFromInput(input)

	for {
		pv := Panels[RobotX][RobotY].Color
		RobotComputer.SetInput(int64(pv))

		color, didHalt, err := RobotComputer.ExecuteUntilOutput()

		if didHalt == true {
			break
		}

		if err != nil {
			panic(err)
		}

		RobotComputer.ClearOutput()

		Panels[RobotX][RobotY].SetColor(int(color))

		dir, didHalt, err := RobotComputer.ExecuteUntilOutput()

		if didHalt == true {
			break
		}

		if err != nil {
			panic(err)
		}

		RobotComputer.ClearOutput()

		TurnRobot(int(dir))

		MoveRobot()
	}

	numTouched := NumberPanelsTouched()

	fmt.Printf("Num Touched: %v\n", numTouched)
}

func Part2() {
	input := common.GetInput("day11/input")

	InitPanels()

	RobotX = WorldDim / 2
	RobotY = WorldDim / 2

	RobotDirection = Up

	RobotComputer = NewComputer()

	RobotComputer.LoadProgramFromInput(input)

	// Start robot on single thite panel
	Panels[RobotX][RobotY].SetColor(1)

	for {
		pv := Panels[RobotX][RobotY].Color
		RobotComputer.SetInput(int64(pv))

		color, didHalt, err := RobotComputer.ExecuteUntilOutput()

		if didHalt == true {
			break
		}

		if err != nil {
			panic(err)
		}

		RobotComputer.ClearOutput()

		Panels[RobotX][RobotY].SetColor(int(color))

		dir, didHalt, err := RobotComputer.ExecuteUntilOutput()

		if didHalt == true {
			break
		}

		if err != nil {
			panic(err)
		}

		RobotComputer.ClearOutput()

		TurnRobot(int(dir))

		MoveRobot()
	}

	//numTouched := NumberPanelsTouched()

	h := 16
	w := 100

	startX := WorldDim/2 - w/2
	startY := WorldDim/2 - h/2

	for y := startY; y < startY+h; y++ {
		for x := startX; x < startX+w; x++ {

			c := Panels[x][y].Color
			var character string

			if c == 0 {
				character = "."
			} else {
				character = "#"
			}

			fmt.Printf("%v", character)
		}

		fmt.Printf("\n")
	}
}

func NumberPanelsTouched() int {
	var sum int = 0

	for _, ps := range Panels {
		for _, p := range ps {
			if p.HasBeenVisited == true {
				sum++
			}
		}
	}

	return sum
}
