package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Go doesnt support inheritance but it supports composition by embedding structs in structs, which allows you to reuse code and compose behaviors by combining structs.
type customer struct {
	firstName string
	lastName  string
	address   address // composition
	account   account
	pin       uint16
}

type account struct {
	number  uint32
	balance float64
}

type address struct {
	streetNumber uint32
	streetName   string
	unitNumber   string
	city         string
	state        string
	postalcode   string
}

// a function that is associated with a struct is called a method. A method is defined with a receiver argument, which is the struct type (or a pointer to the struct type) that the method is associated with. This receiver allows the method to access and modify the fields of the struct.

// this is still a normal function because it operates on the data passed to it but is not inherently associated with any particular type.
func newCustomerAddress(c *customer, a *address) {
	c.address = *a // a is already a pointer to address, therefor *a dereferences the pointer to assign the value
}

// methods, on the other hand, have a receiver and are tied to a specific type, allowing them to modify the state of that type or access its fields directly.
// this would then allow you to call it on a customer instance, such as cust.newCustomerAddress(&addr).
// this function is assigned and therfor part of the customer struct
func (c *customer) newCustomerAddress(a *address) {
	c.address = *a
}

// This method retrieves and formats the customer's address as a string.
func (c *customer) getCustomerAddress() string {
	addr := c.address
	address := strconv.Itoa(int(addr.streetNumber)) + " " + addr.streetName

	// If the unit number is present, include it in the address
	if addr.unitNumber != "" {
		address += " Unit " + addr.unitNumber
	}
	address += ", " + addr.city + ", " + addr.state + " " + addr.postalcode
	return address
}

func (c *customer) getAccountBalance() int {
	return int(c.account.balance)
}

func (c *customer) setAccountDeposit(amount float64) error {
	if amount > 0 {
		c.account.balance = amount
		return nil
	}
	return errors.New("Error: Invalid amount")
}

func (c *customer) setAccountWithdrawl(amount float64, pin uint16) error {
	if c.pin == 0 {
		return errors.New("Pin Not Set")
	}
	if pin != c.pin {
		return errors.New("Incorrect Pin")
	}
	if amount > c.account.balance {
		return errors.New("Insufficient Funds")
	}
	c.account.balance = c.account.balance - amount
	return nil
}

func (c *customer) setCustomerPin(pin uint16) {
	c.pin = pin
}

func main() {
	fmt.Println("project")

	// this is invalid because this is not specifying a datatype, rather trying to assign a type to a variable. NOT ALLOWED
	// customer0 := customer

	// because we add the {} we are intantiating a customer struct with initialization to primitive defaults values
	customer0 := customer{}
	var customer1 customer

	// this is called a struct literal
	customer2 := customer{"John", "Doe", address{12, "azalea ave", "", "parker", "co", "80134"}, account{1234, 0.00}, 1234}

	customer1.firstName = "stephen"
	customer0.firstName = "angela"

	fmt.Printf("customer0: %s %s\n", customer0.firstName, customer0.lastName)
	fmt.Printf("customer1: %s %s\n", customer1.firstName, customer1.lastName)
	fmt.Printf("customer2: %s %s\n", customer2.firstName, customer2.lastName)

	newAddress := address{16140, "Azalea Ave", "", "Parker", "CO", "80134"}

	customer0.newCustomerAddress(&newAddress)

	fmt.Printf("customer0 adress: %v\n", customer0.getCustomerAddress())
	fmt.Printf("customer0 balance: %v\n", customer0.getAccountBalance())
	fmt.Printf("customer0 deposits $50.00\n")
	customer0.setAccountDeposit(50.00)
	fmt.Printf("customer0 balance: %v\n", customer0.getAccountBalance())

	fmt.Println("customer0 withdrawls 40.00 with a pin code of 0")
	err := customer0.setAccountWithdrawl(40.00, 0)
	if err != nil {
		fmt.Printf("Withdrawl Failed: %s\n", err)
	}

	fmt.Println("customer0 sets pin code to 123")
	customer0.setCustomerPin(123)

	fmt.Println("customer0 withdrawls 40.00 with a pin code of 0")
	err = customer0.setAccountWithdrawl(40.00, 0)
	if err != nil {
		fmt.Printf("Withdrawl Failed: %s\n", err)
	}

	fmt.Printf("customer0 balance: %v\n", customer0.getAccountBalance())

	fmt.Println("customer0 withdrawls 40.00 with a pin code of 123")
	err = customer0.setAccountWithdrawl(40.00, 123)
	if err != nil {
		fmt.Printf("Withdrawl Failed: %s\n", err)
	}

	fmt.Printf("customer0 balance: %v\n", customer0.getAccountBalance())

}
