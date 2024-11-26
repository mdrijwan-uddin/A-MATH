package constants

type code int
type messages string

const (
	BadRequest code = 400

	InvalidInputForChip  messages = "Invalid input for chip value"
	InvalidInputForBoard messages = "Invalid input for board position"

	InvalidChipPlacement messages = "Invalid chips placing"
)
