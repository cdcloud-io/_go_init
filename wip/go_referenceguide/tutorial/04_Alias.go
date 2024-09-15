package main

import (
	"fmt"
)

// abbreaviate a function call ie fmt.Println()

var pl = fmt.Println // ok
// var pl = fmt.Println() // error because we are assigning the return of Println() to pl which has no args

func part04() {
	fmt.Println("--------------------------------")
	fmt.Println("part04()")
	fmt.Println("--------------------------------")

	pl("this was printed using abreviation var pl = fmt.Println")
}
