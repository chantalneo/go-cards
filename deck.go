package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string // We're sorta saying type deck is kind of a slice or a subclass or kind of extends all the behavior of a slice of string. Note: extends some class are not terms that we actually use inside of Go

func (d deck) print() { // (d deck) here in the syntax is what we refer to as a receiver. With this line of syntax, we basically allowed any variable of type "deck" the access to the "print" method. Further breakdown of the syntax is available at point 1 below
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Notes:
// 1. (d deck) within syntax... the d is the actual copy of the deck we're working with is available in the function as a variable called 'd'
//    the deck there basically means every variable of type 'deck' can call this function on itself
//    We used d because by convention we usually refer to the receiver with a one or 2 letter abbreviation that matches the type of the receiver. In other OOP, we could see d as self/this

// Quiz 5: Test Your Knowledge: Functions with Receivers
// 1. What would the following code print out? "Harry Potter"
/*	package main

	import "fmt"

	type book string

	func (b book) printTitle() {
    	fmt.Println(b)
	}

	func main() {
    	var b book = "Harry Potter"
    	b.printTitle()
	} */
// 2. Complete the sentence! By creating a new type with a function that has a receiver, we... Are adding a 'method' to any value of that type
// 3. In the following snippet, what does the variable 'ls' represent? A value of laptopSize
/* 	type laptopSize float

	func (ls laptopSize) getSizeOfLaptop() {
    	return ls
	} */
// 4. Is the following code valid? Yes, but it is breaking convention, Go avoids any mention of 'this' or 'self'. While the code is technically valid and will compile, we don't ever reference a receiver value as 'this' or 'self'
/* 	type laptopSize float64

	func (this laptopSize) getSizeOfLaptop() laptopSize {
    	return this
	} */
