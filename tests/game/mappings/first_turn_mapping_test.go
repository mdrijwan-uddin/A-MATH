package mappings_test

import (
	"A-MATH/game/components"
	"A-MATH/game/constants"
	"A-MATH/game/mappings"
	"A-MATH/game/models"
	"log"
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

	var values []string
	values = append(values, string(constants.Fourteen))
	values = append(values, string(constants.Subtraction))
	values = append(values, string(constants.Seven))
	values = append(values, string(constants.Addition))
	values = append(values, string(constants.Eleven))
	values = append(values, string(constants.Equal))
	values = append(values, string(constants.One))
	values = append(values, string(constants.Eight))

	tests := []struct {
		name           string
		board          components.Board
		chipForPlacing []models.ChipForPlacing
		expectedValues []string
	}{
		{
			name:           "Correct chip placement",
			board:          boardForTest,
			chipForPlacing: chips,
			expectedValues: values,
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {

		results := mappings.FirstTurnMapping(tests[0].board, tests[0].chipForPlacing)

		if len(results) != 1 {
			t.Errorf("FirstTurnMapping() = len(results):%d; expected %d", len(results), 1)
		}

		for i, result := range results {
			log.Println(result[0].SquareType)

			if result[0].ChipForCalculating.Value != tests[0].expectedValues[i] {
				t.Errorf("FirstTurnMapping() = value[%d]:%v; expected %v", i, result[0].ChipForCalculating.Value, tests[0].expectedValues[i])
			}

			if result[0].IsPlacedOnBoard != false {
				t.Errorf("FirstTurnMapping() = IsPlacedOnBoard[%d]:%v; expected %v", i, result[0].IsPlacedOnBoard, false)
			}
		}
	})

}
