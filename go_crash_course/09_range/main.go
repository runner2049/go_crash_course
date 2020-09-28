package main

import "fmt"

func main() {
	ids := []int{33, 76, 68, 98, 132, 7}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	mail := map[string]string{"Bob": "bob@mail.com", "Sarah": "sarah@mail.com"}

	for k, v := range mail {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range mail {
		fmt.Println("Name: " + k)
	}
}
