package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	for {
		// Get user input for booking
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		// Validate user input before booking
		// Validate the number of tickets requested is greater than 0 and less than or equal to remaining tickets
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := len(email) >= 6 && strings.Contains(email, "@")
		// Validate the number of tickets requested does not exceed remaining tickets.
		// This check is important to prevent overbooking and ensure that we do not sell more tickets than we have available.
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Printf("Thank you %v %v for your booking. You have booked %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, userTickets, email)

			// Update remaining tickets after successful booking.
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			// Loop through the bookings and extract the first names.
			for _, booking := range bookings {
				// Split the booking string into first name and last name using strings.Fields, which splits the string by whitespace.
				// This allows us to easily extract the first name, which is the first element of the resulting slice.
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of all bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered is too short or doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}
