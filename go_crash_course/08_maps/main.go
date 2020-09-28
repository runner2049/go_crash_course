package main

import "fmt"

func main() {
	// Defina map
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Tom"] = "tom@gmail.com"
	emails["Dan"] = "dan@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add kv
	mail := map[string]string{"Bob": "bob@mail.com", "Sarah": "sarah@mail.com"}

	fmt.Println(mail)
	fmt.Println(mail["Bob"])
	fmt.Println(len(mail))

}
