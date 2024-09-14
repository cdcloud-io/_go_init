package main

import "fmt"

/*
Go structs are a way to group related data together to form complex data structures. They serve as the foundation for creating objects in Go, which is a statically typed language. Unlike classes in object-oriented languages, Go's structs are purely data holding entities and don't inherently support methods or inheritance. However, Go uses interfaces and methods defined on struct types to achieve polymorphism and encapsulation.

*/

/*
### Basic Structure

A struct is defined using the `struct` keyword, followed by a series of fields that define the data types. Each field has a name and a type. Here’s a simple example:

*/

type Person struct {
	Name string
	Age  int
}

/*
### Usage

To create an instance of a struct, you can either specify the values in the order they are declared or use field names to initialize the struct, which enhances readability:
*/
func Part38() {
	// Using order-based initialization
	person1 := Person{"Alice", 30}

	// Using named fields initialization
	person2 := Person{
		Name: "Bob",
		Age:  25,
	}

	employee := Employee{
		Person: Person{Name: "Eve", Age: 28},
		ID:     1,
	}

	person2.SetAge(26)
	fmt.Println(person2.Age) // Outputs: 26

	fmt.Println(person1.Greet()) // Outputs: Hello, my name is Alice

	fmt.Println(employee.Name) // Outputs: Eve

	var t Talker = Person{"Dave", 34}
	fmt.Println(t.Talk()) // Outputs: Hi!

}

/*
### Methods on Structs

While structs do not inherently contain behavior, you can define methods on them. A method is a function with a special receiver argument. Here’s how you might attach a method to our `Person` struct:

*/

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// This method can be called on instances of `Person`, like so:

/*
### Pointers to Structs

To modify a struct or to avoid copying on each method call, pointers to structs are used. This is crucial for large structs or when method functions need to alter the struct:
*/

func (p *Person) SetAge(age int) {
	p.Age = age
}

/*
### Composition over Inheritance

Go doesn’t support inheritance, but it supports composition. You can include an instance of one struct into another, effectively embedding its fields and methods:
*/

type Employee struct {
	Person // Embedded struct
	ID     int
}

/*
### Interfaces and Structs

Interfaces are used to define behavior. Any struct implementing all methods in the interface is said to implement that interface, enabling polymorphic behavior without inheritance:
*/

type Talker interface {
	Talk() string
}

func (p Person) Talk() string {
	return "Hi!"
}
