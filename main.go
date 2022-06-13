package main

import (
	"fmt"
	"sync"
	"time"
)

const conferencaeTickets = 50

var conferenceName string = "Go conference"

var remainingTickets uint = 50

//list of map
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var waitGroup = sync.WaitGroup{}

func main() {

	GreetUser()

	// for {

	firstName, lastName, email, userTicket := getUserInput()
	isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

	if isValidEmail && isValidName && isValidTicket {

		//function print first names
		bookTicket(userTicket, firstName, lastName, email)

		waitGroup.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		// go doSmthn

		firstNames := PrintFirstNames()
		fmt.Printf("The first names are %v bookings\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("We are sorry but we are sold out")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("The first name of last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("The email address you entered is invalid")
		}
		if !isValidTicket {
			fmt.Println("email address you entered is invalid")
		}
	}

	waitGroup.Wait()

}

// }

func GreetUser() {
	fmt.Printf("Welcome to the %v", conferenceName)
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have tota of %v ticket and we have %v remaining\n ", conferencaeTickets, remainingTickets)

}

func PrintFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket

	//create a map for user
	// var userData = make(map[string]string)
	var userData UserData = UserData{firstName: firstName, lastName: lastName, email: email, numberOfTickets: userTicket}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for your booking %v tickets. You will recieve a confirmation email at %v  \n", firstName, lastName, userTicket, email)
	fmt.Printf("We have %v remaining tickets\n", remainingTickets)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v have been sent to %v", userTicket, firstName, lastName, email)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("#####################")
	waitGroup.Done()

}
