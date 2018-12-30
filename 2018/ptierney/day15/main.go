package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

func getInput(fn string) []string {
	file, err := os.Open(fn)

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

func containsCharacter(r rune) bool {
	return r == 'E' || r == 'G'
}

func parseInput(input []string) {
	Characters = make([]*Character, 0)
	AliveCharacters = list.New()
	tempTiles := make([][]*Tile, len(input))
	maxWidth := len(input[0])
	maxHeight := len(input)

	for i, line := range input {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}

		tileLine := make([]*Tile, len(line))

		for ii, char := range line {
			tile := NewTile(char)

			tileLine[ii] = tile

			if containsCharacter(char) {
				MakeCharacter(ii, i, char)
			}
		}

		tempTiles[i] = tileLine
	}

	// Create an empty Grid
	Tiles = make([][]*Tile, maxWidth)

	for x := 0; x < maxWidth; x++ {
		Tiles[x] = make([]*Tile, maxHeight)

		// TODO: remove this
		for y := 0; y < maxHeight; y++ {
			Tiles[x][y] = NewTile('#')
		}
	}

	// Populate the values

	for i, line := range tempTiles {
		for ii, t := range line {
			Tiles[ii][i] = t
			Tiles[ii][i].X = ii
			Tiles[ii][i].Y = i
		}
	}

	// lace Tiles

	for y := 0; y < len(Tiles[0]); y++ {
		for x := 0; x < len(Tiles); x++ {
			if x > 0 {
				Tiles[x][y].WestTile = Tiles[x-1][y]
			}

			if x < len(Tiles)-1 {
				Tiles[x][y].EastTile = Tiles[x+1][y]
			}

			if y > 0 {
				Tiles[x][y].NorthTile = Tiles[x][y-1]
			}

			if y < len(Tiles[0])-1 {
				Tiles[x][y].SouthTile = Tiles[x][y+1]
			}
		}
	}

	// copy characters

	for _, c := range Characters {
		AliveCharacters.PushBack(c)
	}
}

type TileType int

const (
	TileType_open TileType = iota
	TileType_wall TileType = iota
)

type CharacterType int

const (
	CharacterType_goblin CharacterType = iota
	CharacterType_elf    CharacterType = iota
)

type Tile struct {
	Type TileType

	X int
	Y int

	NorthTile *Tile
	SouthTile *Tile
	EastTile  *Tile
	WestTile  *Tile

	Glyph rune // what to display

	Distance int
}

func NewTile(inputRune rune) *Tile {
	tile := new(Tile)

	switch inputRune {
	case '#':
		tile.Type = TileType_wall
		tile.Glyph = '#'
	case '.', 'G', 'E':
		tile.Type = TileType_open
		tile.Glyph = '.'
	default:
		logStr := fmt.Sprintf("Unknown character type in input: %v", string(inputRune))
		log.Fatal(logStr)
	}

	tile.NorthTile = nil
	tile.SouthTile = nil
	tile.EastTile = nil
	tile.WestTile = nil

	tile.Distance = -1

	return tile
}

func (t *Tile) HashKey() string {
	return fmt.Sprintf("%v,%v", t.X, t.Y)
}

func (t *Tile) CharacterOnTile() *Character {
	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Character)

		if t.X == c.X && t.Y == c.Y {
			return c
		}
	}

	return nil
}

// open is not a cave or occupied
func (t *Tile) IsOpen() bool {
	if t.Type != TileType_open {
		return false
	}

	return !t.AliveCharacterIsOnTile()
}

func (t *Tile) AliveCharacterIsOnTile() bool {
	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Character)

		if t.X == c.X && t.Y == c.Y {
			return true
		}
	}

	return false
}

func (t *Tile) SurroundingTiles() []*Tile {
	tiles := make([]*Tile, 4)

	tiles[0] = t.NorthTile
	tiles[1] = t.SouthTile
	tiles[2] = t.EastTile
	tiles[3] = t.WestTile

	return tiles
}

func (t *Tile) SurroundingOpenTiles() *list.List {
	tiles := list.New()

	surrounding := t.SurroundingTiles()

	for _, tile := range surrounding {
		if tile.IsOpen() {
			tiles.PushBack(tile)
		}
	}

	return tiles
}

func (t *Tile) TileIsInTileSlice(ts []*Tile) bool {
	for _, tile := range ts {
		if t.HashKey() == tile.HashKey() {
			return true
		}
	}

	return false
}

func (t *Tile) ShortestDistanceToTile(otherTile *Tile) int {
	// returns the number of steps to the other tile, along the/a shortest route

	return 0
}

func ResetTileDistances() {
	for _, t1 := range Tiles {
		for _, t2 := range t1 {
			t2.Distance = -1
		}
	}
}

func FirstTileInReadingOrder(tiles []*Tile) *Tile {
	// find smallest y

	yMin := tiles[0].Y

	for _, t := range tiles {
		if t.Y < yMin {
			yMin = t.Y
		}
	}

	tilesInY := make([]*Tile, 0)

	for _, t := range tiles {
		if t.Y == yMin {
			tilesInY = append(tilesInY, t)
		}
	}

	// then of the Y line, select the tile with
	// the smallest X

	xMin := tilesInY[0].X
	minIndex := 0

	for i, t := range tilesInY {
		if t.X < xMin {
			xMin = t.X
			minIndex = i
		}
	}

	return tilesInY[minIndex]
}

// A character is either an elf or a goblin
type Character struct {
	Type CharacterType

	X int
	Y int

	HitPoints   int
	AttackPower int

	Glyph rune
}

func NewCharacter(x int, y int, inputChar rune) *Character {
	character := new(Character)

	character.X = x
	character.Y = y

	character.HitPoints = 200
	character.AttackPower = 3

	character.Glyph = inputChar

	switch inputChar {
	case 'G':
		character.Type = CharacterType_goblin
	case 'E':
		character.Type = CharacterType_elf
	default:
		log.Fatal("Unexpected character rune")
	}

	return character
}

func (c *Character) IsAlive() bool {
	return c.HitPoints > 0
}

func (c *Character) MoveToTile(t *Tile) {
	c.X = t.X
	c.Y = t.Y
}

// implies that they are alive
func (c *Character) AllEnemies() []*Character {
	if c.Type == CharacterType_elf {
		return AllAliveCharactersOfType(CharacterType_goblin)
	} else if c.Type == CharacterType_goblin {
		return AllAliveCharactersOfType(CharacterType_elf)
	} else {
		log.Fatal("Unknown character type")
	}

	return make([]*Character, 0)
}

// all found surrending Tiles
func (c *Character) SurroundingTiles() []*Tile {
	tiles := make([]*Tile, 4)

	thisTile := Tiles[c.X][c.Y]

	tiles[0] = thisTile.NorthTile
	tiles[1] = thisTile.SouthTile
	tiles[2] = thisTile.EastTile
	tiles[3] = thisTile.WestTile

	return tiles
}

// This modifies the distances for all tiles to be the distances
func (c *Character) SetDistancesForTiles() {
	ResetTileDistances()

	currentTile := c.CurrentTile()

	tileStack := list.New()
	reachableTiles := list.New()
	evaluatedTiles := make(map[string]*Tile)
	evaluatedTiles[currentTile.HashKey()] = currentTile

	sot := currentTile.SurroundingOpenTiles()

	for e := sot.Front(); e != nil; e = e.Next() {
		e.Value.(*Tile).Distance = 1
	}

	tileStack.PushBackList(currentTile.SurroundingOpenTiles())

	for tileStack.Len() > 0 {
		// evaluate all
		tileElem := tileStack.Front()
		tile := tileElem.Value.(*Tile)
		tileStack.Remove(tileElem)
		evaluatedTiles[tile.HashKey()] = tile

		reachableTiles.PushBack(tile)

		nextTiles := tile.SurroundingOpenTiles()

		for e := nextTiles.Front(); e != nil; e = e.Next() {
			nt := e.Value.(*Tile)

			if _, contains := evaluatedTiles[nt.HashKey()]; contains {
				continue
			}

			nt.Distance = tile.Distance + 1

			tileStack.PushBack(nt)
		}
	}
}

// Reachable tiles by definition have to be open
func (c *Character) AllReachableTiles() []*Tile {
	rtSlice := make([]*Tile, 0)

	for _, t1 := range Tiles {
		for _, t2 := range t1 {
			if t2.Distance > 0 {
				rtSlice = append(rtSlice, t2)
			}
		}
	}

	return rtSlice
}

// the tiles on both sides and up and down, not occupied or
func (c *Character) SurroundingTilesInRange() []*Tile {
	inRangeTiles := make([]*Tile, 0)

	surrounding := c.SurroundingTiles()

	for _, t := range surrounding {
		if t.IsOpen() {
			inRangeTiles = append(inRangeTiles, t)
		}
	}

	return inRangeTiles
}

// every single tile in range to an enemy
func (c *Character) AllTilesInRange() []*Tile {
	enemies := c.AllEnemies()

	inRangeMap := make(map[string]*Tile)

	for _, enemy := range enemies {
		tilesInRange := enemy.SurroundingTilesInRange()

		for _, t := range tilesInRange {
			inRangeMap[t.HashKey()] = t
		}
	}

	allInRange := make([]*Tile, 0)

	for _, value := range inRangeMap {
		allInRange = append(allInRange, value)
	}

	return allInRange
}

func (c *Character) GetMoveTargetTile() *Tile {
	// find all the tiles in range (to an enemy)
	allInRange := c.AllTilesInRange()

	c.SetDistancesForTiles()

	// find the tiles that are reachable
	allReachable := c.AllReachableTiles()

	reachableInRange := make([]*Tile, 0)

	for _, inRangeTile := range allInRange {
		if inRangeTile.TileIsInTileSlice(allReachable) {
			reachableInRange = append(reachableInRange, inRangeTile)
		}
	}

	// find the nearest set of reachable tiles
	// choose the nearest tile in reading order

	if len(reachableInRange) <= 0 {
		return nil
	}

	closestReachable := make([]*Tile, 1)
	closestReachable[0] = reachableInRange[0]
	var closestDist = closestReachable[0].Distance

	for _, t := range reachableInRange {
		if t.Distance == closestDist {
			closestReachable = append(closestReachable, t)
		} else if t.Distance < closestDist {
			closestReachable = make([]*Tile, 1)
			closestReachable[1] = t
			closestDist = t.Distance
		}
	}

	return FirstTileInReadingOrder(closestReachable)
}

func (c *Character) Tick() {
	// move is possible
	currentTile := c.CurrentTile()

	targetTile := c.GetMoveTargetTile()

	if targetTile != nil {

		// move to the target tile

		// choose the boundary point in reading order

		surroundingOpen := currentTile.SurroundingOpenTiles()

		bestSurrounding := make([]*Tile, 1)
		bestSurrounding[0] = surroundingOpen.Front().Value.(*Tile)

		bestDist := bestSurrounding[0].ShortestDistanceToTile(targetTile)

		for e := surroundingOpen.Front().Next(); e != nil; e = e.Next() {
			t := e.Value.(*Tile)
			dist := t.ShortestDistanceToTile(targetTile)

			if dist == bestDist {
				bestSurrounding = append(bestSurrounding, t)
			} else if dist < bestDist {
				bestSurrounding := make([]*Tile, 1)
				bestSurrounding[0] = t
				bestDist = dist
			}
		}

		tileToMoveTo := FirstTileInReadingOrder(bestSurrounding)

		c.MoveToTile(tileToMoveTo)
	}

	// calculate the distance to the nearest tile
	// for all four boundary points
	// #######
	// #4E212#
	// #32101#
	// #432G2#
	// #######

	enemyTiles := c.SurroundingTilesWithEnemies()

	if len(enemyTiles) == 0 {
		return // turn is over
	}

	// select one enemy in range by reading order
	attackTile := FirstTileInReadingOrder(enemyTiles)

	c.Attack(attackTile.CharacterOnTile())
}

func (c *Character) Attack(otherChar *Character) {
	otherChar.HitPoints -= c.AttackPower
}

func (c *Character) CurrentTile() *Tile {
	return Tiles[c.X][c.Y]
}

func (c *Character) tileContainsEnemy(t *Tile) bool {
	charOnTile := t.CharacterOnTile()

	if char == nil {
		return false
	}

	if charOnTile.Type == c.Type {
		return false
	}

	return true
}

func (c *Character) SurroundingTilesWithEnemies() []*Tile {
	tiles := make([]*Tile, 0)

	st := c.CurrentTile().SurroundingTiles()

	for _, tile := range st {
		if c.tileContainsEnemy(tile) {
			tiles = append(tiles, tile)
		}
	}

	return tiles
}

func MakeCharacter(x int, y int, inputChar rune) {
	c := NewCharacter(x, y, inputChar)

	Characters = append(Characters, c)
}

func AllAliveCharactersOfType(t CharacterType) []*Character {
	chars := make([]*Character, 0)

	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		if e.Value.(*Character).Type == t {
			chars = append(chars, e.Value.(*Character))
		}
	}

	return chars
}

var Tiles [][]*Tile
var Characters []*Character
var AliveCharacters *list.List

var RoundCounter int

func characterAtPosition(x int, y int) *Character {
	for _, c := range Characters {
		if c.IsAlive() == false {
			continue
		}

		if c.X == x && c.Y == y {
			return c
		}
	}

	return nil
}

func DisplayTiles() {
	for y := 0; y < len(Tiles[0]); y++ {
		for x := 0; x < len(Tiles); x++ {
			c := characterAtPosition(x, y)
			var printChar rune

			if c != nil {
				printChar = c.Glyph
			} else {
				printChar = Tiles[x][y].Glyph
			}

			fmt.Printf("%v", string(printChar))
		}
		fmt.Printf("\n")
	}
}

func OrderedCharactersForRound() *list.List {
	charList := list.New()

	for y := 0; y < len(Tiles[0]); y++ {
		for x := 0; x < len(Tiles); x++ {
			c := characterAtPosition(x, y)

			if c != nil {
				charList.PushBack(c)
			}
		}
	}

	return charList
}

func OneSideHasWon() bool {
	if AliveCharacters.Len() == 0 {
		log.Fatal("No side has won, everyone is dead")
	}

	assumption := AliveCharacters.Front().Value.(*Character).Type

	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		if e.Value.(*Character).Type != assumption {
			return false
		}
	}

	return true
}

func findAndRemoveDeadCharacters() *Character {
	// loop through the alive characters, finding a dead one. Return it after
	// removing it

	return nil
}

// input is an in-progress list search. It remove the character from the list
// if it is present (it may have already passed the character), and then
// returns the new "Next" element, or nil if at the end of the list
// removes the character from the in pogress list
func removeCharacterFromInProgressList(c *Character, e *list.Element) *list.Element {
	//

	return nil
}

func battleUntilComplete() {

	RoundCounter = 0

	return

	for {
		charactersForRound := OrderedCharactersForRound()

		for e := charactersForRound.Front(); e != nil; {
			c := e.Value.(*Character)

			c.Tick()

			deadChar := findAndRemoveDeadCharacters()

			if deadChar != nil {
				e = removeCharacterFromInProgressList(deadChar, e)
			}

			if OneSideHasWon() {
				break
			}
		}

		// only count FULL rounds
		RoundCounter++

		if OneSideHasWon() {
			break
		}
	}
}

func scoreForAliveCharacters() int {
	scoreSum := 0

	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		scoreSum += e.Value.(*Character).HitPoints
	}

	return scoreSum
}

func GetBattleOutcome(inputFileName string) (int, int, int) {
	parseInput(getInput(inputFileName))

	DisplayTiles()

	battleUntilComplete()

	return RoundCounter, scoreForAliveCharacters(), RoundCounter * scoreForAliveCharacters()
}

func main() {
	r, hp, outcome := GetBattleOutcome("input")

	fmt.Printf("Outcome: %v * %v = %v", r, hp, outcome)
}
