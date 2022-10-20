package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketCount {
			bookTicket(firstName, lastName, email, remainingTickets, conferenceName, userTickets, bookings)

			firstNames := getFirstNames(bookings)
			fmt.Printf("Bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("The conference is booked out. Try again next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("invalid email address")
			}
			if !isValidTicketCount {
				fmt.Println("number of tickets entered is invalid")
			}
			continue
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketCount
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of desired tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, remainingTickets uint, conferenceName string, userTickets uint, bookings []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will contact you at your supplied email address %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
