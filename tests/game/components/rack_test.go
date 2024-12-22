package components_test

import (
	"A-MATH/game/components"
	"testing"
)

const GetTotalChipErrorf = "r.GetTotalChip() function is not equal to %d"
const RackValueChipErrorf = "Rack(%d) Has Chip:%s but expected %s"

func TestRack(t *testing.T) {
	r := components.NewRack()
	result := r.IsEmpty()

	if result != true {
		t.Errorf("IsEmpty() function is false")
	}
}

func TestRackIncreaseAndTotalChip(t *testing.T) {
	r := components.NewRack()
	var chipForAdd = []string{"1", "19"}
	length := len(chipForAdd)

	for _, ch := range chipForAdd {
		r.Add(components.NewChip(ch))
	}

	if r.GetTotalChip() != length {
		t.Errorf(GetTotalChipErrorf, length)
	}

	for i, ch := range chipForAdd {
		if ch != r.Chips[i].Value {
			t.Errorf(RackValueChipErrorf, i, r.Chips[i].Value, ch)
		}
	}

}

func TestRackDecreaseAndTotalChip(t *testing.T) {
	r := components.NewRack()
	var chipForAdd = []string{"1", "19", "5", "+", "=", "+-"}
	var chipForRemove = []string{"19", "+", "="}
	var expectedChipLeft = []string{"1", "5", "+-"}
	length := len(expectedChipLeft)

	for _, ch := range chipForAdd {
		r.Add(components.NewChip(ch))
	}

	for _, ch := range chipForRemove {
		r.Remove(components.NewChip(ch))
	}

	if r.GetTotalChip() != length {
		t.Errorf(GetTotalChipErrorf, length)
	}

	for i, ch := range expectedChipLeft {
		if ch != r.Chips[i].Value {
			t.Errorf(RackValueChipErrorf, i, r.Chips[i].Value, ch)
		}
	}
}

func TestRackIsFull(t *testing.T) {
	r := components.NewRack()
	var chipForAdd = []string{"1", "19", "5", "+", "=", "+-", "8", "16"}
	length := len(chipForAdd)

	for _, ch := range chipForAdd {
		r.Add(components.NewChip(ch))
	}

	if r.GetTotalChip() != 8 {
		t.Errorf(GetTotalChipErrorf, length)
	}

	if !r.IsFull() {
		t.Errorf("r.IsFull() function is false")
	}

	for i, ch := range chipForAdd {
		if ch != r.Chips[i].Value {
			t.Errorf(RackValueChipErrorf, i, r.Chips[i].Value, ch)
		}
	}
}
