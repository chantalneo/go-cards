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

// Quiz 2: Test Your Knowledge: Variable Assignment
// 1. Is the following a valid way of initializing and assigning a value to a variable? Yes
// var bookTitle string = "Harry Potter"
// 2. Is the following a valid way of initializing an assigning a value to a variable? Yes
// fruitCount := 5
// 3. After running the following code, Go will assume that the variable quizQuestionCount is of what type? Integer
// quizQuestionCount := 10
// 4. Will the following code compile?  Why or why not? No, because a variable can only be initialized one time. In this case, the ':=' operator is being used to intialize 'paperColor' twice
// paperColor := "Green"
// paperColor := "Blue"
// 5. Are the two lines following ways of initializing the variable 'pi' equivalent? Yes
// pi := 3.14
// var pi float = 3.14
// 6. This might require a bit of experimentation on your end :)
// Remember that you can use the Go Playground at https://play.golang.org/ to quickly run a snippet of code
// Is the following code valid? Due to a syntax error: non-declaration statement outside function body that'd occur, it's a no
/* 	package main

	import "fmt"
	deckSize := 20

	func main() {
  		fmt.Println(deckSize)
	} */
// 7. We might be able to initialize a variable and then later assign a variable to it. Is the following code valid? Yes
/*	package main

	import "fmt"

	func main() {
  		var deckSize int
  		deckSize = 52
  	fmt.Println(deckSize)
	} */
// 8. Here's another one that might need some testing on your end!
// Remember that you can use the Go Playground at https://play.golang.org/ to quickly run a snippet of code.
// Is the following code valid? Yes, we can initialize a variable outside of a function, we just can't assign a value to it
/* 	package main

	import "fmt"

	var deckSize int

	func main() {
  		deckSize = 50
  		fmt.Println(deckSize)
	} */
// 9. Is the following code valid? Why or why not? No. This is because, variables must first be initialized with the ':=' operator or the 'var variableName type' syntax
/* 	package main

	import "fmt"

	func main() {
 		deckSize = 52
  		fmt.Println(deckSize)
	} */
