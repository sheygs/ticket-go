package shared

import (
	"strings"
)

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (isValidName, isValidEmail, isValidTicketNum bool) {
	isValidName = len(firstName) >= 3 && len(lastName) >= 3

	isValidEmail = strings.Contains(email, "@")

	isValidTicketNum = userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTicketNum
}
