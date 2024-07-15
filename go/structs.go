package main

import "fmt"
import "time"

type user struct {
	userFirstName string
	userLastName string
	userBirthDate string
	createdAt time.Time
}
// function to get user inputs
func (u user)returnUserDetails() {
	fmt.Println(u.userFirstName, u.userBirthDate, u.userLastName)
}
// function to resets user inputs
//when you add an * before youer struct you derefrence

func (u *user)clearUserName() {
	u.userFirstName = ""
	u.userLastName = ""
	if u.userFirstName == "" && u.userLastName == "" {
		fmt.Println("Name is empty")
	}else {
		fmt.Println("Name is not empty")
	}
}
func main () {
	firstname := getUserDetails("What's your firstname: ")
	lastname := getUserDetails("What's your lastname: ")
	birthDate := getUserDetails("what's your birthdate? ")
	

	var appUser user
	appUser = user{
		userFirstName : firstname,
		userLastName : lastname,
		userBirthDate: birthDate,
		createdAt :  time.Now(),
	}
	appUser.returnUserDetails()
	appUser.clearUserName()
	appUser.returnUserDetails()
}





func getUserDetails(promptQuestion string) string {
	fmt.Print(promptQuestion)
	var input string
	fmt.Scan(&input)
	return input
}