// This is a package by itself
// functions must be exported to be used by other packages
// exporting a variable/functions makes it available for all packages in the app
// when possible, use variables as 'local' as possible
//
// In go, to Export a function or anything else (Variables, Structs, Types, etc.) you need to put the first letter as capital
// validateUserInput() = Private | ValidateUserInput() = Public or Exported
package main // Belongs to the same package
import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) { // Public function with capital first letter, can be used by other packages in the module
	// **** OPERATORS & LOGIC ****
	// ** USER INPUT VALIDATION **
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
