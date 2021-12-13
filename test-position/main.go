package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

var x, y int

func main() {
	fmt.Print("Input value of X: ")
	_, err := fmt.Scanln(&x)
	if err != nil {
		panic(err)
	}
	fmt.Print("Input value of Y: ")
	_, err = fmt.Scanln(&y)
	if err != nil {
		panic(err)
	}

	robotgo.Move(x, y)
}
