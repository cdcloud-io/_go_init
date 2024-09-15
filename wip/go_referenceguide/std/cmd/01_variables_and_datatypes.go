package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

/* ## Variables and Constants ## */
var myVariable0 string = "myString" // explicit, scoped to package
var myVariable1 = "myString"        // infered,  scoped to package
// myVariable2 := "myString"        // infered but this type of declaration and assignment must be in a function block.
var myVariable3 string // initialized as ""

func varFunc0() {
	myVariable2 := "myString" // infered but this type of declaration and assignment must be in a function block.
	fmt.Println("varFunc0: myVariable2 is:", myVariable2)

	// block scoped variables
	{
		const myVariable5 string = ""
		myVariable4 := "myString"
		fmt.Println("varFunc0: myVariable4 is:", myVariable4)

	}

	// multi variable initialization
	var var1, var2 = 1, 2
	var3, var4 := "hello", "soda"
	fmt.Println(var1, var2) // output: 1 2
	fmt.Println(var3, var4) // output: hello soda

	// works but avoid as another developer has no clue what the datatype for myVar is without hovering over foo() return type
	myVar0 := foo()
	fmt.Println(myVar0)

	// better way
	var myVar1 int = foo()
	fmt.Println(myVar1)

}

/* ## Constents and Scoping ## */
// constant used for storing values we dont want the program to change.
// constants cannot change and have to be declared and initialized at once.
// Package-level constant (globally scoped within the package)
const GlobalConstant0 = "I am a global constant"
/*
- Global variables can be accessed from any function within the same package.
- Global variables are generally discouraged unless necessary, as they can make the code 
    harder to maintain and understand. It's often better to pass variables as parameters to functions.
*/

func part01() {
	fmt.Println("Hello from 01_variables_and_datatypes")
	varFunc0()

	fmt.Println(GlobalConstant0) // Accessible here

	// Function-level constant (locally scoped within the function)
	const LocalConstant = "I am a local constant"
	fmt.Println(LocalConstant) // Accessible here

	anotherFunction()

	// Uncommenting the following line would cause a compile-time error
	// fmt.Println(AnotherLocalConstant)

	// calling othere functions
	myInts0()
	myUints0()
	myFloats0()
	myStrings0()
	part02()
	myLen0()
	part03()
	part04()

}

func anotherFunction() {
	fmt.Println(GlobalConstant0) // Accessible here

	// Another function-level constant (locally scoped within the function)
	const AnotherLocalConstant = "I am another local constant"
	fmt.Println(AnotherLocalConstant) // Accessible here
}

/* ## Datatypes ##
bool: true/false // defaults to false
floats: float32, float64 // use 64 for accuracy when possible. // defaults to 0
integer: int, int8, int16, int32, int64 // signed whole numbers - int by itself assumed architecture type. ie 64 bit cpu, int == int64 //defaults to 0
unsigned integer: uint, uint8, uint16, uint32, uint64 // unsigned whole numbers // defaults to 0
rune: like a char but used for getting character counts ajusting for utf-8 variable char lenghth // defaults to ''
string: "something in double quotes" // defaults to ''

special datatypes
error: declare in a function // defaults to nil
var err error

*/

func myBools0() {
	var myBool0 bool // initializes to false
	var myBool1 bool = true

	fmt.Println("myBool0 is: ", myBool0)
	fmt.Println("myBool1 is: ", myBool1)

}

func myInts0() {
	// Signed integer types
	fmt.Println("Signed Integer Types:")
	fmt.Printf("int8:   [%d, %d]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:  [%d, %d]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:  [%d, %d]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:  [%d, %d]\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("int:    [%d, %d]\n", math.MinInt, math.MaxInt) // typically int is 32-bit on 32-bit systems

}

func myUints0() {
	// Unsigned integer types
	fmt.Println("Unsigned Integer Types:")
	fmt.Printf("uint8:  [0, %d]\n", math.MaxUint8)
	fmt.Printf("uint16: [0, %d]\n", math.MaxUint16)
	fmt.Printf("uint32: [0, %d]\n", math.MaxUint32)
	fmt.Printf("uint64: [0, %d]\n", uint64(math.MaxUint64))
	fmt.Printf("uint:   [0, %d]\n", uint64(math.MaxUint)) // typically uint is 32-bit on 32-bit systems

	fmt.Println("Unsigned Integer Types:")
	fmt.Printf("uint8:  [0, %d]\n", math.MaxUint8)
	fmt.Printf("uint16: [0, %d]\n", math.MaxUint16)
	fmt.Printf("uint32: [0, %d]\n", math.MaxUint32)
	fmt.Printf("uint64: [0, %d]\n", uint64(math.MaxUint64))
	fmt.Printf("uint:   [0, %d]\n", uint64(math.MaxUint))

}

/*
In Go, the default format used by fmt.Println for floating-point numbers is scientific notation for large and small values.
This is why you see the float printed as 1.23456789e+07 instead of 12345678.9.

To control the formatting of floating-point numbers, you can use the fmt.Printf function with format specifiers.
The %f format specifier allows you to specify the number of decimal places.
fmt.Printf("float32: %.4f\n", float32Num0)
*/
func myFloats0() {
	var float32Num0 float32 = 12345678.9
	fmt.Println("float32: ", float32Num0)             // Println defaults to scientific notation on large numbers 1.xxxx..e+07
	fmt.Printf("float32: %.4f\n", float32Num0)        // Printf will display as intended. xxx9.0000
	fmt.Printf("float32 10.4: %10.4f\n", float32Num0) // Printf will display as intended.

	var float64Num0 float64 = 12345678.9
	fmt.Println("float64: ", float64Num0) // Println defaults to scientific notation on large numbers 1.xxxx..e+07

	/*
		%020.4f: The format specifier 20 means the total width of the output will be 20 characters.
		The 0 indicates that the number will be padded with zeros. The .4 specifies four digits after the decimal point.
	*/

	fmt.Printf("float32: %020.4f\n", float32Num0) // format printed to be a total width of 20 (including the .0000) padded with 0 (zeros)

}

/* Common String Formating Verbs - https://pkg.go.dev/fmt#hdr-Printing
%t	the word true or false
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%p	base 16 notation, with leading 0x //pointer reference
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

*/

func myStrings0() {
	var string0 string // declared with no value, default value becomes ""
	fmt.Println("declared string0: ", string0)

	// OK use of string declaration and initialization
	var string1 string = "hello world"
	var string2 string = "hello\nworld"
	var string3 string = `hello
	  world`
	var string4 string = "hello" + " world" // string concat ok here

	fmt.Println("string1: ", string1)
	fmt.Println("string2: ", string2)
	fmt.Println("string3: ", string3)
	fmt.Println("string4: ", string4)

	// rune is like char
	var myRune0 rune = 'a'
	fmt.Println("myRune0: ", myRune0)
	fmt.Println(myRune0) // 97
	/*
		97 maps to 'a' in UTF-8:

		Decimal: 97
		Hexadecimal: 0x61
		Character: 'a'
	*/

	var utf8CodeDec int = 97
	var utf8CodeHex int = 0x61
	var utf8CodeBin uint8 = 0b1100001
	char0 := rune(utf8CodeDec)
	char1 := rune(utf8CodeHex)
	char2 := rune(utf8CodeBin)
	fmt.Printf("UTF-8 code %d represents the character '%c'\n", utf8CodeDec, char0)
	// UTF-8 code 97 represents the character 'a'

	fmt.Printf("UTF-8 code %x represents the character '%c'\n", utf8CodeHex, char1)
	// UTF-8 code 61 represents the character 'a'

	// print in hexidecimal notation
	fmt.Printf("UTF-8 code 0x%x represents the character '%c'\n", utf8CodeHex, char1)
	// UTF-8 code 0x61 represents the character 'a'

	fmt.Printf("UTF-8 code %b represents the character '%c'\n", utf8CodeBin, char2)
	// UTF-8 code 1100001 represents the character 'a'

}

func myLen0() {
	fmt.Println("we are in myLen0 func")
	// getting lenghth of a string returns bytes - characters encoded outside of the ASCII character set will use more than 1 byte
	fmt.Println(len("hello"))  // 5 bytes
	fmt.Println(len("🤣"))      // 4 bytes
	fmt.Println(len("hello👋")) // 9 bytes

	// to count the lenghth in number of characters in a unicode string, use runes by importing  "unicode/utf8"
	fmt.Println(utf8.RuneCountInString("hello👋")) //6 characters

}

func foo() int {
	return 1
}
