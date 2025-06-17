package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n")

	var firstName string
	var lastName string
	var email string
	var ticket uint8

	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: \n")
	fmt.Scan(&ticket)

	fmt.Printf("Thank you %v %v for buying %v ticket. you will get a confirm mail on your mail %v \n", firstName, lastName, ticket, email)

	remainingTickets = conferenceTickets - ticket

	fmt.Printf("You're buy %v tickets and we have %v tickets remaining.\n", ticket, remainingTickets)

}
