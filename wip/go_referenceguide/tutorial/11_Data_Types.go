package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
Notes:
// primitive datatypes in GO and default assignment

int: 0, float32/64: 0.0, bool: false, string: "", rune "" // Note: always use float64, GO defaults to this on x86_64

int defaults to system architecture x86_64 mean int64

int8 - 1byte
int16 - 2bytes
int32 - 4bytes
int64 - 8bytes

uint8 - 1byte
uint16 - 2bytes
uint32 - 4bytes
uint64 - 8bytes
*/
func part11() {
	fmt.Println("--------------------------------")
	fmt.Println("part11()")
	fmt.Println("--------------------------------")

	fmt.Println("Primitive Data Types in GO")
	fmt.Println("--------------------------")

	var myInt int = 25
	var myUint8 uint8 = 254
	var myFloat float64 = 1.123
	var myBool bool = true
	var myRune rune = 'a'
	var myRune2 rune = '🙂'

	fmt.Println("myInt is: ", myInt)
	fmt.Println("myUint8 is: ", myUint8)
	fmt.Println("myFloat is: ", myFloat)
	fmt.Println("myBool is: ", myBool)
	fmt.Println("myRune is: ", myRune)
	fmt.Println("myRune2 is: ", myRune2)
	// fmt.Printf("myRune is: %c\n", myRune)
	// fmt.Printf("myRune2 is: %c\n", myRune2)
	fmt.Println("")
	fmt.Println("Type of using reflect")
	fmt.Println("----------------------")
	fmt.Println("myInt is type: ", reflect.TypeOf(myInt))
	fmt.Println("myUint8 is type: ", reflect.TypeOf(myUint8))
	fmt.Println("myFloat is type: ", reflect.TypeOf(myFloat))
	fmt.Println("myBool is type: ", reflect.TypeOf(myBool))
	fmt.Println("myRune is type: ", reflect.TypeOf(myRune))
	fmt.Println("myRune2 is type: ", reflect.TypeOf(myRune2))
	fmt.Println("myRune2 is type: ", reflect.TypeOf(myRune2))

	// Signed integer types
	fmt.Println("")
	fmt.Println("Signed Integer Types Value Range:")
	fmt.Println("---------------------")
	fmt.Printf("int8:   [%d, %d]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:  [%d, %d]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:  [%d, %d]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:  [%d, %d]\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("int:    [%d, %d]\n", math.MinInt, math.MaxInt) // typically int is 32-bit on 32-bit systems

	// Unsigned integer types
	fmt.Println("")
	fmt.Println("Unsigned Integer Types Value Range:")
	fmt.Println("-----------------------")
	fmt.Printf("uint8:  [0, %d]\n", math.MaxUint8)
	fmt.Printf("uint16: [0, %d]\n", math.MaxUint16)
	fmt.Printf("uint32: [0, %d]\n", math.MaxUint32)
	fmt.Printf("uint64: [0, %d]\n", uint64(math.MaxUint64))
	fmt.Printf("uint:   [0, %d]\n", uint64(math.MaxUint)) // typically uint is 32-bit on 32-bit systems

}
