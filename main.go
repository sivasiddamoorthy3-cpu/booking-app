package main

import (
	"booking-app/helper"
	"booking-app/models"
	"fmt"
)

const conferenceTickets uint = 50

var (
	bookings              = make([]models.UserData, 0)
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
)

func main() {

	// Greeting the Users with welcome message with the conference details
	greetUsers()

	for {
		// Get user input for booking
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user provided info in the request
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputRequest(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(firstName, lastName, email, userTickets)

			// Getting the first names of the booked users.
			firstNames := getFirstNames()
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	// Loop through the bookings and extract the first names.
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Get user input for booking
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Book the tickets based on user provided info
func bookTickets(firstName, lastName, email string, userTickets uint) {
	fmt.Printf("Thank you %v %v for your booking. You have booked %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, userTickets, email)

	// Update remaining tickets after successful booking.
	remainingTickets = remainingTickets - userTickets

	userData := models.UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings for user: %v\n", bookings)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
