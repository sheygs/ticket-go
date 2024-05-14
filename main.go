package main

import (
	"fmt"
	"strings"
)

// Declare and initialise a variable
const conferenceTickets uint =  60
var remainingTickets uint = 60
var conferenceName = "Go conference"
var bookings []string


func main() {

	greetings()

	for  {
	 firstName,lastName, email, userTickets := getUserInputs();

	 isValidName, isValidEmail, isValidTicketNum := validateUserInput(firstName, lastName, email, userTickets)


     if !isValidName {
	    fmt.Println("firstname or lastname is invalid")
		 break
	  }

	  if !isValidEmail {
       fmt.Println("email is invalid")
		 break
	  }

	  if !isValidTicketNum {
       fmt.Println("ticket number is invalid")
		 break
	  }

      remainingTickets -= userTickets

	   bookings = append(bookings, firstName + " " + lastName)

	   fmt.Printf("%s %s with email ID: %s, ordered for %d tickets. The number of tickets left is %d 🚀\n",firstName, lastName, email, userTickets, remainingTickets)
	   fmt.Printf("These are all the bookings: %v\n", bookings)

	   firstNames := getFirstNames();

      fmt.Printf("first names of bookings are : %v\n",firstNames)

	   if remainingTickets == 0 {
          fmt.Println("All conference tickets already out")
	       break
	   }

	}
}

func greetings(){
	// fmt.Println("Welcome to our " + conferenceName + " booking application 🚀");
	// fmt.Println("Welcome to our ", conferenceName, " booking application 🚀")

 	fmt.Printf("Welcome to our %s booking application. 🚀\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend the %s\n", conferenceName)
}

func getFirstNames() []string {
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){

	 // Input validations
	 isValidName := len(firstName) >= 3 && len(lastName) >= 3

	 isValidEmail := strings.Contains(email,"@")

	 isValidTicketNum := userTickets <= remainingTickets || userTickets > 0

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

func bookTicket(firstName string, lastName string, email string, userTickets uint) []string{
	// do something
	remainingTickets -= userTickets

	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("%s %s with email ID: %s, ordered for %d tickets. The number of tickets left is %d 🚀\n",firstName, lastName, email, userTickets, remainingTickets)
	fmt.Printf("These are all the bookings: %v\n", bookings)

	return bookings
}