package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitSlice := []string{"Apple", "Orange", "Lemon"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
