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

// type PathTile struct {
// 	ThisTile     Tile
// 	PreviousTile Tile
// }

// func NewPathTile(thisTile Tile, previousTile Tile) *PathTile {
// 	pTile := new(PathTile)

// 	pTile.ThisTile = thisTile
// 	pTile.PreviousTile = previousTile

// 	return pTile
// }

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
	tiles := make([]*Tile, 0)

	if t.NorthTile != nil {
		tiles = append(tiles, t.NorthTile)
	}

	if t.SouthTile != nil {
		tiles = append(tiles, t.SouthTile)
	}

	if t.EastTile != nil {
		tiles = append(tiles, t.EastTile)
	}

	if t.WestTile != nil {
		tiles = append(tiles, t.WestTile)
	}

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
	if otherTile == nil {
		log.Fatal("Cannot compute distance with nil")
	}

	previousTileMap := make(map[string]*Tile)

	tileStack := list.New()

	tileStack.PushBack(t)

	for tileStack.Len() > 0 {
		tileElem := tileStack.Front()
		tile := tileElem.Value.(*Tile)
		tileStack.Remove(tileElem)

		nextTiles := tile.SurroundingOpenTiles()

		for e := nextTiles.Front(); e != nil; e = e.Next() {
			nt := e.Value.(*Tile)

			if _, contains := previousTileMap[nt.HashKey()]; contains {
				continue
			}

			previousTileMap[nt.HashKey()] = tile

			tileStack.PushBack(nt)
		}
	}

	path := list.New()

	current := otherTile

	for {
		if current.HashKey() == t.HashKey() {
			break
		}

		path.PushFront(current)

		prev, contains := previousTileMap[current.HashKey()]

		if contains {
			current = prev
		} else {
			// in this case there is no path between the two tiles
			return 2147483647
		}
	}

	return path.Len()
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

	character.Glyph = inputChar

	switch inputChar {
	case 'G':
		character.Type = CharacterType_goblin
		character.AttackPower = 3
	case 'E':
		character.Type = CharacterType_elf
		character.AttackPower = ElfAttackPower
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

	return thisTile.SurroundingTiles()

	return tiles
}

// This modifies the distances for all tiles to be the distances
func (c *Character) SetDistancesForTiles() {
	ResetTileDistances()

	currentTile := c.CurrentTile()

	tileStack := list.New()
	evaluatedTiles := make(map[string]*Tile)

	tileStack.PushBack(currentTile)
	currentTile.Distance = 0

	for tileStack.Len() > 0 {

		//fmt.Printf("TileStack length: %v\n", tileStack.Len())
		//fmt.Printf("EvaluatedTiles length: %v\n", len(evaluatedTiles))

		// evaluate all
		tileElem := tileStack.Front()
		tile := tileElem.Value.(*Tile)
		tileStack.Remove(tileElem)

		nextTiles := tile.SurroundingOpenTiles()

		for e := nextTiles.Front(); e != nil; e = e.Next() {
			nt := e.Value.(*Tile)

			if _, contains := evaluatedTiles[nt.HashKey()]; contains {
				continue
			}

			nt.Distance = tile.Distance + 1

			tileStack.PushBack(nt)

			evaluatedTiles[nt.HashKey()] = nt
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
			closestReachable[0] = t
			closestDist = t.Distance
		}
	}

	return FirstTileInReadingOrder(closestReachable)
}

func (c *Character) IsInRangeToAttack() bool {
	return len(c.SurroundingTilesWithEnemies()) > 0
}

func (c *Character) Tick() {
	// move is possible
	currentTile := c.CurrentTile()

	var targetTile *Tile = nil

	if c.IsInRangeToAttack() == false {
		targetTile = c.GetMoveTargetTile()
	}

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
				bestSurrounding = make([]*Tile, 1)
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

	// select the enemy with the lowest hit points
	lowestScore := enemyTiles[0].CharacterOnTile().HitPoints

	lowestEnemyTiles := make([]*Tile, 1)
	lowestEnemyTiles[0] = enemyTiles[0]

	for _, enemyTile := range enemyTiles {
		hp := enemyTile.CharacterOnTile().HitPoints

		if hp == lowestScore {
			lowestEnemyTiles = append(lowestEnemyTiles, enemyTile)
		} else if hp < lowestScore {
			lowestEnemyTiles = make([]*Tile, 1)
			lowestEnemyTiles[0] = enemyTile
			lowestScore = hp
		}
	}

	// select one enemy in range by reading order
	attackTile := FirstTileInReadingOrder(lowestEnemyTiles)

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

	if charOnTile == nil {
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

func CharacterAtPosition(x int, y int) *Character {
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
	fmt.Printf("Round: %v\n", RoundCounter)

	return

	for y := 0; y < len(Tiles[0]); y++ {
		for x := 0; x < len(Tiles); x++ {
			c := CharacterAtPosition(x, y)
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
			c := CharacterAtPosition(x, y)

			if c != nil {
				charList.PushBack(c)
			}
		}
	}

	return charList
}

var OneSideHasOneFlag bool = false

func OneSideHasWon() bool {
	//	return OneSideHasOneFlag

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
	var deadCharacterElem *list.Element = nil

	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Character)

		if c.IsAlive() == false {
			deadCharacterElem = e
			break
		}
	}

	if deadCharacterElem == nil {
		return nil
	}

	deadCharacter := deadCharacterElem.Value.(*Character)

	AliveCharacters.Remove(deadCharacterElem)

	return deadCharacter
}

// input is an in-progress list search. It remove the character from the list
// if it is present (it may have already passed the character), and then
// returns the new "Next" element, or nil if at the end of the list
// removes the character from the in pogress list
// func removeCharacterFromInProgressList(c *Character, inProgress *list.Element) *list.Element {
// 	var foundElement *list.Element = nil

// 	for e := inProgress; e != nil; e = e.Next() {
// 		thisChar := e.Value.(*Character)

// 		if thisChar == c {
// 			foundElement = e
// 			break
// 		}
// 	}

// 	if foundElement != nil {

// 	}

// 	return nil
// }

func battleUntilComplete() {
	// Display the initial condition
	//DisplayTiles()

	for {

		charactersForRound := OrderedCharactersForRound()

		if OneSideHasWon() {
			break
		}

		RoundCounter++

		roundDeadChars := make(map[*Character]bool)

		for e := charactersForRound.Front(); e != nil; e = e.Next() {
			c := e.Value.(*Character)

			if _, contains := roundDeadChars[c]; contains {
				continue
			}

			c.Tick()

			deadChar := findAndRemoveDeadCharacters()

			if deadChar != nil {
				roundDeadChars[deadChar] = true

				if deadChar.Type == CharacterType_elf {
					AnElfHasDied = true
				}
			}

			if OneSideHasWon() {
				// only count full rounds
				if e.Next() != nil {
					RoundCounter--
				}

				break
			}
		}

		//DisplayTiles()
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
	RoundCounter = 0
	AnElfHasDied = false

	parseInput(getInput(inputFileName))

	//DisplayTiles()

	battleUntilComplete()

	return RoundCounter, scoreForAliveCharacters(), RoundCounter * scoreForAliveCharacters()
}

func PrintAliveCharacterScores() {
	for e := AliveCharacters.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Character)

		fmt.Printf("%v at %v : %v\n", string(c.Glyph), c.CurrentTile().HashKey(), c.HitPoints)
	}
}

var ElfAttackPower int = 3
var AnElfHasDied bool = false

func main() {
	for {
		ElfAttackPower++
		r, hp, outcome := GetBattleOutcome("input")

		if AnElfHasDied == false {
			fmt.Printf("Outcome: %v * %v = %v\n", r, hp, outcome)
			break
		} else {
			fmt.Printf("An elf died at attack power %v\n", ElfAttackPower)
		}
	}
}
