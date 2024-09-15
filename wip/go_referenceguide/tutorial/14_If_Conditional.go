package main

import (
	"fmt"
	"strconv"
)

/*
Notes:
Logical Operators: Used to combine boolean expressions (&&, ||, !)
Conditional Operators: Used to compare values and return boolean results (==, !=, <, >, <=, >=)
Control Flow: Logical and conditional operators are essential for controlling the flow of a program,
particularly in if, for, and switch statements.

Logical:
AND (&&)
OR (||)
NOT (!)

Conditionals:
Equal (==)
Not Equal (!=)
Less Than (<)
Greater Than (>)
Less Than or Equal (<=)
Greater Than or Equal (>=)
*/
func part14() {
	fmt.Println("--------------------------------")
	fmt.Println("part14()")
	fmt.Println("--------------------------------")

	//  AND (&&)
	// The AND operator returns true if both operands are true. Otherwise, it returns false.
	// Syntax: expr1 && expr2
	a1 := true
	b1 := false
	fmt.Println("a1 && b1 =", a1 && b1) // Output: false

	//  OR (||)
	// The OR operator returns true if at least one of the operands is true. If both operands are false, it returns false.
	// Syntax: expr1 || expr2
	a2 := true
	b2 := false
	fmt.Println("a2 || b2 =", a2 || b2) // Output: true

	//  NOT (!)
	// The NOT operator inverts the boolean value of its operand. If the operand is true, it returns false, and vice versa.
	// Syntax: !expr
	a3 := true
	fmt.Println("!a =", !a3) // Output: false

	// 	Equal (==)
	// Checks if two values are equal.
	// Syntax: a == b
	a4 := 5
	b4 := 5
	fmt.Println("a4 == b4 =", a4 == b4) // Output: true

	// 	Not Equal (!=)
	// Checks if two values are not equal.
	// Syntax: a != b
	a5 := 5
	b5 := 3
	fmt.Println("a5 != b5 =", a5 != b5) // Output: true

	// 	Less Than (<)
	// Checks if the value on the left is less than the value on the right.
	// Syntax: a < b
	a6 := 3
	b6 := 5
	fmt.Println("a6 < b6 =", a6 < b6) // Output: true

	// 	Greater Than (>)
	// Checks if the value on the left is greater than the value on the right.
	// Syntax: a > b
	a7 := 5
	b7 := 3
	fmt.Println("a7 > b7 =", a7 > b7) // Output: true

	// 	Less Than or Equal (<=)
	// Checks if the value on the left is less than or equal to the value on the right.
	// Syntax: a <= b
	a8 := 3
	b8 := 3
	fmt.Println("a8 <= b8 =", a8 <= b8) // Output: true

	// 	Greater Than or Equal (>=)
	// Checks if the value on the left is greater than or equal to the value on the right.
	// Syntax: a >= b
	a9 := 5
	b9 := 3
	fmt.Println("a9 >= b9 =", a9 >= b9) // Output: true

	// 	Logical and conditional operators are frequently used in control flow statements to make decisions based on conditions.
	// Example: Using 'if' statement with logical and conditional operators
	age := 20
	isStudent := true

	// Using logical and conditional operators in an if statement
	if age >= 18 && isStudent {
		fmt.Println("Eligible for student discount")
	} else {
		fmt.Println("Not eligible for student discount")
	}

	// Example: Using 'for' loop with conditional operator
	// Using conditional operator in a for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		-------------------------------------------------------------------------
		if, if..else, if..else if..else

		'if' statement executes a block of code if a specified condition evaluates to true.
				if condition {
				code to be executed if the condition is true
			}
	*/

	age = 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	/* if...else statement executes one block of code if a condition is true, and another block of code if the condition is false.
		if condition {
	    // code to be executed if the condition is true
	} else {
	    // code to be executed if the condition is false
	}

	*/
	age = 16

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	/*
		if...else if...else Statement
		The if...else if...else statement allows you to test multiple conditions. It executes the first block of code whose condition evaluates to true. If none of the conditions is true, it executes the else block.

		if condition1 {
		// code to be executed if condition1 is true
		} else if condition2 {
				// code to be executed if condition2 is true
		} else {
				// code to be executed if none of the conditions are true
		}
	*/
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	/*
		Using Short Statement in if
		In Go, you can include a short statement to execute before the condition check. This is useful for variable initialization that is used only within the scope of the if statement.

		if initializer; condition {
		// code to be executed if the condition is true
		}
	*/

	// if err == nil on num, err := strconv.Atoi("123")
	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Println("Converted number:", num)
	} else {
		fmt.Println("Conversion failed:", err)
	}

}
