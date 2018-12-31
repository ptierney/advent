package main

import (
	"testing"
)

func printFailure(t *testing.T, expected int, outcome int) {
	t.Fatalf("Expected %v but received %v", expected, outcome)
}

func testBattleFile(t *testing.T, filename string, expected int) {
	_, _, outcome := GetBattleOutcome(filename)

	if outcome != expected {
		printFailure(t, expected, outcome)
	}
}

func setupBasicTests() {
	parseInput(getInput("test_input_1"))
}

func TestAliveCharactersCount(t *testing.T) {
	setupBasicTests()

	if AliveCharacters.Len() != 6 {
		printFailure(t, 6, AliveCharacters.Len())
	}
}

func TestAllAliveCharactersOfType(t *testing.T) {
	setupBasicTests()

	numElf := len(AllAliveCharactersOfType(CharacterType_elf))
	numGob := len(AllAliveCharactersOfType(CharacterType_goblin))

	if numElf != 2 {
		printFailure(t, 2, numElf)
	}

	if numGob != 4 {
		printFailure(t, 4, numGob)
	}
}

func TestCharacterAtPosition(t *testing.T) {
	setupBasicTests()

	char := CharacterAtPosition(2, 1)

	if char.Type != CharacterType_goblin {
		t.Fatalf("Expected to find Goblit")
	}

	char = CharacterAtPosition(4, 2)

	if char.Type != CharacterType_elf {
		t.Fatalf("Expected to find Elf")
	}
}

func TestGetTargetTile(t *testing.T) {
	setupBasicTests()

	goblin := CharacterAtPosition(2, 1)

	goblinTarget := goblin.GetMoveTargetTile()

	if goblinTarget.X != 4 || goblinTarget.Y != 1 {
		t.Fatalf("Goblin target incorrect, is: %v, %v",
			goblinTarget.X, goblinTarget.Y)
	}

	goblin = CharacterAtPosition(3, 4)

	goblinTarget = goblin.GetMoveTargetTile()

	if goblinTarget.X != 3 || goblinTarget.Y != 2 {
		t.Fatalf("2nd Goblin target incorrect")
	}

	elf := CharacterAtPosition(5, 4)

	elfTarget := elf.GetMoveTargetTile()

	if elfTarget.X != 3 || elfTarget.Y != 5 {
		t.Fatalf("Elf target incorrect")
	}
}

func TestGetTargetTileNull(t *testing.T) {
	setupBasicTests()

	goblin := CharacterAtPosition(5, 3)

	goblinTarget := goblin.GetMoveTargetTile()

	if goblinTarget != nil {
		t.Fatalf("Expected a null target")
	}
}

func TestAllTilesInRange(t *testing.T) {
	setupBasicTests()

	goblin := CharacterAtPosition(2, 1)

	tilesInRange := goblin.AllTilesInRange()

	if len(tilesInRange) != 3 {
		printFailure(t, 3, len(tilesInRange))
	}
}

func TestSurroundingTilesWithEnemies(t *testing.T) {
	setupBasicTests()

	goblin := CharacterAtPosition(2, 1)

	enemies := goblin.SurroundingTilesWithEnemies()

	if len(enemies) != 0 {
		printFailure(t, 0, len(enemies))
	}

	elf := CharacterAtPosition(4, 2)

	enemies = elf.SurroundingTilesWithEnemies()

	if len(enemies) != 1 {
		printFailure(t, 1, len(enemies))
	}
}

func TestShortestDistanceToTile(t *testing.T) {
	setupBasicTests()

	t1 := Tiles[3][1]

	t2 := Tiles[4][1]

	dist := t1.ShortestDistanceToTile(t2)

	if dist != 1 {
		printFailure(t, 1, dist)
	}

	t1 = Tiles[2][2]

	t2 = Tiles[5][5]

	dist = t1.ShortestDistanceToTile(t2)

	if dist != 8 {
		printFailure(t, 8, dist)
	}

	t1 = Tiles[3][2]

	t2 = Tiles[4][1]

	dist = t1.ShortestDistanceToTile(t2)

	if dist != 2 {
		printFailure(t, 2, dist)
	}
}

func TestShortestDistanceToTileFailure(t *testing.T) {
	parseInput(getInput("test_input_1.1"))

	t1 := Tiles[3][2]

	t2 := Tiles[4][1]

	// this point is unreachable
	dist := t1.ShortestDistanceToTile(t2)

	if dist != 2147483647 {
		printFailure(t, 2147483647, dist)
	}
}

func TestTick(t *testing.T) {
	setupBasicTests()

	goblin := CharacterAtPosition(2, 1)

	goblin.Tick()

	if goblin.X != 3 || goblin.Y != 1 {
		t.Fatalf("Goblin tick in wrong position")
	}

}

func TestFirstTileInReadingOrder(t *testing.T) {
	setupBasicTests()

	tile := Tiles[2][1]

	surrounding := tile.SurroundingTiles()

	first := FirstTileInReadingOrder(surrounding)

	if first.X != 2 || first.Y != 0 {
		t.Fatalf("Incorrect reading order, got: %v, %v", first.X, first.Y)
	}
}

func TestOrderedCharactersForRound(t *testing.T) {
	setupBasicTests()

	ordered := OrderedCharactersForRound()

	first := ordered.Front().Value.(*Character)

	if first.X != 2 || first.Y != 1 {
		t.Fatalf("Incorrect first character, got: %v, %v", first.X, first.Y)
	}
}

func TestInput1(t *testing.T) {
	testBattleFile(t, "test_input_1", 27730)
}

func TestInput2(t *testing.T) {
	testBattleFile(t, "test_input_2", 36334)
}

func TestInput3(t *testing.T) {
	testBattleFile(t, "test_input_3", 39514)
}

// func TestInput4(t *testing.T) {
// 	testBattleFile(t, "test_input_4", 27755)
// }

// func TestInput5(t *testing.T) {
// 	testBattleFile(t, "test_input_5", 28944)
// }

// func TestInput6(t *testing.T) {
// 	testBattleFile(t, "test_input_6", 18740)
// }
