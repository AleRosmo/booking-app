package main // Application belongs to the package, it's the package name of our app

import (
	"fmt"     // Import package 'format' which is named 'fmt' - Go std. Library, search Go Standard Library for further packages
	"strings" // Import strings
)

func main() { // Entry function
	// You need to tell Go the data type always but when assigning a value and the data type will be in inferred
	// It is possible to define also another data type, more data types: https://www.tutorialspoint.com/go/go_data_types.htm
	// Specifying data types that fits the variable allows you to automatically validate the input

	var conferenceName = "Go Conference" // Variables declared with 'var' at beginning - Variables definied beginning func
	const conferenceTickets = 50         // Static don't use 'var' but 'const'
	var remainingTickets uint = 50

	// Arrays and Slices
	// Arrays have fixed dimension that must be declared, Slices don't need a fixed dimension
	// Arrays and Slices must be declared with a fixed size of how many elements fits indexed from 0..49 and the data type - all values must be same data type
	// var bookings = [50]string{"Ale", "Palle", "Ugo"} // Array with values init, use {} for empty array - .
	// bookings[0] = firstName + " " + lastName         // Assigning values to Array in specific positions
	// var bookings [50]string 							// Array init with no values
	// bookings := []string{} 							// We can use all the ways to declare variables with it's alternatives mixed
	var bookings []string // Slice init with no values and no size

	// 'Syntatic Sugar' is the term to define features in a language that allows to do something more easily
	// Doesn't add functionality that it didn't already have eg.:

	varName := "value" // Variable definition: no 'var' keyword and no type defined, use colon before assignation - No constants, No Types
	fmt.Printf("varName has: '%v' assigned with Go-specific syntax\n", varName)
	fmt.Printf("conferenceTickets is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) // The Placeholder '%T' displays the type

	// Example of prints
	// Print belongs to fmt package.
	// Print() = Inline Print | Println() = Print new line | Printf() Print formatted text

	// fmt.Println("Welcome to", conferenceName, "booking application")                                           // Use ',' to chain and puts automatically whitespace
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available") // Multiple variables can be used
	fmt.Println("Get your ticekts here to attend")

	// Annotation verbs (or placeholders) can be used with 'Printf' function to format variables passed in
	// '%v' is the default format, there are others available to display value
	// All verbs can be found here: https://pkg.go.dev/fmt#hdr-Printing

	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)                                         // Formats a string into the provided args (which are variables or whatever)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets) // Multiple variables can be used also here but variables must be passed in order

	// Loops
	// Go has only 'for' loop
	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string // String unassigned variable initialization - must have the data type 'string' at the end
		var lastName string
		var email string
		var userTickets uint // Integer unassigned variable initialization

		// Pointers or "Special Variables"
		// A pointer is a variable that points to the memory address of another variable.
		// How it looks in memory:
		// 						  	  Memory Address |	Memory
		// 						 	---------------------------
		// Variable: tickets --->			0xc01401 |	'50' 		<--┐
		// 									0xc01402 |	'..'		   | points to
		// Pointer: &tickets --->			0xc01403 |	'0xc01401'  ---┘
		//
		// We reference the memory address of a variable using '&' in front of the variable name
		// and read the value of the user input referencing the pointer in memory for that variable
		fmt.Print("Please enter your first name: ") // Ask user for their name with Scan which gets the User Input
		fmt.Scan(&firstName)
		// We are getting user input passing to 'Scan' function not the Value of 'userName' which is empty but the reference, which will be the memory address
		// The Scan() function will read the user input and assign it to the user variable in memory
		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Please enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// User input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isValidTicketNumber := city == "Singapore" || city == "London"
		// Negation
		// !isValidTicketNumber := city == "Singapore" || city == "London"
		// isInvalidTicketNumber := city != "Singapore" || city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			// Operations on variables - both must be same type 'uint' because math between types 'uint' and 'int' cannot work
			remainingTickets = remainingTickets - userTickets

			// Slice is an abstraction of an Array, variable-lenght or get sub-array of its own, Index-based and have size but resized when needed - it's a Dynamic Array
			bookings = append(bookings, firstName+" "+lastName) // Built-in Go Function that add elements at end of the slice and return updated slice value, grows the slice if more space is needed.

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// For each
			// To do a 'for each' we need to use 'range' which iterates over elements for different data structures (not only array/slices)
			// For arrays and slices 'range' provids index and value for each element

			firstNames := []string{} // Declare new empty list (slice) for first names

			for _, booking := range bookings { // for each, range returns index and value in 1st and 2nd position, we can ditch the index with a Blank Identifier '_'
				var names = strings.Fields(booking)       // "Fields" splits string with whitespace as separator - Return slice with split elements
				firstNames = append(firstNames, names[0]) // Append the name to the slice of first names
			}
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			// If - Else
			// Work as any other language, condition is either True or False (bool)
			if remainingTickets == 0 {
				// End program
				fmt.Println("Our conference is booked out. Come back next year.")
				break // Break exits the loop
			}
		} else { // Must be written like this
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")

			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			continue // Continue exit from this element without executing the code afterwards and continues to the next element
		}

	}

	// Switch case
	// Useful to execute different code for different or specific values when multiple are given
	// This solves the issue of having too many if's for each single case - allows variables to be tested for equality against list of values
	// Default solves the case if no match is found

	city := "London"

	switch city {
	case "New York":
		//code
	case "Signapore":
		//code
	case "London", "Berlin": // Chain with comma, this code will be used for both the cases
		//code
	case "Mexico City":
		//code
	case "Hong Kong":
		//code
	default: // Case when no match is found
		fmt.Print("No valid city selected")
	}

	// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets) // Multiple args print
	// fmt.Println("Empty printing variables to suppress errors.", bookings2)

}

// Examples

// Arrays

// fmt.Printf("The whole array: %v\n", bookings)
// fmt.Printf("The whole Array/Slice: %v\n", bookings)
// fmt.Printf("The first value: %v\n", bookings[0])
// fmt.Printf("Array/Slice type: %T\n", bookings)
// fmt.Printf("Array/Slice lenght: %v\n", len(bookings))

// bookings[0] = firstName + " " + lastName // Assigning values to Array in specific positions

// For / For Each // For conditional

// for {} - This runs forever is like 'for true'
// for index, value := range booking  { }
// for remainingTickets > 0 && len(booking) < 50 { }
