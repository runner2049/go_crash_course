package main

import "fmt"

func greetings(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greetings("Tom"))
	fmt.Println(getSum(1, 4))
}
