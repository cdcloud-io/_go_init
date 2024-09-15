package main

import (
	"fmt"
	"time"
	"unsafe"
)

func part04() {
	fmt.Println("Hello from 04_arrays_slices_maps_and_loops")

	var myArray0 [3]int32 // declare an array called myArray with 3 elements of int32 - therefore this array consumes 12bytes of memory - initializes to 0 if only declared
	myArray0[0] = 123
	myArray0[1] = 456
	//myArray[2] = 789

	fmt.Println("value at index0: ", myArray0[0])
	fmt.Println("value at index1: ", myArray0[1])
	fmt.Println("value at index2: ", myArray0[2])

	fmt.Println("value at element 1 and 2: ", myArray0[1:3]) // 3 is not inclusive
	fmt.Println("value at element 0 and 2: ", myArray0[1:2]) // 2 is not inclusive

	// print memory addresses for each element
	fmt.Println(myArray0)
	fmt.Println(&myArray0)
	fmt.Printf("memory address of myArray is %p\n", &myArray0)
	fmt.Println(&myArray0[0])
	fmt.Println(&myArray0[1])
	fmt.Println(&myArray0[2])

	var myArray1 [3]int32
	myArray1[0] = 999
	myArray1[1] = 888

	// how to set a pointer variable
	// 1st we need a pointer declaration
	var myArrayPointer *[3]int32

	// then we can assign an address to the pointer
	myArrayPointer = &myArray1
	fmt.Println("unsafe array pointer: ", myArrayPointer)

	fmt.Println("Array1:", myArray1)
	fmt.Printf("Address of array1: %p\n", &myArray1)
	fmt.Printf("Address of element 0: %p\n", &myArray1[0])
	fmt.Printf("Address of element 1: %p\n", &myArray1[1])
	fmt.Printf("Address of element 2: %p\n", &myArray1[2])

	// Using unsafe to get the addresses
	fmt.Printf("Unsafe address of array1: %x\n", uintptr(unsafe.Pointer(&myArray1)))
	fmt.Printf("Unsafe address of element 0: %x\n", uintptr(unsafe.Pointer(&myArray1[0])))
	fmt.Printf("Unsafe address of element 1: %x\n", uintptr(unsafe.Pointer(&myArray1[1])))
	fmt.Printf("Unsafe address of element 2: %x\n", uintptr(unsafe.Pointer(&myArray1[2])))

	type MyStruct struct {
		a int
		b float64
		c string
	}

	var s MyStruct
	fmt.Printf("Offset of 'a': %d\n", unsafe.Offsetof(s.a))
	fmt.Printf("Offset of 'b': %d\n", unsafe.Offsetof(s.b))
	fmt.Printf("Offset of 'c': %d\n", unsafe.Offsetof(s.c))

	// declare and init an array
	var intArr0 [3]int32 = [3]int32{1, 2, 3}

	fmt.Printf("Address of element 0: %p\n", &intArr0[0])
	fmt.Printf("Address of element 1: %p\n", &intArr0[1])
	fmt.Printf("Address of element 2: %p\n", &intArr0[2])

	fmt.Printf("Value of element 0: %d\n", intArr0[0])
	fmt.Printf("Value of element 1: %d\n", intArr0[1])
	fmt.Printf("Value of element 2: %d\n", intArr0[2])

	fmt.Println("value of intArr0: ", intArr0)

	// declare and infer fixed size
	/*
		In Go, the [... ] syntax is used to define an array with a size that is determined by the number of elements provided in the array literal.
	*/
	intArr1 := [...]int32{40, 30, 60, 20}
	fmt.Printf("Address of element 0: %p\n", &intArr1[0])
	fmt.Printf("Address of element 1: %p\n", &intArr1[1])
	fmt.Printf("Address of element 2: %p\n", &intArr1[2])
	fmt.Printf("Address of element 2: %p\n", &intArr1[3])

	fmt.Printf("Value of element 0: %d\n", intArr1[0])
	fmt.Printf("Value of element 1: %d\n", intArr1[1])
	fmt.Printf("Value of element 2: %d\n", intArr1[2])
	fmt.Printf("Value of element 2: %d\n", intArr1[3])

	fmt.Println("value of intArr1: ", intArr1)

	// ok lets benchmark

	/*
		## Slices ##
		slices are just wrappers around arrays.  they provide the additional metadata to facilitate
		a more general, powerful and convenient interface to sequences of data
		https://go.dev/doc/effective_go

			type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}

		note: slices
		n Go, slices are indeed implemented as a struct, but the actual struct type for slices is not accessible in your code. The Go runtime defines slices internally, and they are represented in a way that is not directly accessible or modifiable through your code. This is why you can't access the fields like array, len, or cap directly on a slice type like []int32.
	*/

	// by omiting the lenghth value, we now have a slice
	fmt.Println("#################################################################")
	var intSlice0 []int32 = []int32{4, 69, 70} // under the hood, [4, 69, 70]
	fmt.Println("before append: ", intSlice0)
	fmt.Printf("Address of intSlice0 before append: %p\n", &intSlice0)
	fmt.Printf("the length is %v, with capacity %v\n", len(intSlice0), cap(intSlice0))

	/* todo
	fmt.Println("showing struct offsets (can only be done with unsafe and reflect)")
	fmt.Println("#################################################################")
	// Accessing the underlying array using unsafe.SliceData
	// dataPtr := unsafe.SliceData(intSlice0) // no good
	dataPtr := unsafe.Pointer(&intSlice0)
	fmt.Printf("Data pointer: %p\n", dataPtr)

	// Get the slice header
	sliceHeader := unsafe.SliceData(intSlice0)

	// Print the slice header details
	// fmt.Printf("Slice Header: Data=%p, Len=%d, Cap=%d\n", sliceHeader , sliceHeader.Len, sliceHeader.Cap)
	fmt.Println("sliceheader is: ", *sliceHeader)
	const M, N = unsafe.Sizeof(intSlice0), unsafe.Sizeof(sliceHeader)
	fmt.Println(M, N) // 16 32

	// Access elements using the slice header
	arrayPtr := unsafe.Pointer(sliceHeader)
	firstElement := *(*int32)(arrayPtr)
	secondElement := *(*int32)(unsafe.Pointer(uintptr(arrayPtr) + unsafe.Sizeof(firstElement)))
	thirdElement := *(*int32)(unsafe.Pointer(uintptr(arrayPtr) + 2*unsafe.Sizeof(firstElement)))

	// Print the elements
	fmt.Printf("First element: %d\n", firstElement)
	fmt.Printf("Second element: %d\n", secondElement)
	fmt.Printf("Third element: %d\n", thirdElement)
	*/
	intSlice0 = append(intSlice0, 77, 688) // the append function takes in a slice as param1 and the value you want to append as 2nd arg. it then returns a slice with a new element appended.

	fmt.Println("after append: ", intSlice0)
	fmt.Printf("Address of intSlice0 after append: %p\n", &intSlice0)
	fmt.Printf("the length is %v, with capacity %v\n", len(intSlice0), cap(intSlice0))

	// [4,69,70,77,688,*] length is the elements that have int32 value, and the capacity includes the * which are just empty UNACCESIBLE elements.  if you try to read the index on a *, or you will get a index out of range error.

	// using the append function to concatenate slices.
	var intSlice1 []int32 = []int32{8, 9}
	intSlice0 = append(intSlice0, intSlice1...)
	fmt.Println("after append: ", intSlice0)
	fmt.Printf("Address of intSlice0 after append: %p\n", &intSlice0)
	fmt.Printf("the length is %v, with capacity %v\n", len(intSlice0), cap(intSlice0))

	// declaring slices with make
	var intSlice2 []int32 = make([]int32, 3) // 3 is the length of slice.
	fmt.Printf("Address of intSlice2: %p\n", &intSlice2)
	fmt.Println("value of intSlice2: ", intSlice2)
	fmt.Printf("the length is %v, with capacity %v\n", len(intSlice2), cap(intSlice2))
	// Address of intSlice2: 0xc000122090
	// the length is 3, with capacity 3

	var intSlice3 []int32 = make([]int32, 3, 8) // optionally we can also specify the capacity of 8 otherwise the default becomes the length of the slice
	// note ^: for performance its better to know what the capacity size should be for your program so it does not have to reallocate the underlying array which is a big perfomance hit.
	fmt.Printf("Address of intSlice3: %p\n", &intSlice3)
	fmt.Println("value of intSlice3: ", intSlice3)
	fmt.Printf("the length is %v, with capacity %v\n", len(intSlice3), cap(intSlice3))
	// Address of intSlice3: 0xc0001220a8
	// the length is 3, with capacity 8

	// append some data
	intSlice3 = append(intSlice3, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("Address of intSlice3 after append: %p\n", &intSlice3)
	fmt.Println("value of intSlice3 after append: ", intSlice3)
	fmt.Printf("the length after append is %v, with capacity after append is %v\n", len(intSlice3), cap(intSlice3))
	// value of intSlice3 after append:  [0 0 0 1 2 3 4 5 6 7 8]
	// the length after append is 11, with capacity after append is 16

	/*
		## MAPS ##
		map[string]int32
		maps are a set of {"key": "value"} pairs where you can lookup a value by its key.

		using make func to declare a map
		var myMap0 map[string]uint8 = make(map[string]uint8)  // this makes a map where the keys are of type string, and the values are of type uint8

		NOTE:  maps always return, even if there is no key, it will return the default value of the datatype ie 0 for int
		for this maps return an extra option value called `ok` which is a `bool` = to false if value is not in map, and true if it is.
	*/
	fmt.Println("######## MAPS ########")
	// declare a map
	var myMap0 map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap0)

	// declare and initialize a map - dont need make
	var myMap1 = map[string]uint8{"Adam": 23, "Sarah": 32, "Angela": 44, "Stephen": 48, "Aleks": 35}
	fmt.Println(myMap1["Adam"])  // 23
	fmt.Println(myMap1["Jason"]) // 0  remember the default value for uint8 is 0.  this can be bad as we wont get an error when querying invalid data.

	// using the optional ok return from MAP to see if the value was found via bool value of true/false on `ok`
	var ageJason, ok = myMap1["Jason"]
	if ok {
		fmt.Printf("age is: %v\n", ageJason)
	} else {
		fmt.Println("Invalid Name")
	}

	// because ok is declared
	var ageAdam uint8
	ageAdam, ok = myMap1["Adam"]
	if ok {
		fmt.Printf("age is: %v\n", ageAdam)
	} else {
		fmt.Println("Invalid Name")
	}

	// delete(map, key) function to remove a key:value from a map - this will DELETE by REFERENCE so no RETURN value is given
	delete(myMap1, "Adam")
	ageAdam, ok = myMap1["Adam"]
	if ok {
		fmt.Printf("age is: %v\n", ageAdam)
	} else {
		fmt.Println("Invalid Name")
	}

	/*
		## LOOPS ##
		Loops are fundamental control structures in programming that allow you to execute a block of code multiple times. In Go, the primary looping construct is the `for` loop, which is highly versatile and can be used in various forms to handle different looping scenarios.
	*/

	// iterate over a map,array,slice,string and channels using a for loop we can use the range keyword within our for loop
	for name := range myMap1 { // name is initialized by the for statment itself
		fmt.Printf("Name: %v\n", name) // when iteration over a map, no order is preserved. so the output can be listed differently
	}

	// to return the values as well
	for name, age := range myMap1 { // name is initialized by the for statment itself
		fmt.Printf("Name: %v: Age: %v\n", name, age) // when iteration over a map, no order is preserved. so the output can be listed differently
	}

	// we can iterate thru arrays/slices like this where i is the index and v is the value in the array or slice
	for i, v := range intArr1 { // i is scoped to the loop
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	// fmt.Println(i)  i here is undefined

	// Go does not have while loops per se, but this loop type is acheived with the for keyword again
	// this is your while loop in go, it will continue till i >= 10
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1 // go does not have i++
	}
	fmt.Println(i) // i is still in scope and value is 10

	// you can also omit the condition and break when i is >= 10
	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		//i = i + 1
		// or
		i++
	}

	// you can also acheive with a tradition loop with 3 distinct parts separated by semi-colon;
	// they are:
	// initialization: i = 0
	// condition: i < 10
	// post: i++ (this gets executed every time the loop is completed) choices: i++, i--, i += 10 (increment by 10), i -= 10 (decrement by 10), i *= 10 (mult by 10), i /= 10 (divide by 10)
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	// call the slice perf func to show how prealloacting capacity increases performance
	slicePerf()
}

/*
The use of underscores in numeric literals (1_000_000) for better readability was introduced in Go 1.13. These underscores are ignored by the compiler, making it easier to read large numbers.
*/

func slicePerf() {
	var n int = 1_000_000
	fmt.Println(n)
	var testSlice0 = []int{}
	var testSlice1 = make([]int, 0, n) // using make to declare a slice len 0 and a capacity of n (1_000_000)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice0, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice1, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
