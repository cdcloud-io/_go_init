package main

import (
	"fmt"
)

/*
Notes:
In Go, type casting refers to the conversion of a value from one data type to another. This is useful when you need to perform operations between different types or when interfacing with APIs that require specific types. Go supports both implicit and explicit type conversions, with explicit conversions being the more common and preferred method.

Explicit Type Casting
Explicit type casting in Go is done using the syntax: T(v), where T is the target type and v is the value to be converted.

Type Conversion Caveats
- Precision Loss: When converting between types with different precision (e.g., float to int), you may lose information.

- Compatibility: Not all types can be converted directly. For instance, you cannot directly convert a string to an int; you need to use a function like strconv.Atoi from the strconv package.

- Safe Conversions: Always ensure that conversions make logical sense and are safe to perform to avoid runtime errors or data corruption.
*/
func part12() {
	fmt.Println("--------------------------------")
	fmt.Println("part12()")
	fmt.Println("--------------------------------")

	// casting int to float
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("i = 42, float64(i) is: ", f)

	// casting float to int
	f = 42.9
	i = int(f)
	fmt.Println("f = 42.9, int(f) is: ", i)

	// byte to slice
	var s string = "hello"
	var b []byte = []byte(s)

	// Print the memory address of the string s
	fmt.Printf("Memory address of s string is: %p\n", &s)

	// Print the contents of the slice b
	fmt.Println("Slice b which is the byte representation of the string hello is:", b)

	// Print the address of the slice header b
	fmt.Printf("Memory address of slice header b: %p\n", &b)

	// Print the address of the underlying array
	if len(b) > 0 {
		fmt.Printf("Memory address of the underlying array: %p\n", &b[0])
	}
	/*
		In Go, when you use &b on a slice, you are getting the address of the slice header, not the address of the underlying array. The slice header contains a pointer to the underlying array, the length of the slice, and its capacity. However, when you print &b, fmt.Println does not automatically dereference it to show the address; instead, it shows the contents of the slice in a readable format.
	*/

	// byte to slice
	var myByte []byte = []byte{104, 101, 108, 108, 111}
	var mySting string = string(b)
	fmt.Printf("Byte Slice: %v, String: %s\n", myByte, mySting)


	// type assertion for interfaces
	  var myInterface interface{} = "hello"
    
    // Type assertion
    i2 := myInterface.(string)
    fmt.Println("i2 is: ", i2)
    
    // Type assertion with check
    i2, ok := myInterface.(string)
    if ok {
        fmt.Println("String:", i2)
    } else {
        fmt.Println("Not a string")
    }

}
