package main

import (
	"fmt"
	"strings"
)

// one entry point per app
func main() {
	// declare and initialise a variable

	conferenceName := "Go conference"
	// var conferenceName string = "Go conference"
	// mark a variable as unchanged
   const conferenceTickets =  60
	var remainingTickets uint = 60

	// print out the "typeof" value
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to our " + conferenceName + " booking application 🚀");
	// fmt.Println("Welcome to our ", conferenceName, " booking application 🚀")
	fmt.Printf("Welcome to our %s booking application. 🚀\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend the %s\n", conferenceName)


	 // declare a variable
	 var firstName string
	 var lastName string
    var email string
	 var userTickets uint

	 // 1. arrays
	 // var bookings [50]string

	 // 2. slice
	 var bookings []string //  or var bookings =  []string{}
	 // bookingsArr := []string{}

	 // empty list that will store "firstNames"
	 firstNames := []string{}

	// "for-loop" as an infinite loop
	// no conditions
	for {
    fmt.Println("What is your first name ?: ")
	 // ask for "firstname" input
	 // pointers. we are passing the memory address of the variable
	 fmt.Scan(&firstName)
	 fmt.Println(&firstName) // p the memory address of the variable

	 fmt.Println("What is your last name ?: ")
	 fmt.Scan(&lastName) 	// ask for "lastname" input

	 fmt.Println("What is your email ?: ")
	 fmt.Scan(&email)

	 fmt.Println("How many tickets are you booking ?: ")
	 fmt.Scan(&userTickets)

	 for {
 	    if userTickets > remainingTickets {
	      fmt.Printf("Tickets not sufficient, you may have to order %d or less\n", remainingTickets)

		   fmt.Println("How many tickets are you booking ?: ")
	      fmt.Scan(&userTickets)

		    if userTickets <= remainingTickets {
            break
		    }
	    }
	 }



	 // fmt.Printf("\nThe array: %v", bookings)
	 // fmt.Printf("\nfirst value: %s", bookings[0])
	 // fmt.Printf("\narray type: %T", bookings)
	 // fmt.Printf("\narray length: %d", len(bookings))

	 remainingTickets -= userTickets
	 bookings = append(bookings, firstName + " " + lastName + ",")

	 fmt.Printf("%s %s with email ID: %s, ordered for %d tickets. The number of tickets left is %d 🚀\n",firstName, lastName, email, userTickets, remainingTickets)
	 fmt.Printf("These are all the bookings: %v\n", bookings)

	 // "for-each" loop
	 for _, booking := range bookings {
		 // you can use "strings.Field()" too
		 names := strings.Split(booking, " ")
		 firstNames = append(firstNames, names[0] + ",")
	 }

    fmt.Printf("first name is: %v\n",firstNames)

	 // var ticketsLeft bool = remainingTickets == 0
	 // ticketsLeft := remainingTickets == 0

	 if remainingTickets == 0 {
     fmt.Println("All conference tickets already out")
	  break
	 }

	}

	// will not execute because loop is infinite. unless it terminates
	// fmt.Println("These are all the bookings: \v", bookings)
}


