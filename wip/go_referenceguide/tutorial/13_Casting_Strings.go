package main

import (
	"fmt"
	"strconv"
)

/*
Notes:
In Go, the strconv package provides functions for converting between strings and various other data types. This package is essential for tasks where you need to parse numeric values from strings or format numbers as strings. Below is a comprehensive guide on how to use the strconv package for string casting.

String to Integer: Use strconv.Atoi or strconv.ParseInt.
Integer to String: Use strconv.Itoa or strconv.FormatInt.
String to Float: Use strconv.ParseFloat.
Float to String: Use strconv.FormatFloat.
String to Boolean: Use strconv.ParseBool.
Boolean to String: Use strconv.FormatBool.

*/
func part13() {
	fmt.Println("--------------------------------")
	fmt.Println("part13()")
	fmt.Println("--------------------------------")

	// string to int
	// To convert a string to an integer, you can use the strconv.Atoi or strconv.ParseInt function.

	s1 := "123"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted integer:", i1)
	}

	// strconv.ParseInt
	// More flexible than strconv.Atoi, allows specifying the base and bit size.
	// Converts a string to an int64.

	s2 := "123"
	i2, err := strconv.ParseInt(s2, 10, 64) // base 10, 64-bit size
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted integer:", i2)
	}

	// Converting Integer to String
	// To convert an integer to a string, you can use the strconv.Itoa or strconv.FormatInt function.

	// strconv.Itoa
	// Converts an int to a string.
	i3 := 123
	s3 := strconv.Itoa(i3)
	fmt.Println("Converted string:", s3)

	// strconv.FormatInt
	// Converts an int64 to a string.
	// Allows specifying the base for the conversion.
	i4 := int64(123)
	s4 := strconv.FormatInt(i4, 10) // base 10
	fmt.Println("Converted string:", s4)

	// Converting String to Float
	// To convert a string to a float, you can use the strconv.ParseFloat function.
	s5 := "123.45"
	f5, err := strconv.ParseFloat(s5, 64) // 64-bit precision
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted float:", f5)
	}

	// Converting Float to String
	// To convert a float to a string, you can use the strconv.FormatFloat function.
	f6 := 123.45
	s6 := strconv.FormatFloat(f6, 'f', 2, 64) // format as float, 2 decimal places, 64-bit precision
	fmt.Println("Converted string:", s6)

	// Converting String to Boolean
	// To convert a string to a boolean, you can use the strconv.ParseBool function.
	s7 := "true"
	b7, err := strconv.ParseBool(s7)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted boolean:", b7)
	}

	// Converting Boolean to String
	// To convert a boolean to a string, you can use the strconv.FormatBool function.
	b8 := true
	s8 := strconv.FormatBool(b8)
	fmt.Println("Converted string:", s8)

}
