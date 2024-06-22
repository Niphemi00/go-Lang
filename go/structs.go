package main

import "fmt"

func hello () {
	userDetails()
}
// introduction to structs


func userDetails() {
	Name := getUserDetails("What's your name: ")
	age := getUserDetails("What is your age as a string: ")
	fmt.Println("Hello, " + Name + "! You are " + age + " years old.")
}

func getUserDetails(promptQuestion string) string {
	fmt.Print(promptQuestion)
	var input string
	fmt.Scan(&input)
	return input
}