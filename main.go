package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Breaking down of the syntax, check point 1 below
	card := "Ace of Spades"   // This is the 100% equivalent of the code right above. Here we rely upon the compiler to infer that card is supposed to contain a string
	card = "Five of Diamonds" // Important note on point 3 below
	fmt.Println(card)
}

// Notes:
// 1. var card string = "Ace of Spades"
//    var tells go we're about to create a new variable. var is simply the short of variable
//    card is the name of the variable
//    string tells go that only a "string" will ever be assigned to this variable. Go is a statically typed language. So whenever we define a variable we assign it a type
//    "Ace of Spades" is basically the value assigned to this variable "card" using an equal sign
//
// 2. Basic Go Types (this is not an exhaustive list of basic types by the way)
//    Type    | Example
//    bool    | true, false
//    string  | "Hi!", "How's it going?"
//    int     | 0, -10000, 99999
//    float64 | 10.000001, 0.00009, -100.003
//
// 3. Only use := syntax when you're defining a new variable. To resign a new value to in this case, our variable "card", simply use the following:
//    card = "Five of Diamonds"
