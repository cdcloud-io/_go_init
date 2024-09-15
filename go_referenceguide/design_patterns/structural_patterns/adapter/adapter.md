# Adapter (structural design pattern)

![Adapter](_img/adapter.png)

## References

- [Refactoring.guru](https://refactoring.guru/design-patterns)
- [GeeksForGeeks](https://www.geeksforgeeks.org/adapter-pattern)

## Document status

- [ ] Complete (Verified)
- [ ] Complete (Unverified)
- [x] In Progress

🗓️ Last updated: 09-14-2024

## Summary

- Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.
- Adapter is a structural design pattern.
- AKA: Wrapper

---

## Introduction

Adapter Method or Adapter Design Patterns also knows as wrapper, convert the interface of a class into another interface clients expect. Adapter lets structs (class) work together that couldn’t otherwise because of incompatible interfaces.

## ⚠️ **Important Callouts** ⚠️

- Callout1
  - Sub-point1
  - Sub-point2

- Callout2
  - Sub-point1
  - Sub-point2

---

## Definition(s)

- **Term1**:
  - Definition1
  - Definition2

---

## Usage Guidelines

- When need to ADD functionality without CHANGING your Class/Objects and Methods
- Facilitates Single Responsibility Principle. You can separate the interface or data conversion code from the primary business logic of the program.
- Open/Closed Principle. You can introduce new types of adapters into the program without breaking the existing client code, as long as they work with the adapters through the client interface.
- Connecting your code to a 3rd party library with incompatible interfaces to your existing code.

### When ***NOT*** to use

- Specific business logic within a code base dont always need to have adaters when you control the client and business code. Adapters would just add un needed complexity.

## Components

![Adapter Components](_img/adapter_components.png)

1. **Target Interface**

   - Description: Defines the interface expected and used by the client. It represents the set of operations that the client code can use.
   - Role: It’s the common interface that the client code interacts with.

2. **Adaptee**

   - Description: The existing class or system with an incompatible interface that needs to be integrated into the new system. This is the object (struct instance) used by the `Adapter` to reuse the existing functionality and modify them for desired use.
   - Role: It’s the class or system that the client code cannot directly use due to interface mismatches.

3. **Adapter**

   - Description: A class that implements the target interface and internally uses an instance of the adaptee to make it compatible with the target interface. This is the `Wrapper` which implements the target interface and modifies the specific request available from the Adaptee class.
   - Role: It acts as a bridge, adapting the interface of the adaptee to match the target interface.

4. **Client**

   - Description: The code that uses the target interface to interact with objects. It remains unaware of the specific implementation details of the adaptee and the adapter.  The `Client` will interact with the `Adapter`.

   - Role: It’s the code that benefits from the integration of the adaptee into the system through the adapter.

## Diagrams

- ![Friendly Name](_img/image1.png)

- Description or explanation of the diagram.

## Implementation

- project directory structure:

```sh
apple-charger # Hexagonal Architecture Layout
├── bin
├── cmd
│   └── main.go
├── go.mod
├── internal
│   ├── adapter
│   │   └── adapter.go
│   ├── application
│   │   └── application.go
│   └── client
│       └── client.go
├── Makefile
└── pkg
    └── target
        └── target.go

```

## Code Block

- we have an existing apple charging app, and we need to extend it to support android without changing the existing code.
- the existing apple charging app could be implemented as follows, which also facilitates HEXAGONAL ARCHITECTURE.

```go target.go
package client
// UML - TARGET IMPLEMENTATION
// Note: the `target` implementation is the code that will become the incompatible package to charging `android`
//       this target could be an internal or 3rd party library that we are using for the current version of
//       the app.  Later, down the road, when a customer request adding `android` functionality,
//       we will do so without modifying the current code.


// - `mobile` interface is implemented by apple struct which means we can pass an instance of the `apple`
// struct to the 'chargeMobile(m mobile) method'

// - the `mobile` interface method signature `chargeAppleMobile()` is implemented on the 'client' struct.
// which makes the `client` struct also a `mobile` interface
type mobile interface {
   chargeAppleMobile()
}

// the concrete implementation of the `mobile` interface.

// UML - CONCRETE PROTOTYPE IMPLEMENTATION
type apple struct {}

// create a method / receiver function to implement the mobile interface
// this is the concrete implementation written out
func (a *apple) chargeAppleMobile() {
   fmt.Println("Charging Apple Device")

}

// -------- end: target.go -------- //
```

```go client.go
package client
// UML - CLIENT IMPLEMENTATION (BUSINESS/CORE LOGIC)
// Note(s): - in HEXAGONAL ARCHITECTURE, the BUSINESS/CORE LOGIC should not have any internal or external dependancies

type client struct {}

// client has a method called chargeMobile that accept an 'm' of type 'mobile' defined in the Client

// client receiver function
func (c *client) chargeMobile(m mobile) {
   // when we pass an `apple` instance to `chargeMobile(m mobile)` the (m mobile) becomes (apple)
   // which effectively makes the `m.chargeAppleMobile` --> `apple.chargeAppleMobile` and this
   // is ok because by satisfying the contract. we are essentially calling the `apple` receiver function
   // inside `chargeMobile` as tough we were calling it directly
   m.chargeAppleMobile()
   // effective statement becomes: apple.chargeAppleMobile()
}

// -------- end: client.go -------- //
```

```go application.go
package main
// main program (APPLICATION LOGIC)
// Note(s): - real-world scenario: /cmd/app/main.go should call the `application.go` in /internal/app/
//          - in HEXAGONAL ARCHITECTURE, the APPLICATION LOGIC wrap the CORE LOGIC and gets wrapped by INTERFACE LOGIC
//            which will facilitate the DRIVEN and DRIVING interfaces (check, maybe is Adapters)


// Note: application.go will get additional code in second half when we implement the android mobile phone
func main() {
   apple := &apple{} // create an instance of apple struct
   client := &client{} // create an instance of client struct

   // - call the chargeMobile method of the `client` which accepts a `mobile` interface type
   // - the apple instance is also a `mobile` type as it implements the method signature - the interface contract
   // - we access `chargeMobile` via the .(dot) notation because it is a receiver function (method) of the `client` struct
   // - the receiver fuctions param `(m mobile)` becomes an `apple` struct
   // - *** IMPORTANT *** the `chargeMobile` method parameters will become the apple struct, which gives the
   //                     `chargeMobile` method access to the chargeAppleMobile() method within the apple struct
   client.chargeMobile(apple) 
}

// -------- end: application.go -------- //
```

## Application Output

```sh
// output
// > Charging Apple Device
```

## Implement Android Charging

- option 1: rewrite the existing code to create more abstract charging method that will charge everything.
- option 2: create an ADAPTER in between using the Adapter Pattern, keeping the existing code and functionality unmodified.

### Choosing OPTION 2

- write an adapter that serves the purpose of charging the android (this will take less time, and serve our purpose)

```go adapter.go
package adapter

// UML - ADAPTEE
// This contains the specific functionality or the specific requirement changed functionality reaquested by the client
type android struct {}

func (a *android) chargeAndroidMobile() {
   fmt.Println("Charging Android Device")
}

// UML - ADAPTER
// - This `adapter` will take the input, and convert that `client` input into a form to charge the android device
// - The `adapter`  will be responsible for providing the output back to the `client`
// - The `client` will not know from where this output comes in
// - The job of the adapter is to call the/or charge the android mobile phone

// - create a struct for the android adapter
// - the `androidAdapter` struct also has to implement the `mobile` interface type by satisfying the contract
// - by satisfying the `mobile` interface, this will be used to extend the functionality of the code.
type androidAdapter struct {
   android *android
}

// create the receiver function / method, that satisfies the mobile contract.
func (aa *androidAdapter) chargeAppleMobile() {
   aa.android.chargeAndroidMobile()
}

// -------- end: adapter.go -------- //
```

// modify the existing application.go

```go application.go
package main
// main program (APPLICATION LOGIC)
// Note(s): - real-world scenario: /cmd/app/main.go should call the `application.go` in /internal/app/
//          - in HEXAGONAL ARCHITECTURE, the APPLICATION LOGIC wrap the CORE LOGIC and gets wrapped by INTERFACE LOGIC
//            which will facilitate the DRIVEN and DRIVING interfaces (check, maybe is Adapters)


// Note: application.go will get additional code in second half when we implement the android mobile phone
func main() {
   // The original/inital requirement
   apple := &apple{} // create an instance of apple struct
   client := &client{} // create an instance of client struct

   // - call the chargeMobile method of the `client` which accepts a `mobile` interface type
   // - the apple instance is also a `mobile` type as it implements the method signature - the interface contract
   // - we access `chargeMobile` via the .(dot) notation because it is a receiver function (method) of the `client` struct
   // - the receiver fuctions param `(m mobile)` becomes an `apple` struct
   // - *** IMPORTANT *** the `chargeMobile` method parameters will become the apple struct, which gives the
   //                     `chargeMobile` method access to the chargeAppleMobile() method within the apple struct
   client.chargeMobile(apple)


   // Implement the additional requirement of charging android mobile phone.

   // - create an instance of the `android` struct, and `androidAdapter`
   android := &android{}
   androidAdapter := &androidAdapter{

      // this android adapter expects an `android` so we will give it the `android` instance we just created.
      android: android,

   }

   // now we can use the existing `client` instance on the `androidAdapter` which implements the contract.
   client.chargeMobile(androidAdapter)
}

// -------- end: application.go -------- //




## Additional Explanation(s)

[Hexagonal Architecture Considerations]

Core (Business Logic)

- The client struct represents the core business logic. It depends on abstractions (mobile interface) rather than concrete implementations.

Ports and Adapters

- Port: The mobile interface acts as a port, defining the operations that can be performed.
- Adapter: The apple struct acts as an adapter, implementing the mobile interface.

Isolation

- The core logic (client) is isolated from external dependencies, facilitating easier extension (e.g., adding Android support later without modifying existing code).

---
