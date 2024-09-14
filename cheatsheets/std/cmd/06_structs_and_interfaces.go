package main

import (
	"fmt"
)

/*
	- a struct is used to define your own datatype.
	- to create a struct we use the 'type' keyword because again we are defining a TYPE here.
	- we need the '<type/stuctName>' of our 'TYPE' followed by the 'struct' keyword and curly brace '{}' no spaces  'struct{}'
	ie. 'type myType struct{}'  // no spaces after struct for the curly brace.  this is not a function block rather a Fields block
	which does not have a space

	- structs can hold mixed types in the form of "FIELDS" which we can define by NAME.
*/

// here we define a gasEngine TYPE which holds a miles per gallon field of type unsigned int 8 (uint8)
// and a gallons fieid of the same type which is uint8 which represents how many gallons of fuel are left
// unless its an anonymous struct (struct with no name usually defined in scope of a function) structs
// should be defined globally to the package.  so we will comment out the struct bellow and copy to the top of
// file before the part06 func call
/*
type <typeName> struct {
    FieldName1 FieldType1
    FieldName2 FieldType2
}
TypeName: This is the name you give to the struct type. It defines a new data type.
FieldName: These are the names of the fields within the struct.
FieldType: These are the data types of the fields within the struct.
*/

// global named structs definitions

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

// with this struct defined, we can go into our function and create a new variable of type 'gasEngine'

func part06() {
	fmt.Println("Entering part06() structs and interfaces")

	// with gasEngine struct defined, lets create a variable of type 'gasEngine'
	// note: if you just declare a var of type gasEngine, it is considered a ZERO value struct
	// meaning the assosiated fields will be initialized with their default values.
	// in our case uint8 defaults to 0
	// so just by declaring 'myEngine' we have a 'myEngine' that looks like this:
	/*
		myEngine{
		  mpg: 0
			gallons: 0
		}
	*/
	var myEngine0 gasEngine
	fmt.Println(myEngine0)                        // output: {0 0}
	fmt.Println(myEngine0.mpg, myEngine0.gallons) // output: 0 0

	// to declare and set the values we use the 'struct literal' syntax like this"
	var myEngine1 = gasEngine{mpg: 25, gallons: 15} // notice we need to use the assignment operator =
	fmt.Println(myEngine1)                          // output: {25 15}

	// you can also omit the 'FIELD' names and as long as you know the order, you can declare and assign like this:
	var myEngine2 = gasEngine{26, 16} // notice we need to use the assignment operator =
	fmt.Println(myEngine2)            // output: {26 16}

	// you can also set the values by name directly like this: <objName>.<fieldName>
	var myEngine3 = gasEngine{30, 10} // notice we need to use the assignment operator =
	// using the <objName>.<fieldName>
	myEngine3.mpg = 31
	fmt.Println(myEngine3) // output: {31 10}

	part6a()

}

// The fields of a struct can be anything you want... even another struct

// lets define an owner struct and make that a field in gasEngine2 struct that defines the gas engines owners name
type owner struct {
	name string
}

type gasEngine2 struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

func part6a() {

	// declare and initialize myEngine2
	// this declaration fails because "Stephen" is of Type string, and ownerInfo is of Type owner. //uncomment to see error

	// var myEngine3 gasEngine2{10, 20, "Stephen"}

	// you can initialize 3 ways.. create a var of type owner

	// var owner0 owner = owner{name:"stephen"} // output: {stephen}
	// var owner0 owner = owner{"stephen"} // output: {stephen}
	owner0 := owner{"stephen"} // output: {stephen}
	fmt.Println(owner0)

	// then we can make a gasEngine2 with owner var owner0
	var myEngine2 gasEngine2 = gasEngine2{10, 20, owner0}
	fmt.Println(myEngine2) // output: {10 20 {stephen}}

	// or inline initialization with owner (meaning we are defining and creating the owner as part of the parameters)
	var myEngine3 gasEngine2 = gasEngine2{10, 20, owner{"alex"}}
	fmt.Println(myEngine3) // output: {10 20 {alex}}

	// to access owner information
	fmt.Println(myEngine3.ownerInfo.name) // output: alex // ownerinfo is the field, which is a struct with another field called 'name'

	part6b()

}

// you can also set a field as a struct.  so instead of having an object like this:
/*
  myEngine{
	mpg: 25
	gallons: 15
	ownerInfo.name "alex"
	}

	we can set the field 'ownerInfo' as just the owner struct which default to type owner.  you can do this with any type.
	we will also add a field 'int' which is also of type 'int'
	type myEngine struct{
	mpg uint8
	gallons uint8
	owner
	int
  }

	now when we extantiate an object the value will be as follows:

	myEngine{
	mpg: 25
	gallons: 15
	name: "alex"
	int: 10
	}

*/

type anotherEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

func part6b() {

	var myCar0 anotherEngine = anotherEngine{10, 20, owner{"charlie"}, 10}
	fmt.Println("myCar0 owner is: ", myCar0.name)
	fmt.Printf("myCar0 gallons is: %v and mpg is: %v\n", myCar0.gallons, myCar0.mpg)
	fmt.Printf("myCar0 has a field called int which is: %v\n\n", myCar0.int)

	part6c()

}

// ANONYMOUS Structs

// you can define aunonymous struct which dont have a name type.
// aunonymous structs have to be declared/define and intialized it in the same place.

func part6c() {

	// create and define an anounymous struct.
	var myAnonEngine = struct {
		mpg     uint8
		gallons uint8
	}{22, 10} // this is not reusable.  if you wanna create another struct like this, you have to rewrite the entire definition again

	fmt.Printf("\nmyAnonEngine values - mpg: %v - gallons: %v\n", myAnonEngine.mpg, myAnonEngine.gallons)

	part6d()

}

// METHODS

// struct also have a concept of methods which we can use as well.
// methods are functions that are directly tied to the struct and have access to the struct instance itself

type fuelEngine struct {
	mpg     uint8
	gallons uint8
}

// let say we want a method that calculate the miles left on a gas tank
// a method is just like a functions except we add an assignment to the fuelEngine type with the (e fuelEngine)
// this function now has access to all the fields and even other methods assigned to the fuelEngine type
func (e fuelEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// this is now a function because its not tied to the struct itself, but takes in a struct method
func canMakeIt(e fuelEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Printf("Warning: you dont have enough gas to drive %v miles\n\n", miles)
	}
}

func part6d() {
	fmt.Printf("\n\nfunc part 6d\n")
	var myFuelEngine fuelEngine = fuelEngine{8, 100} // this is like oop where we define a class on this line
	// then extantiate on of its methods. 'myFuelEngine.milesLeft()'

	// to access the method now we can do something like this:
	fmt.Printf("Total miles left in tank: %v\n", myFuelEngine.milesLeft()) //output is 32???
	//note: we are multiplying 2 uint8 and returning a uint8, however the multiplicaton exceeds the size
	// of uint8 (255), it wraps around due to overflow

	// can we make it there
	canMakeIt(myFuelEngine, 200)

	part6e()

}

// INTERFACES

/*

suppose we have 2 different engine types. a gas and electric engine.  the electric engine has different fields.
- instead of mpg, it has mile per kilowatt hour:  mpkwh
- instead of gallons it has kilowatt hours: kwh which specifies how much charge is left in the battery

also we will create a similar milesLeft() method

*/

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// electricEngines milesLeft() method signature with fullfils the engine interface contract.
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// now currently our canMakeIt() function only takes in a gasEngine. but what if we wanted to take in any engine type
// *** this is were interfaces come in ***

// to define an interface type, we use the keyword type, followed by the interfaceName followed by the keyword interface
// type <typeName> interface

// in our canMakeIt() functions from earlier, we see all we really need is a miles left method on the engine struct
// which takes no parameters and returns an uint8 // this is also called the method signature
/*
func canMakeIt(e fuelEngine, miles uint8) { // <-- the first line before the { is also called the method signature
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Printf("Warning: you dont have enough gas to drive %v miles\n\n", miles)
	}
}

- we can specify this method signature in our interface like this:

type engine interface {
	milesLeft() uint8
}

- now we can change the canMakeIt function (now canMakeIt2) signature to take in an engine (more generic) rather then a specific engine.
- our new function can now take in anything for engine AS LONG AS IT IMPLEMENTS a milesLeft() method with the
signature we specified in out interface. milesLeft()

*/

// with this new function, we can apply to a wider range of engine types. hooray
func canMakeIt2(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Printf("Warning: you dont have enough gas to drive %v miles\n\n", miles)
	}
}

type engine interface {
	milesLeft() uint8
}

func part6e() {

	var myElectricMotor electricEngine = electricEngine{20, 200}
	canMakeIt2(myElectricMotor, 100)

	part6f()

}

// additional fun with Go way of doing enums.
// go does not have enums, instead this is how we enforce values

// Define a custom type for fuelType
type fuelType int

// Define constants for the fuelType enum
const (
	Gas      fuelType = iota // Gas is 0
	Diesel                   // Diesel is 1
	Electric                 // Electric is 2
	Hydrogen                 // Hydrogen is 3
)

// String method to convert fuelType to a string representation
/*
	In Go, the [... ] syntax is used to define an array with a size that is determined by the number of elements provided in the array literal.
*/
func (f fuelType) String() string {
	return [...]string{"Gas", "Diesel", "Electric", "Hydrogen"}[f]
}

/*
Step by Step
- First, we define the fuelType enum using iota:

type fuelType int

const (
    Gas fuelType = iota
    Diesel
    Electric
    Hydrogen
)

- fuelType is a custom type based on int

- The constants Gas, Diesel, Electric, and Hydrogen are of type fuelType and are assigned the values 0, 1, 2, and 3 respectively, due to the use of iota.

- Next, we define a String method for the fuelType type:

func (f fuelType) String() string {
    return [...]string{"Gas", "Diesel", "Electric", "Hydrogen"}[f]
}

- array is intialized with [...]

- The [...] syntax creates an array of strings with the elements "Gas", "Diesel", "Electric", and "Hydrogen"

- The length of this array is inferred from the number of elements, which is 4 in this case.
  This array has the type [4]string because it has 4 elements.

- The expression [f] accesses the element at index f in the array.

- Since f is of type fuelType and can be 0 (Gas), 1 (Diesel), 2 (Electric), or 3 (Hydrogen),
  it directly maps to the corresponding string in the array.
  For example, if f is 1, this expression evaluates to "Diesel".

*/

type airPlaneEngine struct {
	rpm  uint32
	fuel fuelType
}

func part6f() {
	var myFuel fuelType = Gas
	fmt.Printf("myFuel.String() is %v, Datatype: %T\n", myFuel.String(), myFuel.String()) // Output: Gas
	fmt.Printf("myFuel is %v, Datatype: %T\n", myFuel, myFuel)

	var myJetEngine airPlaneEngine = airPlaneEngine{10_000, Diesel}
	fmt.Printf("myJetEngine uses %v for fuel, and rotates at %v RPM\n\n", myJetEngine.fuel.String(), myJetEngine.rpm)

}
