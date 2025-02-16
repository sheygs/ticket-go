package main

import (
	"fmt"
	shared "ticket-go/utils"
)

const conferenceTickets uint = 60

var remainingTickets uint = 60
var conferenceName = "Go conference"

var bookings = make([]User, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetings()

	for {
		firstName, lastName, email, userTickets, err := getUserInputs()

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		isValidName, isValidEmail, isValidTicketNum := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("`firstname or lastname` is invalid")
			continue
		}

		if !isValidEmail {
			fmt.Println("`email` is invalid")
			continue
		}

		if !isValidTicketNum {
			fmt.Println("`ticket number` is invalid")
			continue
		}

		bookTicket(firstName, lastName, email, userTickets)

		firstNames := getFirstNames()

		fmt.Printf("first names of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All conference tickets already sold out")
			break
		}
	}
}

func greetings() {
	fmt.Printf("Welcome to our %s booking application. 🚀\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend the %s\n", conferenceName)
}

func getFirstNames() []string {
	// empty list
	firstNames := []string{}

	// "for each" loop
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName+",")
	}

	return firstNames
}

func getUserInputs() (firstName, lastName, email string, userTickets uint, err error) {
	fmt.Println("Enter first name: ")
	// Pointers - we're passing the memory address of the variable
	if _, err = fmt.Scan(&firstName); err != nil {
		return "", "", "", 0, err
	}

	fmt.Println("Enter last name: ")
	if _, err = fmt.Scan(&lastName); err != nil {
		return "", "", "", 0, err
	}

	fmt.Println("Enter email: ")
	if _, err = fmt.Scan(&email); err != nil {
		return "", "", "", 0, err
	}

	fmt.Println("How many tickets are you booking?: ")
	if _, err = fmt.Scan(&userTickets); err != nil {
		return "", "", "", 0, err
	}

	return firstName, lastName, email, userTickets, nil
}

func bookTicket(firstName, lastName, email string, userTickets uint) {
	remainingTickets -= userTickets

	user := User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("%s %s with email ID: %s, ordered for %d tickets. The number of tickets left is %d 🚀\n", firstName, lastName, email, userTickets, remainingTickets)
	fmt.Printf("These are all the bookings: %v\n", bookings)
}
