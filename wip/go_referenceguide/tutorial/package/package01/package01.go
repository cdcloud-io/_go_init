package package01

import (
	"fmt"
)

var MyString1 string = "I am accessible outside of package main"
var mystring2 string = "I am NOT accesible outside of package main"

func package01(){
	fmt.Println("--------------------------------")
	fmt.Println("package01()")
	fmt.Println("--------------------------------")
}