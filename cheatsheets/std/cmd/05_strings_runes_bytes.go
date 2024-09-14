package main

import (
	"fmt"
	"strings"
	"time"
)

func part05() {

	// IMPORTANT: func part05() does not depend on the string imports.
	// func part05a() Does!!

	// this will also demonstrate the variable encoding length of UTF-8

	var myString0 = "résumé" // those accents are unicode 2 byte characters (non-ASCII)
	fmt.Println(myString0)

	// indexing a string
	indexed0 := myString0[0] // should be 114 uint8
	indexed1 := myString0[1] // the true value is 233 but will return 195 because its not reading the second byte.  the range keyword will automatically

	fmt.Printf("value at index0 is: %v\n", indexed0) // outputs 114
	fmt.Printf("value at index1 is: %v\n", indexed1) // outputs 195 (not technically correct as we are only caluculating value of byte 1 of 2)

	// lets print out the value and the type of the index using printf
	fmt.Printf("Output of index0: value %%v and Type %%T: ") // double %% escapes the % and prints it literally
	fmt.Printf("%v, %T\n", indexed0, indexed0)               // output: 114, uint8

	fmt.Printf("Output of index1: value %%v and Type %%T: ") // double %% escapes the % and prints it literally
	fmt.Printf("%v, %T\n", indexed1, indexed1)               // output: 114, uint8

	// iterate over the string with the range keyword where i is the index and v is the value
	for i, v := range myString0 {
		fmt.Printf("index: %v value: %v memory: %v Type: %T\n", i, v, &i, v)
	}
	/*
		index: 0 value: 114 memory: 0xc000012118 Type: int32
		index: 1 value: 233 memory: 0xc000012130 Type: int32
		index: 3 value: 115 memory: 0xc000012138 Type: int32
		index: 4 value: 117 memory: 0xc000012140 Type: int32
		index: 5 value: 109 memory: 0xc000012148 Type: int32
		index: 6 value: 233 memory: 0xc000012150 Type: int32

		notice the index goes from 1 to 3
		- strings still need to be represented as binary data 01011010. the early days we used ASCII encoding 7bits. for 128 english characters
		- to represent an extended set of character, we went to unicode.
		- utf-32 is 4 bytes 32bits but this will waste so much memory for each character. a would be 00000000 00000000 00000000 01100001
		- utf-8 solves this with variable bit encoding. so a would stay the same as in ASCII, 01100001 - the padded 0s ar dropped
	*/

	/*
		  - takeaway: when you are dealing with strings in Go, you are dealing with a value whode underlying representation is an array of bytes
			- this is also why taking the lenghth of a string is its length in the number of bytes and not the number of chars (characters)
	*/
	fmt.Printf("\nThe length of 'myString0' is %v bytes \n\n", len(myString0))

	// an easier way to deal with iterating and indexing strings is to cast them to an array of runes rather than dealing
	// with the underlying byte array of string.

	// - runes are just 'Unicode' point numbers which represent the character

	// here is what we get when we do that with our string
	myString2 := []rune("résumé")
	fmt.Printf("myString2 is an array of rune with values: %v\n\n", myString2) // myString2 is an array of rune with values: [114 233 115 117 109 233]

	fmt.Printf("not the len(myString2) is: %v\n", len(myString2))
	fmt.Printf("this is more closely compared to couting the array of char is c\n")
	fmt.Printf("so runes are just 'Unicode' point numbers which represent the character")

	// now when we iterate over a run we get sequential index numbers
	for i, v := range myString2 {
		fmt.Printf("index: %v value: %v memory: %v Type: %T\n", i, v, &i, v)
	}
	/*
		index: 1 value: 233 memory: 0xc000012178 Type: int32
		index: 2 value: 115 memory: 0xc000012180 Type: int32
		index: 3 value: 117 memory: 0xc000012188 Type: int32
		index: 4 value: 109 memory: 0xc000012190 Type: int32
		index: 5 value: 233 memory: 0xc000012198 Type: int32
	*/

	// to declare a rune we use single quotes 'a' like char in c

	var myRune0 = 'a'
	fmt.Printf("\nmyRune0 = %v\n\n", myRune0)

	// the following code throws an error: invalid argument: myRune0 (variable of type rune) for len
	// fmt.Printf("\nlength of myRune0 is : %v", len(myRune0))

	/*
		In Go, a rune is an alias for int32 and represents a Unicode code point. Therefore, a rune itself is always a single code point, and its length in terms of runes is always 1. However, if you want to find out how many bytes are used to encode that rune in UTF-8, you need to convert it to a string first and then get the length of that string.
	*/

	var myRune1 rune = '世' // Example rune (Chinese character)

	// Convert the rune to a string and get the length in bytes
	runeLength := len(string(myRune1))

	fmt.Printf("myRune1 is: %v length of myRune1 in bytes is: %v\n\n", string(myRune1), runeLength)

	/*
		By converting the rune to a string, you can then use len to get the number of bytes that the UTF-8 encoding of that rune occupies.
		len function: When used on a string, the len function returns the number of bytes in the string.
	*/

	// STRING BUILDING
	// as we saw before, we can concatenate strings using the += symbol like this
	var strSlice0 = []string{"m", "y", " ", "s", "t", "r", "i", "n", "g", " ", "s", "l", "i", "c", "e", "0", "."}
	fmt.Printf("\nstrSlice0 = %v\n", strSlice0)

	var catStr0 = ""
	var catStr0Time = time.Now()
	for i := range strSlice0 {
		catStr0 += strSlice0[i] // see NOTE below
	}
	var catStr0TimeItTook = time.Since(catStr0Time)
	fmt.Println("Concatonated []string slice is:")
	fmt.Printf("%v\n", catStr0)
	fmt.Printf("Time is took to concatonate: %v\n\n", catStr0TimeItTook) // ~900 ns  thats SLOW

	// ** NOTE: Strings are IMMUTABLE in GO, meaning I cannot modify them once created
	// uncomment line below to show error trying to modify an index element of a string
	// catStr0[0] = 'a'

	// NOTE: so it shows that when we are concatonating a string and assigning it to a variable like this: catStr0 += strSlice0[i]
	// we are actually creating a completely new string every time. WHICH IS PRETTY INEFFICIENT. for this we can use Go's string
	// builder by importing the strings package.  we will use the string builder in next function call 05a

	// now we will jump to function part05a() which requires the strings import package to solve these inefficiencies
	// note: comment out the fucntion call and the function when testing strings raw without the help of the strings import
	part05a()

}

func part05a() {

	// here we will depend and use the strings import package to solve the inefficiencies of strings
	// Once we import Go strings package, we can envoke the string builder using the same example in func05

	var strSlice1 = []string{"m", "y", " ", "s", "t", "r", "i", "n", "g", " ", "s", "l", "i", "c", "e", "1", "."}
	fmt.Printf("\nstrSlice1 = %v\n", strSlice1)

	// var catStr1 = "" we will omit initializing an empty string and in sted initialize the string.Builder
	var strBuilder strings.Builder // we need to imports strings package here
	var catStr1Time = time.Now()

	for i := range strSlice1 {
		strBuilder.WriteString(strSlice1[i]) // instead of using the += operator, call the WriteString() method
		// what is happening here is an internal array is created that is appended with each call of the WriteSting() method
	}

	var catStr1 = strBuilder.String() // finally we call the String() method. This creates a new string created from the
	// appended values of the internal array created by the WriteString() method.
  // this is way more performant as outlined by the timers.

	var catStr1TimeItTook = time.Since(catStr1Time)
	fmt.Println("stringBuilder string slice is:")
	fmt.Printf("%v\n", catStr1)
	fmt.Printf("Time is took to concatonate: %v\n\n", catStr1TimeItTook) // ~300 ns MUCH FASTER
}
