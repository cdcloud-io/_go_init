package main

import (
	"fmt"
)

/*
Notes:
define a variable: var <name> <type>
define a const: const <name> <type> = <value>  // const variables cannot change, but they have to be initialized with a value.

*/
// outside of a functions, variables are scoped to the package
var MyString1 string = "I am accessible outside of package main"
var myString2 string = "I am NOT accesible outside of package main"

func part10() {
	fmt.Println("--------------------------------")
	fmt.Println("part10()")
	fmt.Println("--------------------------------")

	fmt.Println("MyString1: ", MyString1)
	fmt.Println("myString2: ", myString2)

	// declared only will make GO initialize to defaults
	var var1 bool    // initializes to false
	var var2 int     // initializes to 0
	var var3 float64 // initializes to 0
	var var4 string  // initializes to ""

	fmt.Println("var1 declared only -> default initial value: ", var1)
	fmt.Println("var2 declared only -> default initial value: ", var2)
	fmt.Println("var3 declared only -> default initial value: ", var3)
	fmt.Println("var4 declared only -> default initial value: ", var4)

	// set values to declared variable
	var1 = true
	var2 = 1
	// var2 = 1.1 error, not type float
	var3 = 1.01
	var4 = "howdy"

	fmt.Println("var1 = true: ", var1)
	fmt.Println("var2 = 1: ", var2)
	fmt.Println("var3 = 1.01: ", var3)
	fmt.Println("var4 = \"howdy\": ", var4) // escape the "" to make it print

	// multiple variables on single line.
	var v1, v2 = 1, 4        // infered int
	var v3, v4 uint32 = 2, 4 // explicit uint32
	var v5 = "hello"

	// within a function you can drop the var keyword
	v6 := "v6"

	// variables are MUTABLE so they can be changed as long as the DATATYPE is the same.
	// v5 = 1 // not allowed. type string
	v5 = "goodbye" // ok

	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("v3: ", v3)
	fmt.Println("v4: ", v4)
	fmt.Println("v5: ", v5)
	fmt.Println("v6: ", v6)

}
