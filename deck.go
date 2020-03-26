package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string // We're sorta saying type deck is kind of a slice or a subclass or kind of extends all the behavior of a slice of string. Note: extends some class are not terms that we actually use inside of Go

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
