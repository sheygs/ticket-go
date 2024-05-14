package main

import (
	"fmt"
	shared "go-web/utils"
)

// Declare and initialise a variable
const conferenceTickets uint =  60
var remainingTickets uint = 60
var conferenceName = "Go conference"

// create an empty list of user struct;
// Initialize the size to 0
var bookings  = make([]User, 0)

type User struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {

	greetings()

	for  {
	 firstName,lastName, email, userTickets := getUserInputs();

	 isValidName, isValidEmail, isValidTicketNum := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


     if !isValidName {
	    fmt.Println("'firstname/lastname' is invalid")
		 break
	  }

	  if !isValidEmail {
       fmt.Println("'email' is invalid")
		 break
	  }

	  if !isValidTicketNum {
       fmt.Println("'ticket number' is invalid")
		 break
	  }

      bookTicket(firstName, lastName, email, userTickets)

	   firstNames := getFirstNames();

      fmt.Printf("first names of bookings are : %v\n",firstNames)

	   if remainingTickets == 0 {
          fmt.Println("All conference tickets already sold out")
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

	 // "for each" loop
	 for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName + ",")
	 }

	 return firstNames
}

func getUserInputs()(string, string, string, uint){
	 var firstName string
	 var lastName string
    var email string
	 var userTickets uint

	 fmt.Println("Enter first name: ")
	 // Pointers - '&'
	 // we're passing the memory address of the variable
	 fmt.Scan(&firstName)

	 fmt.Println("Enter last name: ")
	 fmt.Scan(&lastName)

	 fmt.Println("Enter email: ")
	 fmt.Scan(&email)

	 fmt.Println("How many tickets are you booking ?:")
	 fmt.Scan(&userTickets)

	 return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint){

	remainingTickets =  remainingTickets - userTickets

	user := User {
		firstName: firstName,
		lastName: lastName,
		email:email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("%s %s with email ID: %s, ordered for %d tickets. The number of tickets left is %d 🚀\n",firstName, lastName, email, userTickets, remainingTickets)
	fmt.Printf("These are all the bookings: %v\n", bookings)
}