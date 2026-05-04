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
		if firstName == "" || lastName == "" || email == "" {
			fmt.Println("First name, last name and email cannot be empty.")
			return
		}

		// Validate the number of tickets requested is greater than 0 and less than or equal to remaining tickets
		if userTickets <= 0 {
			fmt.Println("Number of tickets must be greater than 0.")
			return
		}

		// Validate the number of tickets requested does not exceed remaining tickets.
		// This check is important to prevent overbooking and ensure that we do not sell more tickets than we have available.
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			fmt.Printf("Please try booking again with a valid number of tickets.\n")
			fmt.Printf("Do you want to continue to book again? (yes/no): ")
			var continueBooking string
			fmt.Scan(&continueBooking)
			// If the user does not want to continue booking, exit the program.
			// This allows the user to exit gracefully if they do not want to try booking again after an invalid input.
			if continueBooking != "yes" {
				return
			}

			// If the user wants to continue booking, skip the rest of the loop and start the next iteration to get new input.
			// This allows the user to correct their input without exiting the program.
			continue
		}

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
	}
}
