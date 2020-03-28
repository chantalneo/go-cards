package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handSize int) (deck, deck) { // The reason why we didn't choose to use a receiver here is because we'd then have e.g. cards.deal(5). It'd give the wrong idea that we're modifying the slice. Instructor will also address more reasons regarding this at the core of next topic
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is a very default permissions that basically means anyone can read and write this file
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // 'err' has a value of type 'error'. If nothing went wrong, it will have a value of 'nil', a.k.a. which essentially means no value in Go
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // 0 would mean success
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// Approach:
	// for each index, card in cards
	//    Generate a random number between 0 and len(cards)-1
	//    Swap the current card and the card at cards[randomNumber]
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d { // We do not need to always get a reference to the element that we are iterating over. So I'm not getting using i, card
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] // Fancy one-line swap right here. This is an expression to swap the elements at both i and new position inside of our slice
	}
}

// Notes:
// 1. (d deck) within syntax... the d is the actual copy of the deck we're working with is available in the function as a variable called 'd'
//    the deck there basically means every variable of type 'deck' can call this function on itself
//    We used d because by convention we usually refer to the receiver with a one or 2 letter abbreviation that matches the type of the receiver. In other OOP, we could see d as self/this
//
// 2. Byte slice is like an array where every element inside of it corresponds to an ASCII character code. String "Hi there!" would be [72 105 32 116 104 101 114 101 33] in byte slice
//    To do conversion, we can refer to the Dec column of http://www.asciitable.com/
//    - We can do type conversion in Go by writing something like: []byte("Hi There!"). []byte refers to the type we want and "Hi There!" is the value we have
//
// 3. Testing in Go isn't RSpec, mocha, jasmine, selenium, etc!
//    - To make a test, create a new file ending in _test.go
// 	    E.g. deck_test.go
//    - To run all tests in a package, run the following command:
//      go test

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

// Quiz 6: Test Your Knowledge: Multiple Return Values
// 1. In the following code snippet, what will the value and type of 'title' and 'pages' be? Title will be string with value of "War and Peace". Pages will be int with value of 1000
/* 	func getBookInfo() (string, int) {
    	return "War and Peace", 1000
	}

	title, pages := getBookInfo() */
// 2. What will the following program log out? "red" "yellow" "blue". Note: We aren't limited to only two return values at a time
/* 	package main

	import "fmt"

	func main() {
    	color1, color2, color3 := colors()

    	fmt.Println(color1, color2, color3)
	}

	func colors() (string, string, string) {
    	return "red", "yellow", "blue"
	} */
// 3. What will the following program log out? "Red is an awesome color"
/* 	package main

	import "fmt"

	func main() {
   		c := color("Red")

   		fmt.Println(c.describe("is an awesome color"))
	}

	type color string

	func (c color) describe(description string) (string) {
   		return string(c) + " " + description
	} */
// 4. Which of the following best explains the describe function listed below? 'describe' is a function with a receiver of type 'color that requires an argument of type 'string', then returns a value of type 'string'
/* 	package main

	import "fmt"

	func main() {
   		c := color("Red")

   		fmt.Println(c.describe("is an awesome color"))
	}

	type color string

	func (c color) describe(description string) (string) {
   		return string(c) + " " + description
	} */
// 5. This is a tricky question about something that we'll cover in much greater detail later on. Think back to our current "cards" program, where we have the following code
// After calling "deal" and passing in "cards", does the list of strings that the "cards" variable point at change?
// In other words, did we modify the 'cards' slice by calling 'deal'? No, 'cards' will be the same before and after calling 'deal'
/* 	func main() {
    	cards := newDeck()

    	hand, remainingCards := deal(cards, 5)

    	hand.print()
    	remainingCards.print()
	} */
