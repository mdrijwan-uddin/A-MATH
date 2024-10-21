package main

import (
	"A-MATH/model/game"
	"fmt"
)

// type MyError struct{}

// func (m *MyError) Error() string {
// 	return "boom"
// }

// func SayHello() (string, error) {
// 	return "", &MyError{}
// }

func main() {
	c := game.NewChip("x")
	d := game.NewChip("/")
	fmt.Println(c == d)

	a := game.NewChip("1")
	fmt.Println(a)

	// s, err := SayHello()
	// if err != nil {
	// 	fmt.Println("unexpected error: err:", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("The string:", s)

}
