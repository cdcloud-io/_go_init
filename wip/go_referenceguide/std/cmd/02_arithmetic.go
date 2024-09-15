package main

import "fmt"

/*
 */
func part02() {
	fmt.Println("Hello from 02_arithmetic")

	//var floatNum0 float64 = 1.0
	var floatNum1 float64 = 1.1
	var intNum0 uint8 = 10
	//var intNum1 uint64 = 25
	var result0 uint8 = uint8(floatNum1) + intNum0

	fmt.Println("result0: ", result0)

	var intNumX int
	var intNum1 int = 20
	intNum2 := 30
	intNum3 := 0

	fmt.Println("unitialized intNumX = ", intNumX)

	/*
		Print, Println does not allow string concatonation - next line is an error
		fmt.Println("num2 / num1 = " + intNum2/intNum1 + "Remainder: " + intNum2%intNum1)

		Instead, you need to convert the integers to strings first. Also, for a cleaner output
		and better formatting, you should use fmt.Sprintf or fmt.Printf.
	*/

	// Using fmt.Sprintf to format the string with integer values
	result := fmt.Sprintf("num2 / num1 = %d Remainder: %d", intNum2/intNum1, intNum2%intNum1)
	fmt.Println(result)

	// Alternatively, using fmt.Printf directly
	fmt.Printf("num2 / num1 = %d Remainder: %d\n", intNum2/intNum1, intNum2%intNum1)

	fmt.Println("num3 / num2 = ", intNum3/intNum2)

	// divide by 0 runtime error - will compile
	// fmt.Println("num2 / num3 = ", intNum2/intNum3)

	/*
	output:
	panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.arithmetic0()
        /home/stephen/code/workspaces/go-workspace/cheatsheets/std/cmd/02_arithmetic.go:42 +0x28a
main.main()
        /home/stephen/code/workspaces/go-workspace/cheatsheets/std/cmd/01_variables_and_datatypes.go:49 +0xaf
exit status 2
*/



}
