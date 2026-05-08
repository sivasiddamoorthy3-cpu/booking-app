package helper

import "strings"

func ValidateUserInputRequest(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	// Validate user input before booking
	// Validate the number of tickets requested is greater than 0 and less than or equal to remaining tickets
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := len(email) >= 6 && strings.Contains(email, "@")
	// Validate the number of tickets requested does not exceed remaining tickets.
	// This check is important to prevent overbooking and ensure that we do not sell more tickets than we have available.
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
