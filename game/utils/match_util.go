package utils

import (
	"A-MATH/game/constants"
	"math/rand"
)

func RandomString(n int) string {
	letter := constants.RandomLetter
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
