package main

import (
	"errors"
	"fmt"
)

// () where the parameters go
// {} stores the logic of the function

func part03() {
	fmt.Println("Hello from 03_functions_and_control_structures")

	myFunc1("im calling myFunc1")
	var div0 int = intDivision0(6, 3)
	fmt.Println(div0)

	// call a multi-return function and store each return into respective var.  int1 and int2
	int1, int2 := intDivision1(11, 5)
	fmt.Printf("11 / 5 is %d Remainder %d\n", int1, int2)

	// call a multi-return function that could return an err, store each return into respective var, int3 and int4, and handle the err
	/* math refresh
	Numerator: Top number, tells how many parts you have.
	Denominator: Bottom number, tells into how many parts the whole is divided.
	Fraction: Represents a part of a whole.
	*/
	numerator := 6
	denominator := 0
	div0, remainder0, err := intDivision2(numerator, denominator)
	if err != nil {
		fmt.Println(err)
	} else if remainder0 == 0 { //only runs if parent condition is false and its condition is true
		fmt.Printf("%v / %v is %v\n", numerator, denominator, div0)
	} else { // only runs if all previous conditions were false
		fmt.Printf("%v / %v is %v Remainder %v\n", numerator, denominator, div0, remainder0)
	}

	// some conditionals
	if 3 == 3 && 4 == 4 { // the 4 == 4 only checks if 3 == 3 is true
		fmt.Println("conditions true")
	}

	if 4 == 3 || 4 == 4 { // this condition will check both values if the first one is false
		fmt.Println("conditions true")
	}

	// case statements: with expression (num)
	num := 2

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}

	div1, remainder1, err := intDivision2(numerator, denominator)

	// case statement: without expression (put the expressions into the case block)
	switch {
	case err != nil:
		fmt.Println(err)
	case remainder1 == 0:
		fmt.Printf("%v / %v is %v\n", numerator, denominator, div1)
	default:
		fmt.Printf("%v / %v is %v Remainder %v\n", numerator, denominator, div1, remainder1)

	}

	switch remainder1 {
	case 0:
		fmt.Println("the division was exact with no remainders")
	case 1, 2:
		fmt.Println("the division was close with a few remainders")
	default:
		fmt.Println("the division was not even close")

	}
}

func myFunc1(inputName string) {
	fmt.Println(inputName)

}

// function that takes in 2 ints and returns an int
func intDivision0(numerator int, denominator int) int {
	//var result int = numerator/denominator
	// return result
	// or
	return numerator / denominator

}

// multi return function that return 2 ints (int, int)
func intDivision1(numerator int, denominator int) (int, int) {
	//var result int = numerator/denominator
	return numerator / denominator, numerator % denominator

}

// multi return function that return 2 ints and an error (int, int, error) - requires importing errors package
// functions should return errors and its the job of the caller to handle the error
func intDivision2(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("error: cannot divide by 0 ")
		return -1, -1, err //Using -1 as a sentinel value indicates a special condition or error state that has a specific meaning defined by the developer.
	}
	return numerator / denominator, numerator % denominator, nil

}
