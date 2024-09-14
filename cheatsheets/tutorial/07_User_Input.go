package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Notes:
Explanation of Buffers
What is a Buffer?
-----------------
- Temporary Storage: A buffer temporarily holds data that is being transferred between two locations. It is typically implemented as a memory array.

- Data Transfer: Buffers are used to manage the differences in rate at which data is processed or transferred. For example, when reading data from a disk (slow) and processing it in memory (fast), a buffer helps to manage the data flow smoothly.

How Buffers Work
----------------
- Reading and Writing: When reading data from a slow device, data is first read into a buffer. The program then processes the data from the buffer, allowing the program to work with the data at its own pace.

- Buffered I/O: In buffered I/O operations, data is read into a buffer before being processed. This reduces the number of I/O operations, which can be expensive in terms of time and resources.

Where are Buffers Created
-------------------------
Most often HEAP allocation

The creation and management of buffers are typically handled by the runtime, and they can be allocated on either the heap or the stack, depending on how they are used. However:
In most cases, buffers used for I/O operations are allocated on the HEAP because they need to be dynamically sized and persist beyond the scope of a single function call.
*/
func part07() {
	fmt.Println("--------------------------------")
	fmt.Println("part07()")
	fmt.Println("--------------------------------")

	fmt.Println("what is your name?")
	// setups a buffered reader that gets text from keyboard from os.Stdin
	reader := bufio.NewReader(os.Stdin)

	// we are gonna store this input inside of name, but we are also going
	// to protect the application from crashing by handling an error. The function
	// signature will show if an error can be returned.
	//  ReadString() can return an error `func (b *bufio.Reader) ReadString(delim byte) (string, error)`

	// if you did not want to handle the error (BAD PRACTICE) you can omit the err by using a _ instead
	// name, _ := reader.ReadString('\n')

	name, err := reader.ReadString('\n') // delimeter byte is set to \n which is <enter>
	if err == nil {
		fmt.Println("hello, ", name)
	} else {
		log.Fatal(err)
	}

}
