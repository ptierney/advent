package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput() []string {
	file, err := os.Open("input")
	//file, err := os.Open("test_input_2")

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

func maxWidthOfInput(input []string) int {
	maxLen := len(input[0])

	for _, s := range input {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	return maxLen
}

func parseInput(input []string) {
	Carts = make([]*Cart, 0)
	tempTiles := make([][]*TrackTile, len(input))
	maxWidth := len(input[0])
	maxHeight := len(input)

	for i, line := range input {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}

		tileLine := make([]*TrackTile, len(line))

		for ii, char := range line {
			tile := NewTile(char)

			tileLine[ii] = tile

			if containsCart(char) {
				MakeCart(ii, i, char)
			}
		}

		tempTiles[i] = tileLine
	}

	// Create an empty Grid
	Tiles = make([][]*TrackTile, maxWidth)

	for x := 0; x < maxWidth; x++ {
		Tiles[x] = make([]*TrackTile, maxHeight)

		for y := 0; y < maxHeight; y++ {
			Tiles[x][y] = NewTile(' ')
		}
	}

	// Populate the values

	for i, line := range tempTiles {
		for ii, t := range line {
			Tiles[ii][i] = t
		}
	}
}

type TrackType int

const (
	TrackType_empty      TrackType = iota
	TrackType_vertical   TrackType = iota
	TrackType_horizontal TrackType = iota
	// A right curve is a /
	// The cart turns right when it is traveling up, from bottom to top
	TrackType_rightCurve TrackType = iota
	// A left curve is a \
	// The curve turns left when it is traveling up, from bottom to top
	TrackType_leftCurve TrackType = iota

	TrackType_intersection TrackType = iota
)

type TurnType int

const (
	TurnType_left     TurnType = iota
	TurnType_straight TurnType = iota
	TurnType_right    TurnType = iota
)

type TrackTile struct {
	Type TrackType

	Character rune
}

type Cart struct {
	X int
	Y int

	XPrev int
	YPrev int

	Character rune

	LastTurn TurnType
}

func containsCart(inputChar rune) bool {
	switch inputChar {
	case '>', '<', 'v', '^':
		return true
	}

	return false
}

func (c *Cart) MoveStraight() {
	c.Y += (c.Y - c.YPrev)
	c.X += (c.X - c.XPrev)
}

func (c *Cart) TurnLeft() {
	startX := c.X
	startY := c.Y

	c.X += (startY - c.YPrev)
	c.Y -= (startX - c.XPrev)
}

func (c *Cart) TurnRight() {
	startX := c.X
	startY := c.Y

	c.X -= (startY - c.YPrev)
	c.Y += (startX - c.XPrev)
}

func NewCart(x int, y int, inputChar rune) *Cart {
	cart := new(Cart)

	cart.Character = inputChar

	cart.LastTurn = TurnType_right // technically false but works

	cart.X = x
	cart.Y = y

	switch inputChar {
	case '>':
		cart.XPrev = x - 1
		cart.YPrev = y
	case '<':
		cart.XPrev = x + 1
		cart.YPrev = y
	case 'v':
		cart.XPrev = x
		cart.YPrev = y - 1
	case '^':
		cart.XPrev = x
		cart.YPrev = y + 1
	default:
		log.Fatal("Unexpected cart character")
	}

	return cart
}

// Makes the cart and adds it to the carts
func MakeCart(x int, y int, inputChar rune) {
	c := NewCart(x, y, inputChar)

	Carts = append(Carts, c)
}

var Tiles [][]*TrackTile
var Carts []*Cart

func NewTile(inputChar rune) *TrackTile {
	t := new(TrackTile)
	t.Character = inputChar

	switch inputChar {
	case '-', '>', '<':
		t.Type = TrackType_horizontal
	case '|', '^', 'v':
		t.Type = TrackType_vertical
	case '\\':
		t.Type = TrackType_leftCurve
	case '/':
		t.Type = TrackType_rightCurve
	case '+':
		t.Type = TrackType_intersection
	case ' ':
		t.Type = TrackType_empty
	default:
		logStr := fmt.Sprintf("Unknown character type in input: %v", inputChar)
		log.Fatal(logStr)
	}

	// store the basic track version of the character
	switch inputChar {
	case '<', '>':
		t.Character = '-'
	case '^', 'v':
		t.Character = '|'
	}

	return t
}

func cartAtPosition(x int, y int) *Cart {
	for _, c := range Carts {
		if c.X == x && c.Y == y {
			return c
		}
	}

	return nil
}

func DisplayTrackTiles() {
	for y := 0; y < len(Tiles[0]); y++ {
		for x := 0; x < len(Tiles); x++ {
			c := cartAtPosition(x, y)
			var printChar rune

			if c != nil {
				printChar = c.Character
			} else {
				printChar = Tiles[x][y].Character
			}

			fmt.Printf("%v", string(printChar))
		}
		fmt.Printf("\n")
	}
}

func cartsIntersect() (*Cart, *Cart) {
	for i := 0; i < len(Carts)-1; i++ {
		for j := i + 1; j < len(Carts); j++ {
			c1 := Carts[i]
			c2 := Carts[j]
			if c1.X == c2.X && c1.Y == c2.Y {
				return c1, c2
			}
		}
	}

	return nil, nil
}

func removeCarts(c1 *Cart, c2 *Cart) {
	removeCart(c1)
	removeCart(c2)
}

func removeCart(cart *Cart) {
	for i, c := range Carts {
		if c == cart {
			Carts = append(Carts[:i], Carts[i+1:]...)
			break
		}
	}
}

func tickCart(cart *Cart) {
	tile := Tiles[cart.X][cart.Y]

	startX := cart.X
	startY := cart.Y

	switch tile.Type {
	case TrackType_empty:
		log.Fatalf("Cart is on empty track at <%v, %v>\n", cart.X, cart.Y)
	case TrackType_vertical, TrackType_horizontal:
		cart.MoveStraight()
	case TrackType_rightCurve:
		if cart.YPrev == cart.Y { // moving horizontally
			cart.TurnLeft()
		} else {
			cart.TurnRight()
		}
	case TrackType_leftCurve:
		if cart.YPrev == cart.Y { // moving horizontally
			cart.TurnRight()
		} else {
			cart.TurnLeft()
		}
	case TrackType_intersection:
		switch cart.LastTurn {
		case TurnType_right:
			cart.TurnLeft()
			cart.LastTurn = TurnType_left
		case TurnType_left:
			cart.MoveStraight()
			cart.LastTurn = TurnType_straight
		case TurnType_straight:
			cart.TurnRight()
			cart.LastTurn = TurnType_right
		}
	}

	cart.XPrev = startX
	cart.YPrev = startY
}

func tickCarts() {
	for _, c := range Carts {
		tickCart(c)
	}
}

func getCartsForTick() []*Cart {
	carts := make([]*Cart, 0)

	for x := 0; x < len(Tiles); x++ {
		for y := 0; y < len(Tiles[0]); y++ {
			c := cartAtPosition(x, y)

			if c == nil {
				continue
			}

			carts = append(carts, c)
		}
	}

	return carts
}

func main() {
	parseInput(getInput())

	for {

		currentCarts := getCartsForTick()

		for _, c := range currentCarts {
			tickCart(c)

			cart1, cart2 := cartsIntersect()

			if cart1 != nil {
				fmt.Printf("Crash at: <%v,%v>\n", cart1.X, cart1.Y)

			}

			removeCarts(cart1, cart2)
		}

		if len(Carts) > 1 {
			continue
		}

		fmt.Printf("Final Cart at: <%v,%v>\n", Carts[0].X, Carts[0].Y)

		break
	}
}
