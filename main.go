package main

import (
	"fmt"
	"strings"
)


func main() {
	// Declare and initialise a variable
   const conferenceTickets =  60
	var remainingTickets uint = 60
	conferenceName := "Go conference"

	greetUsers(conferenceName,conferenceTickets, remainingTickets)

	var bookings []string

	// Infinite "for-loop"
	for  {
	 firstName,lastName, email, userTickets := getUserInputs();

	 isValidName, isValidEmail, isValidTicketNum := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)


     if !isValidName {
	    fmt.Println("firstname or lastname is invalid")
	  }

	  if !isValidEmail {
       fmt.Println("email is invalid")
	  }

	  if !isValidTicketNum {
       fmt.Println("ticket number is invalid")
	  }

      remainingTickets -= userTickets

	   bookings = append(bookings, firstName + " " + lastName)

	   fmt.Printf("%s %s with email ID: %s, ordered for %d tickets. The number of tickets left is %d 🚀\n",firstName, lastName, email, userTickets, remainingTickets)
	   fmt.Printf("These are all the bookings: %v\n", bookings)

	   firstNames := getFirstNames(bookings);

      fmt.Printf("first names of bookings are : %v\n",firstNames)

	   if remainingTickets == 0 {
          fmt.Println("All conference tickets already out")
	       break
	   }

	}
}

func greetUsers(conferenceName string, tickets int, remainingTickets uint){
	// fmt.Println("Welcome to our " + conferenceName + " booking application 🚀");
	// fmt.Println("Welcome to our ", conferenceName, " booking application 🚀")

 	fmt.Printf("Welcome to our %s booking application. 🚀\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d remaining tickets.\n", tickets, remainingTickets)
	fmt.Printf("Get your tickets to attend the %s\n", conferenceName)
}

func getFirstNames(bookings []string) []string {
	 // empty list
	 firstNames := []string{}

	 // "for-each" loop
	 for _, booking := range bookings {
		// Alternative - "strings.Field()"
		names := strings.Split(booking, " ")
		firstNames = append(firstNames, names[0])
	 }

	 return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){

	 // Input validations
	 isValidName := len(firstName) >= 3 && len(lastName) >= 3

	 isValidEmail := strings.Contains(email,"@")

	 isValidTicketNum := userTickets <= remainingTickets && userTickets > 0

	 return isValidName, isValidEmail, isValidTicketNum
}

func getUserInputs()(string, string, string, uint){
	 var firstName string
	 var lastName string
    var email string
	 var userTickets uint

	 fmt.Println("What is your first name? :")
	 // Pointers -'&'. we are passing the memory address of the variable
	 fmt.Scan(&firstName)

	 fmt.Println("What is your last name? :")
	 fmt.Scan(&lastName)

	 fmt.Println("What is your email? :")
	 fmt.Scan(&email)

	 fmt.Println("How many tickets are you booking? :")
	 fmt.Scan(&userTickets)

	 return firstName, lastName, email, userTickets
}

func bookTickets(){
	// do something
}