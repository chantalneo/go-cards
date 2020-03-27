package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string // We're sorta saying type deck is kind of a slice or a subclass or kind of extends all the behavior of a slice of string. Note: extends some class are not terms that we actually use inside of Go

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // append function does not modify the existing slice. Instead it returns a new slice that we then assign back to a variable of cards
		}
	}

	return cards
}

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
//
// 2. Complete the sentence! By creating a new type with a function that has a receiver, we... Are adding a 'method' to any value of that type
//
// 3. In the following snippet, what does the variable 'ls' represent? A value of laptopSize
/* 	type laptopSize float

	func (ls laptopSize) getSizeOfLaptop() {
    	return ls
	} */
//
// 4. Is the following code valid? Yes, but it is breaking convention, Go avoids any mention of 'this' or 'self'. While the code is technically valid and will compile, we don't ever reference a receiver value as 'this' or 'self'
/* 	type laptopSize float64

	func (this laptopSize) getSizeOfLaptop() laptopSize {
    	return this
	} */
//
// 5. In the for loops, we replaced i and j with underscores to tell Go that we understand that there are variables there but we just don't care about it and we don't want to use it. Thus, getting rid of errors
//
// 6. Slices are zero-indexed. Go has a little helper built in to the basic syntax of the language to help us select slices. The syntax is as follows:
// - cards[startIndexIncluding : upToNotIncluding]
//   At such if we call cards[0:2], we get a slice of cards containing cards[0] and cards[1] only
// - We can optionally leave off the numbers on either side of the colon to have go automatically infer that we want to start from the beginning or go all the way to the end of a slice
//   At such cards[:2] is as good as cards[0:2]
//   And if I use [2:], it means I'll take cards from cards[2] to the end of the cards slice
