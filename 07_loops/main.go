package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for f := 1; f <= 10; f++ {
		fmt.Println(f)
	}

	// FizzBuzz
	for a := 1; a <= 100; a++ {
		if a%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if a%3 == 0 {
			fmt.Println("Fizz")
		} else if a%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(a)
		}
	}
}
