# Receiver Function (Method)

A **receiver function** refers to a method that is associated with a specific type. In Go, you can define methods on types (both struct types and named types) by specifying a **receiver**. The receiver appears in the method declaration and represents the instance of the type on which the method is called.

Here's an overview of what receiver functions are and how they work in Go:

---

## **What Is a Receiver in Go?**

A **receiver** is a parameter in a method definition that binds the method to a specific type. It appears between the `func` keyword and the method name, enclosed in parentheses.

### **Syntax:**

```go
func (receiver Type) MethodName(parameters) returnType {
    // method body
}
```

- **receiver**: A variable name representing the instance of the type.
- **Type**: The type to which the method is bound.
- **MethodName**: The name of the method.
- **parameters**: Optional input parameters.
- **returnType**: Optional return type(s).

---

## **Types of Receivers**

There are two types of receivers in Go:

1. **Value Receivers**: The receiver is a copy of the value.
2. **Pointer Receivers**: The receiver is a pointer to the value.

### **Value Receivers**

When you use a value receiver, the method operates on a copy of the value passed to it. This means any modifications to the receiver inside the method do not affect the original value.

**Example:**

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### **Pointer Receivers**

Pointer receivers allow the method to modify the original value. It's also more efficient for large structs because it avoids copying the entire struct.

**Example:**

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

---

## **When to Use Pointer vs. Value Receivers**

- **Use Pointer Receivers When:**
  - The method needs to modify the receiver's value.
  - The struct is large, and you want to avoid copying it.
  - Consistency: If some methods need pointer receivers, it's idiomatic to use them for all methods of the type.

- **Use Value Receivers When:**
  - The method doesn't need to modify the receiver's value.
  - The receiver is a small, copyable value (like basic types).

---

## **Example: Using Receiver Functions in Go**

Let's put it all together with a complete example.

```go
package main

import "fmt"

// Define a struct type
type Circle struct {
    Radius float64
}

// Method with a value receiver
func (c Circle) Area() float64 {
    return 3.1416 * c.Radius * c.Radius
}

// Method with a pointer receiver
func (c *Circle) SetRadius(r float64) {
    c.Radius = r
}

func main() {
    // Create an instance of Circle
    c := Circle{Radius: 5}

    // Call method with value receiver
    fmt.Println("Area:", c.Area()) // Output: Area: 78.54

    // Call method with pointer receiver
    c.SetRadius(10)
    fmt.Println("New Radius:", c.Radius) // Output: New Radius: 10
    fmt.Println("New Area:", c.Area())   // Output: New Area: 314.16
}
```

**Explanation:**

- `Area()` is a method with a value receiver; it calculates the area without modifying the circle.
- `SetRadius()` is a method with a pointer receiver; it modifies the radius of the circle.

---

## **Key Points**

- **Methods vs. Functions**: In Go, a method is a function with a receiver argument.
- **Receivers and Interfaces**: Methods with receivers allow types to satisfy interfaces, enabling polymorphism.
- **Receiver Naming**: It's common to use short, lowercase names (like `r` for receivers) to keep code concise.

---

## **Benefits of Using Receiver Functions**

- **Encapsulation**: Methods allow you to encapsulate behavior related to a type.
- **Code Organization**: Grouping methods with their associated types improves code readability.
- **Reusability**: Methods can be reused across different parts of the program where the type is used.

---

## **Common Use Cases**

- **Struct Methods**: Adding behavior to structs, like calculations or transformations.
- **Custom Types**: Defining methods on custom types to implement interfaces.
- **Immutable vs. Mutable Types**: Deciding whether a type should be immutable (value receiver) or mutable (pointer receiver).

---

## **Understanding Method Sets**

The set of methods that a type has is known as its **method set**.

- **For a type `T`**:
  - The method set includes all methods with receiver type `T`.
- **For a pointer to `T` (`*T`)**:
  - The method set includes all methods with receiver type `*T` **and** `T`.

This means that if you have a method with a pointer receiver, you can still call it on a value of type `T` because Go will automatically take the address of `T`.

**Example:**

```go
type Counter int

func (c *Counter) Increment() {
    *c++
}

func main() {
    var count Counter
    count.Increment() // Go treats this as (&count).Increment()
    fmt.Println(count) // Output: 1
}
```

---

## **Conclusion**

Receiver functions (methods) in Go are a powerful feature that allows you to associate functions with types, enabling object-oriented programming principles like encapsulation and polymorphism. Understanding how to use value and pointer receivers effectively is crucial for writing idiomatic and efficient Go code.

---

**Further Reading:**

- [Go Language Specification - Method Declarations](https://golang.org/ref/spec#Method_declarations)
- [Effective Go - Methods](https://golang.org/doc/effective_go.html#methods)
- [Go by Example - Methods](https://gobyexample.com/methods)
