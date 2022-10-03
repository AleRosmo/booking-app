package main // Application belongs to the package, it's the package name of our app

import "fmt" // Import package 'format' which is named 'fmt' - Go std. Library, search Go Standard Library for further packages

func main() { // Entry function
	// You need to tell Go the data type always but when assigning a value and the data type will be in inferred
	// It is possible to define also another data type, more data types: https://www.tutorialspoint.com/go/go_data_types.htm
	// Specifying data types that fits the variable allows you to automatically validate the input

	var conferenceName = "Go Conference" // Variables declared with 'var' at beginning - Variables definied beginning func
	const conferenceTickets = 50         // Static don't use 'var' but 'const'
	var remainingTickets uint = 50

	// var bookings = [50]string{"Ale", "Palle", "Ugo"} // Array with values init, use {} for empty array - Arrays must be declared with a fixed size of how many elements fits indexed from 0..49 and the data type - all values must be same data type.
	// bookings[0] = firstName + " " + lastName         // Assigning values to Array in specific positions
	//var bookings [50]string // Array init with no values
	var bookings []string   // Slice init with no values and no size
	bookings2 := []string{} // We can use all the ways to declare variables with it's alternatives mixed

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

	fmt.Printf("Welcome to %v booking application\n", conferenceName)                                           // Formats a string into the provided args (which are variables or whatever)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets) // Multiple variables can be used also here but variables must be passed in order

	var firstName string // String unassigned variable initialization - must have the data type 'string' at the end
	var lastName string
	var email string
	var userTickets uint // Integer unassigned variable initialization

	// Pointers or "Special Variables"
	// A pointer is a variable that points to the memory address of another variable.
	// How it looks in memory:
	// 						  	  Memory Address |	Memory
	// 						 		---------------------------
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

	// Operations on variables - both must be same type 'uint' because math between types 'uint' and 'int' cannot work
	remainingTickets = remainingTickets - userTickets

	// Slice is an abstraction of an Array, variable-lenght or get sub-array of its own, Index-based and have size but resized when needed - it's a Dynamic Array
	bookings = append(bookings, firstName+" "+lastName) // Built-in Go Function that add elements at end of the slice and return updated slice value, grows the slice if more space is needed.

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all the bookings: %v\n", bookings)

	// Loops
	// Go has only 'for' loop

	// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets) // Multiple args print
	fmt.Println("Empty printing variables to suppress errors.", bookings2)

}

// Examples

// Arrays

// fmt.Printf("The whole array: %v\n", bookings)
// fmt.Printf("The whole Array/Slice: %v\n", bookings)
// fmt.Printf("The first value: %v\n", bookings[0])
// fmt.Printf("Array/Slice type: %T\n", bookings)
// fmt.Printf("Array/Slice lenght: %v\n", len(bookings))

// bookings[0] = firstName + " " + lastName // Assigning values to Array in specific positions
