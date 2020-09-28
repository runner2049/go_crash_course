package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int 64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var

	var num int = 5
	var isCool = true

	// Shorthand
	//name := "Tom"
	//email := "tom@gmail.com"

	name, email := "Tom", "tomgmail.com"

	size := 1.3

	fmt.Println(name, email, num, isCool)
	fmt.Printf("%T\n", size)
}
