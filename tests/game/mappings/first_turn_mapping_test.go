package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"testing"
)

func TestFirstTurnMapping(t *testing.T) {
	boardForTest := components.NewBoard()

	var chips []models.ChipForPlacing
	chips = append(chips, models.NewChipForPlacing([2]int{1, 8}, components.NewChip("14")))
	chips = append(chips, models.NewChipForPlacing([2]int{2, 8}, components.NewChip("-")))
	chips = append(chips, models.NewChipForPlacing([2]int{3, 8}, components.NewChip("7")))
	chips = append(chips, models.NewChipForPlacing([2]int{4, 8}, components.NewChip("+")))
	chips = append(chips, models.NewChipForPlacing([2]int{5, 8}, components.NewChip("11")))
	chips = append(chips, models.NewChipForPlacing([2]int{6, 8}, components.NewChip("=")))
	chips = append(chips, models.NewChipForPlacing([2]int{7, 8}, components.NewChip("1")))
	chips = append(chips, models.NewChipForPlacing([2]int{8, 8}, components.NewChip("8")))

	tests := struct {
		name                   string
		board                  components.Board
		chipForPlacing         []models.ChipForPlacing
		expectedIsPlaceOnBoard [][]bool
		expectedValues         [][]string
		expectedSquareType     [][]string
	}{
		name:                   "Correct first turn chip placement",
		board:                  boardForTest,
		chipForPlacing:         chips,
		expectedIsPlaceOnBoard: [][]bool{{false, false, false, false, false, false, false, false}},
		expectedValues: [][]string{
			{
				string(constants.Fourteen),
				string(constants.Subtraction),
				string(constants.Seven),
				string(constants.Addition),
				string(constants.Eleven),
				string(constants.Equal),
				string(constants.One),
				string(constants.Eight),
			},
		},
		expectedSquareType: [][]string{
			{
				string(constants.RedSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.OrangeSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.NormalSquare),
				string(constants.CenterSquare),
			},
		},
	}

	t.Run(tests.name, func(t *testing.T) {

		results, _ := mappings.Management(tests.board, tests.chipForPlacing)

		if len(results) != 1 {
			t.Errorf("FirstTurnMapping() = len(results):%d; expected %d", len(results), 1)
		}

		testStraightMapChecking(t, results, tests.expectedIsPlaceOnBoard, tests.expectedValues, tests.expectedSquareType)
	})

}
