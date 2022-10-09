
# Notes

## ARRAYS / SLICES 

 - Arrays have fixed dimension that must be declared
 - Slices don't need a fixed dimension

Arrays and Slices must be declared with a fixed size of how many elements fits indexed from 0..49 and the data type - all values must be same data type

Slices are an abstraction of an Array, variable-lenght or get sub-array of its own, Index-based and have size but resized when needed - it's a Dynamic Array

    var  bookings  = [50]string{"Ale", "Palle", "Ugo"} 
    
Array with values init, use {} for empty array

bookings[0] = firstName +  " "  + lastName // Assigning values to Array in specific positions

var  bookings [50]string  // Array init with no values

bookings  := []string{} // We can use all the ways to declare variables with it's alternatives mixed

var  bookings []string  // Slice init with no values and no size

  

var  bookings  =  make([]map[string]string, 0) // Initializing a list (Slice) of maps creating it with 'make' and specifying a size after , like 0 - Curly braces like Slices here don't work.

// Specifying a size won't matter but is required - being a Slice it will anyways grow in size automatically

  

** APPENDING **

  

bookings  =  append(bookings, firstName+" "+lastName) // Built-in Go Function that add elements at end of the slice and return updated slice value, grows the slice if more space is needed.

  

**** ANNOTATION VERBS (Placeholders) ****

  

Annotation verbs (or placeholders) can be used with 'Printf' function to format variables passed in

'%v' is the default format, there are others available to display value

All verbs can be found here: https://pkg.go.dev/fmt#hdr-Printing

  

**** CHANNELS ****

  

Allow for easy and safe communication between goroutines

Help when goroutines share data or are dependent on each other, mitigating concurrency issues

  

**** CONDITIONS ****

  

If - Else

Work as any other language, condition is either True or False (bool)

  

**** FUNCTIONS ****

func  printFirstNames(bookings []string) {} // Declaring 'booking' slice parameter of type string in input for function

  

** RETURNING VALUES **

  

In Go you also have to define the oputput parameters including the type  explicitly

Return parameteres goes after the input parameters out of the () brackets and must specify only the type

Multiple  return output parameter types must be declared into another set of () brackets after the input parameters

  

func  printFirstNames(bookings []string) []string {} // Declaring 'booking' slice parameter of type string in input and a Slice parameter of type string in output for function

func  validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { // Declaring three outputs of 'bool' data type

return firstBool, secondBool, thirdBool // Returns can be separate with a comma

}

  

** NO INPUT ONLY RETURN **

  

func getUserInput () (string, string, string, uint) {}

  

** ASSIGNING RETURN **

  

isValidName, isValidEmail, isValidTicketNumber  :=  validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

  
  

**** GOROUTINES - THREADING & CONCURRENCY ****

  

We need to instanciate different Threads for things that might take long time

This allows to not block the code if one function takes longer or blocks

To run something in a spearate thread simply use 'go' in front of the statement/expression

'go' abstract the creation of a thread itself

With only 'go' the main thread (the application itself) exits and all the rest stops also.

WaitGroup waits for all the threads launched by the application to finish, it is part of the 'sync'  package

Provides basic sync functionalities

  

WaitGroup has 3 functions

waitGroupName.Add() = Sets number of goroutines to wait (increases the counter by the provided number) - parameter is number of threads, must be excuted before instancing with 'go' and the params is how many functions are instanced with threads in the same block

waitGroupName.Wait() = Wait for all threads specified in 'Add()', waits until counter is 0  - must be executed ad the end of the main thread in the main function

waitGroupName.Done() = Executed in the function that runs a sperate thread to tell it's over with it and the thread can be killed - must be put at end of logic in function

  

**** MAPS ****

  

It's a collection of key, value pairs

Created with keyword 'map' and specifying types for [key] and 'value'

  

map[string]string  // this only defines type of map

  

Maps unique keys to values, you can use the key later to retrieve that value

All keys have same data type - All values have the same data type

  

We need an expression to create the empty map which is 'make', a built-in function

  

make(map[string]string) // Invokes an expression that creates the map

  
  

**** VARIABLES AND CONSTANTS ****

  

You need to tell Go the data type  always but when assigning a value and the data type  will be in inferred

It is possible to define also another data type, more data types: https://www.tutorialspoint.com/go/go_data_types.htm

Specifying data types that fits the variable allows you to automatically validate the input

bookings[0] = firstName +  " "  + lastName // Assigning values to Array in specific positions

  

*** PACKAGE LEVEL VARIABLES ***

Variables outside function defined on top of everything

Globals, basically

Must be defined with 'var' keyword, cannot use ':='

  

**** OPERATORS & LOGIC ****

  

** USER INPUT VALIDATION **

isValidName  :=  len(firstName) >=  2  &&  len(lastName) >=  2

isValidEmail  := strings.Contains(email, "@")

isValidTicketNumber  := userTickets >  0  && userTickets <= remainingTickets

  

** OPERATING ON VARIABLES**

remainingTickets  = remainingTickets - userTickets // Operations on variables - both must be same type 'uint' because math between types 'uint' and 'int' cannot work

isValidTicketNumber  := city ==  "Singapore"  || city ==  "London"

  

** NEGATION **

!isValidTicketNumber  := city ==  "Singapore"  || city ==  "London"

isInvalidTicketNumber  := city !=  "Singapore"  || city !=  "London"

  

**** LOOPS ****

  

Go has only 'for' loop

Types of loops 'for' loops: For / For Each // For conditional

for {} - This runs forever is like 'for true'

for remainingTickets >  0  &&  len(booking) <  50 { }

  

** FOR EACH **

  

for  index, value  :=  range booking { }

To do a 'for each' we need to use 'range' which iterates over elements for different data structures (not only array/slices)

For arrays and slices 'range' provids index and value for each element

** ENDING LOOPS **

  

continue  // Continue exit from this element without executing the code afterwards and continues to the next element

  

break  // Break exits the loop

  

**** SCOPE ****

  

3 levels of scope

Local - Declared within the function or the block - can only be used in that function or block

Package - Declared outside all functions - can be also in different fiels but they all must be part of the same package

Global - Declared outside all functions and First Letter is Uppercase - can be used across all packages

  

**** SWITCH/CASE ****

Useful to execute different code for different or specific values when multiple are given

This solves the issue of having too many if's for each single case - allows variables to be tested for equality against list of values

Default solves the case if no match is found

  

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

  

**** STRINGS ****

  

"Fields" splits string with whitespace as separator - Return slice with split elements

var names = strings.Fields(booking)

  

**** STRUCTS ****

  

Collect different data types of data

Enables to to declare key, value pairs but with mixed data types

Define a custom Type where we can define what properties should have this structure

Like a Class but without inheritance

  

**** POINTERS ****

  

Also called "Special Variables"

A pointer is a variable that points to the memory address of another variable.

How it looks in memory:

Memory Address | Memory

---------------------------

Variable: tickets ---> 0xc01401 | '50' <--┐

0xc01402 | '..' | points to

Pointer: &tickets ---> 0xc01403 | '0xc01401' ---┘

  

We reference the memory address of a variable using '&' in front of the variable name

and read the value of the user input referencing the pointer in memory for that variable

  

fmt.Print("Please enter your first name: ") // Ask user for their name with Scan which gets the User Input

fmt.Scan(&firstName) // Assignment referencing the variable pointer

fmt.Print("Please enter your first name: ") // Ask user for their name with Scan which gets the User Input

  

We are getting user input passing to 'Scan' function not the Value of 'userName' which is empty but the reference, which will be the memory address

The Scan() function will read the user input and assign it to the user variable in memory

  

**** PRINT ****

  

Example of prints

Print belongs to fmt package.

Print() = Inline Print | Println() = Print new line | Printf() Print formatted text

  

fmt.Println("Welcome to", conferenceName, "booking application") // Use ',' to chain and puts automatically whitespace

fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available") // Multiple variables can be used

fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets) // Multiple args print

fmt.Printf("conferenceTickets is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) // The Placeholder '%T' displays the type

fmt.Printf("The whole array: %v\n", bookings)

fmt.Printf("The whole Array/Slice: %v\n", bookings)

fmt.Printf("The first value: %v\n", bookings[0])

fmt.Printf("Array/Slice type: %T\n", bookings)

fmt.Printf("Array/Slice lenght: %v\n", len(bookings))

  

**** SYNTACTIC SUGAR ****

'Syntatic Sugar' is the term to define features in a language that allows to do something more easily

Doesn't add functionality that it didn't already have eg.:

  

varName := "value" // Variable definition: no 'var' keyword and no type defined, use colon before assignation - No constants, No Types

fmt.Printf("varName has: '%v' assigned with Go-specific syntax\n", varName)

  

**** OTHER ****

  

Go has a distinction between "" and ''