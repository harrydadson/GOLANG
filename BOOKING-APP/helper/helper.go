package helper

import "Strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
