# Video Notes for GoLang
### These notes are all the same program. However each example builds off the previous one.
## Contents
- [Example One](#example-one)
- [Variables and Constants](#variables-and-constants)
- [Data Types](#data-types)
- [Getting User Input Data](#getting-user-input-data)
- [Book Ticket Logic](#book-ticket-logic)
- [Arrays and Slices](#arrays-and-slices)
- [Loops](#loops)
- [If Statement](#if-statement)
- [If Else Statement](#if-else-statement)
- [For Loops Extended](#for-loops-extended)
- [Validating User Input](#validating-user-input)
- [Functions](#functions)
- [Creating Packages](#creating-packages)
- [Maps](#maps)
- [Structs](#structs)
- [Concurrency](#concurrency)


## Example One
``` go
// This declares that the code belongs to the main package.
// In Go, the main package is the default name for a package that makes 
// an executable program, as opposed to a shared library.
package main

// Here we're importing the "fmt" package, which provides 
// functions for formatted I/O operations (like printing to the console).
import "fmt"

// The main function is the entry point of the program.
// Every executable Go program must have this function.
func main() {
    // fmt.Println is a function from the "fmt" package.
    // It prints the string "Hello World" followed by a newline to the console.
    fmt.Println("Hello World")
}

```
## Variables and Constants
``` go
// The code belongs to the main package, which means this will create an executable program.
package main

// Here, the "fmt" package is imported, which provides functions for formatted I/O operations.
import "fmt"

// This is the main function, the entry point of the Go program.
func main() {

    // Declare and initialize a variable named 'conferenceName' with the value "Go Conference".
    var conferenceName string = "Go Conference"
    
    // Declare a constant named 'conferencetTickets' with the value 50.
    // Constants cannot be changed after they are initialized.
    const conferencetTickets int = 50
    
    // Declare and initialize a variable named 'remainingTickets' with the value 50.
    var remainingTickets uint = 50

    // Print a welcome message followed by the conference name.
    fmt.Println("Welcome to", conferenceName, "booking applicaiton")
    
    // This line is commented out. It's an alternative way of printing the welcome message
    // fmt.Printf("Welcome to %v booking application\n", conferenceName)
    
    // Print the total number of tickets and how many are still available.
    fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets, "are still available")
    
    // This line is commented out. It's an alternative way of printing the ticket details
    // fmt.Printf("We have a total of %v tickets and %v are still available\n", conferencetTickets, remainingTickets)
    
    // Print a message prompting users to get their tickets.
    fmt.Println("Get your tickets here to attend")
}
```
## Data Types
#### You need to tell Go Compiler, the data type when declaring the variable
#### Type Inference: BUT, Go can infer the type when you assign a value
``` go
// The code belongs to the main package, which means this will produce an executable program.
package main

// The "fmt" package is imported, which provides functions for formatted I/O operations.
import "fmt"

// This is the main function, the entry point of the Go program.
func main() {

    // Declare and initialize a variable named 'conferenceName' with the value "Go Conference".
    // ":=" is a shorthand for declaration and assignment in Go.
    conferenceName := "Go Conference"
    
    // Declare a constant named 'conferencetTickets' with the value 50.
    // Constants cannot be changed after they're initialized.
    const conferencetTickets = 50
    
    // Declare and initialize a variable named 'remainingTickets' with the value 50.
    var remainingTickets = 50

    // Print the types of the variables 'conferencetTickets', 'remainingTickets', and 'conferenceName'.
    fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferencetTickets, remainingTickets, conferenceName)

    // Print a welcome message followed by the conference name.
    fmt.Println("Welcome to", conferenceName, "booking applicaiton")
    
    // This line is commented out. It's an alternative way of printing the welcome message
    // fmt.Printf("Welcome to %v booking application\n", conferenceName)
    
    // Print the total number of tickets and how many are still available.
    fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets, "are still available")
    
    // This line is commented out. It's an alternative way of printing the ticket details
    // fmt.Printf("We have a total of %v tickets and %v are still available\n", conferencetTickets, remainingTickets)
    
    // Print a message prompting users to get their tickets.
    fmt.Println("Get your tickets here to attend")

    // Declare a variable 'userName' of type string.
    var userName string
    // Declare a variable 'userTickets' of type int.
    var userTickets int
    
    // Prompt the user to enter their first name.
    fmt.Println("Enter your first name: ")
    // Scan for the user's input and store it in the 'userName' variable.
    fmt.Scan(&userName)

    // Assign the value 2 to the 'userTickets' variable.
    userTickets = 2
    // Print the user's name and the number of tickets they booked.
    fmt.Println("User", userName, "booked", userTickets, "tickets.")
}
```
## Getting User Input Data
``` go
// The code belongs to the main package, which means this will create an executable program.
package main

// Here, the "fmt" package is imported, which provides functions for formatted I/O operations.
import "fmt"

// This is the main function, the entry point of the Go program.
func main() {
    // This is a shorthand declaration and assignment of the value "Go Conference" 
    // to the variable 'conferenceName'.
    conferenceName := "Go Conference"
    
    // This line declares a constant named 'conferencetTickets' and assigns it a value of 50.
    // Constants in Go, once set, cannot be changed.
    const conferencetTickets = 50
    
    // This line declares a variable named 'remainingTickets' and initializes it with a value of 50.
    var remainingTickets = 50

    // Using the Printf function from the fmt package, this line prints the types 
    // (%T) of the 'conferencetTickets', 'remainingTickets', and 'conferenceName' variables.
    fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferencetTickets, remainingTickets, conferenceName)

    // This line prints a welcome message followed by the conference name.
    fmt.Println("Welcome to", conferenceName, "booking applicaiton")
    
    // This line prints out how many total tickets there are 
    // and how many of those tickets are still available.
    fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets, "are still available")
    
    // This line encourages users to purchase tickets.
    fmt.Println("Get your tickets here to attend")

    // This declares a string variable 'userName' but doesn't assign it any value yet.
    var userName string
    
    // This declares an integer variable 'userTickets' but doesn't assign it any value yet.
    var userTickets int
    
    // This line prompts the user to enter their first name.
    fmt.Println("Enter your first name: ")
    
    // The Scan function reads the user's input (presumably their name) from the console 
    // and assigns it to the 'userName' variable.
    fmt.Scan(&userName)

    // Assigns a value of 2 to the 'userTickets' variable.
    userTickets = 2
    
    // Prints out a message saying the user (by their provided name) booked 2 tickets.
    fmt.Println("User", userName, "booked", userTickets, "tickets.")
}
```
## Book Ticket Logic
``` go
// The code belongs to the main package, which means this will create an executable program.
package main

// Here, the "fmt" package is imported, which provides functions for formatted I/O operations.
import "fmt"

// This is the main function, the entry point of the Go program.
func main() {

    // Shorthand declaration and assignment of the value "Go Conference" to the variable 'conferenceName'.
    conferenceName := "Go Conference"
    
    // Declaration of a constant named 'conferencetTickets' with the value of 50.
    const conferencetTickets = 50
    
    // Declaration and initialization of the variable 'remainingTickets' with a value of 50.
    // The type of this variable is 'uint', which stands for unsigned integer. It can only have non-negative values.
    var remainingTickets uint = 50

    // Using the Printf function, this line prints the types of the variables 'conferencetTickets', 
    // 'remainingTickets', and 'conferenceName'.
    fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

    // Printing a welcome message followed by the conference name.
    fmt.Println("Welcome to", conferenceName, "booking applicaiton")
    
    // Printing out the total number of tickets and how many are still available.
    fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets, "are still available")
    
    // Printing a message to encourage users to get their tickets.
    fmt.Println("Get your tickets here to attend")

    // Declaring string variables to store user's first name, last name, and email.
    var firstName, lastName, email string
    // Declaring an unsigned integer variable to store the number of tickets a user wants.
    var userTickets uint
    
    // Prompting the user to enter their first name.
    fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)

    // Prompting the user to enter their last name.
    fmt.Println("Enter your last name: ")
    fmt.Scan(&lastName)

    // Prompting the user to enter their email address.
    fmt.Println("Enter your email address: ")
    fmt.Scan(&email)

    // Prompting the user to specify the number of tickets they want.
    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)

    // Updating the 'remainingTickets' variable by subtracting the number of tickets the user has booked.
    remainingTickets = remainingTickets - userTickets

    // Printing a thank you message with the user's name and the number of tickets they booked, 
    // as well as mentioning where the confirmation will be sent.
    fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
    
    // Printing how many tickets are still available for the conference.
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
```
## Arrays and Slices 
#### Slice is an abstraction of an array
#### Slices are more flexible and powerful: variable-length or get a sub-array of its own
#### Slices are also index-based and have a size, but is resized when needed
#### Append adds the elements at the end of the slice
#### Append grows the slice if a greater capacity is needed and returns the updated slice value
``` go
// This code belongs to the main package, which is the default package for Go applications.
package main

// Importing the "fmt" package for formatted I/O operations like printing to the console.
import "fmt"

// The main function, where the execution of the program starts.
func main() {

    // Shorthand declaration and assignment of the string "Go Conference" to the variable 'conferenceName'.
    conferenceName := "Go Conference"
    
    // Declaring a constant named 'conferencetTickets' with a value of 50. It cannot be changed during execution.
    const conferencetTickets = 50
    
    // Declaring and initializing a variable 'remainingTickets' with a value of 50.
    // The type of this variable is 'uint', which stands for unsigned integer.
    var remainingTickets uint = 50

    // Using the Printf function, this line prints the types of the variables 'conferencetTickets',
    // 'remainingTickets', and 'conferenceName'.
    fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

    // Printing a welcome message to the console.
    fmt.Println("Welcome to", conferenceName, "booking applicaiton")
    
    // Printing out the total number of tickets and the number of tickets that are still available.
    fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
    
    // Printing information to encourage users to get their tickets.
    fmt.Println("Get your tickets here to attend")

    // Declaring a slice named 'bookings' to store booked names. Slices are flexible and can grow in size.
    var bookings []string

    // Alternate way to declare and initialize an empty slice. Commented out for reference.
    // bookings := []string{}
    
    // Declaring string variables to capture user input.
    var firstName, lastName, email string
    // Declaring an unsigned integer variable to capture the number of tickets a user wants to book.
    var userTickets uint

    // Prompting the user for input and reading it into the relevant variables.
    fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)
    fmt.Println("Enter your last name: ")
    fmt.Scan(&lastName)
    fmt.Println("Enter your email address: ")
    fmt.Scan(&email)
    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)

    // Updating the 'remainingTickets' variable based on the number of tickets booked by the user.
    remainingTickets = remainingTickets - userTickets
    // Appending the user's name to the 'bookings' slice.
    bookings = append(bookings, firstName+" "+lastName)

    // Various commented-out lines that demonstrate how to print the contents, type, and length of the bookings slice.
    //fmt.Printf("The whole array: %v\n", bookings)
    //fmt.Printf("The whole slice: %v\n", bookings)
    //fmt.Printf("The first value: %v\n", bookings[0])
    //fmt.Printf("The array type: %T\n", bookings)
    //fmt.Printf("The slice type: %T\n", bookings)
    //fmt.Printf("The array length: %v\n", len(bookings))
    //fmt.Printf("The slice length: %v\n", len(bookings))

    // Printing a thank you message and other details about the user's booking.
    fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
		"tickets. You will receive a confirmation email at", email)
    
    // Printing the number of tickets remaining after the user's booking.
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

    // Printing all the names in the 'bookings' slice, which represent the users who've booked tickets.
    fmt.Printf("These are all our bookings: %v\n", bookings)
}
```
## Loops
#### You only have the "for loop" loop
#### Range iterates over elements for different data strictures (so not only arrays and slices) (line 279)
#### string.Fields() splits the string with white space as seperator (line 280)
#### _ Blank identifier to ignore a variable you dont want to use (line 279)
#### with Go you need to make unused variables explicit
``` go
// The main package is the default package for Go applications.
package main

// Importing standard libraries for I/O and string operations.
import (
	"fmt"
	"strings"
)

// The main function is the entry point of the application.
func main() {

	// Defining and initializing some essential variables.
	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	// Displaying the data types of the given variables.
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	// Welcoming users to the booking application.
	fmt.Println("Welcome to", conferenceName, "booking application")

	// Informing users about ticket availability.
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")

	// Prompting users to make a booking.
	fmt.Println("Get your tickets here to attend")

	// Infinite loop to continually accept bookings.
	for {
		// Variables to store user input.
		var firstName, lastName, email string
		var userTickets uint

		// Collecting user details and desired number of tickets.
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Updating the count of available tickets.
		remainingTickets = remainingTickets - userTickets
		
		// Storing the name of the person making the booking.
		bookings = append(bookings, firstName+" "+lastName)

		// Confirming the booking to the user.
		fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
			"tickets. You will receive a confirmation email at", email)

		// Displaying updated ticket availability.
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// Initialize an empty slice to store first names from the bookings.
		firstNames := []string{}

		// Loop through each booking in the bookings slice.
		for _, booking := range bookings {
    
    			// Split the booking string (which presumably has names) into a slice of words.
    			// Typically, the string might look like "First Last", so "Fields" will create ["First", "Last"].
    			names := strings.Fields(booking)

    			// Append the first word (or the first name) from the split result to our firstNames slice.
    			firstNames = append(firstNames, names[0])
		}
		// Displaying all the first names of users who've booked tickets.
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}
}

```
## If Statement
#### Break statement terminates the for loop and continues with the code right after the for loop
#### Continue statement causes loop to skip the remainder of its body and immediately retesting its condition
``` go
// Declare the main package. Every executable Go application must have one 'main' package.
package main

// Import the necessary packages.
import (
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"strings"  // This package provides simple functions to manipulate UTF-8 encoded strings.
)

// main function is the entry point of every Go program.
func main() {

	// Declare and initialize a variable for the name of the conference.
	conferenceName := "Go Conference"

	// Declare a constant that specifies the total number of available conference tickets.
	const conferencetTickets = 50

	// Declare and initialize a variable for the number of remaining tickets.
	var remainingTickets uint = 50  // uint means "unsigned integer". 

	// Declare a slice to store the names of those who've booked. 
	// A slice is a dynamic list in Go that can grow or shrink in size.
	var bookings []string

	// Print the type of each variable using format specifiers.
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	// Print a welcome message to the console.
	fmt.Println("Welcome to", conferenceName, "booking application")

	// Print the number of total and remaining tickets.
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")

	// Another print statement prompting users to get their tickets.
	fmt.Println("Get your tickets here to attend")

	// Start of the infinite loop to continue taking bookings until manually stopped or tickets run out.
	for {
		// Declare string variables to store user input for their personal details.
		var firstName string
		var lastName string
		var email string

		// Declare a uint variable to store the number of tickets the user wants to book.
		var userTickets uint

		// Prompt user for first name and read it into the firstName variable.
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		// Similar prompts and reads for last name, email, and number of tickets.
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Check if the number of tickets user wants to book is more than the remaining tickets.
		if userTickets > remainingTickets {
			// Print a message informing the user they can't book more than the available tickets.
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n",
				remainingTickets, userTickets)
			continue  // This will skip the rest of this loop iteration, taking the user back to input prompts.
		}

		// Deduct the number of tickets user booked from the remaining tickets.
		remainingTickets = remainingTickets - userTickets

		// Add the user's full name to the bookings slice.
		bookings = append(bookings, firstName+" "+lastName)
		
		// Print a thank-you message to the user and inform them of their booking details.
		fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
			"tickets. You will receive a confirmation email at", email)

		// Print how many tickets are still available.
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// Declare a slice to store only the first names of those who've booked.
		firstNames := []string{}

		// Loop through each booking to extract and store first names.
		for _, booking := range bookings {
			names := strings.Fields(booking)  // This splits the full name into separate words.
			firstNames = append(firstNames, names[0])  // Assuming the first word is the first name.
		}

		// Print out all the first names of those who've booked.
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		// Check if there are no tickets left.
		if remainingTickets == 0 {
			// If true, print a message to inform the user and then break out of the infinite loop.
			fmt.Println("Our conference is fully booked. Thank you for your interest!")
			break
		}
	}
}

```
## If Else Statement
#### Or you can do something like this for an is else statement:
``` go
// Declare the main package. Every executable Go application must have one 'main' package.
package main

// Import the necessary packages.
import (
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"strings"  // This package provides simple functions to manipulate UTF-8 encoded strings.
)

// main function is the entry point of every Go program.
func main() {

	// Declare and initialize a variable for the name of the conference.
	conferenceName := "Go Conference"

	// Declare a constant that specifies the total number of available conference tickets.
	const conferencetTickets = 50

	// Declare and initialize a variable for the number of remaining tickets.
	var remainingTickets uint = 50  // uint means "unsigned integer". 

	// Declare a slice to store the names of those who've booked. 
	// A slice is a dynamic list in Go that can grow or shrink in size.
	var bookings []string

	// Print the type of each variable using format specifiers.
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	// Print a welcome message to the console.
	fmt.Println("Welcome to", conferenceName, "booking application")

	// Print the number of total and remaining tickets.
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")

	// Another print statement prompting users to get their tickets.
	fmt.Println("Get your tickets here to attend")

	// Start of the infinite loop to continue taking bookings until manually stopped or tickets run out.
	for {
		// Declare string variables to hold user's first name, last name, and email.
		var firstName string
		var lastName string
		var email string

		// Declare a variable to hold the number of tickets the user wants to book.
		// Using 'uint' to ensure the number of tickets is always non-negative.
		var userTickets uint

		// Prompt the user for their first name.
		fmt.Println("Enter your first name: ")
		// Scan the user's input and store it in the 'firstName' variable.
		fmt.Scan(&firstName)

		// Prompt the user for their last name.
		fmt.Println("Enter your last name: ")
		// Scan the user's input and store it in the 'lastName' variable.
		fmt.Scan(&lastName)

		// Prompt the user for their email address.
		fmt.Println("Enter your email address: ")
		// Scan the user's input and store it in the 'email' variable.
		fmt.Scan(&email)

		// Prompt the user for the number of tickets they wish to book.
		fmt.Println("Enter number of tickets: ")
		// Scan the user's input and store it in the 'userTickets' variable.
		fmt.Scan(&userTickets)

		// Check if the number of tickets the user wants is less than or equal to the number of remaining tickets.
		if userTickets <= remainingTickets {
			// Deduct the booked tickets from the total remaining tickets.
			remainingTickets = remainingTickets - userTickets

			// Add the user's first and last name to the bookings slice.
			bookings = append(bookings, firstName+" "+lastName)

			// Commented-out lines for potentially useful debugging information about the 'bookings' slice.

			// Print a confirmation message thanking the user for their booking.
			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will receive a confirmation email at", email)

			// Notify the user about the number of remaining tickets.
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// Declare a new slice to hold the first names of all attendees.
			firstNames := []string{}

			// Loop through the bookings and extract the first name of each attendee.
			for _, booking := range bookings {
				names := strings.Fields(booking) // Split the full name into individual words.
				firstNames = append(firstNames, names[0]) // Assuming the first word is the first name.
			}

			// Print the list of all attendees' first names.
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// Check if all the tickets have been booked.
			if remainingTickets == 0 {
				// Notify the user that all the tickets have been sold out.
				fmt.Println("Our conference is booked up. Sorry about your luck!")
				break // Exit the loop.
			}

		} else {
			// If the user is trying to book more tickets than what's available, inform them.
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}
}

```
#### You can also do this with the for loop
``` go
// Declare the main package. Every executable Go application must have one 'main' package.
package main

// Import the necessary packages.
import (
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"strings"  // This package provides simple functions to manipulate UTF-8 encoded strings.
)

// main function is the entry point of every Go program.
func main() {
	// Initializing a string variable to store the conference name.
	conferenceName := "Go Conference"

	// Declaring a constant to store the total number of tickets for the conference.
	const conferencetTickets = 50

	// Initializing a uint (unsigned integer) variable to track the number of remaining tickets.
	var remainingTickets uint = 50

	// Declaring a slice to store the names of attendees who have booked tickets.
	var bookings []string

	// Displaying the data types of the defined variables and constants.
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	// Welcoming users to the conference booking application.
	fmt.Println("Welcome to", conferenceName, "booking application")

	// Informing users about the total number of tickets and how many are still available.
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")

	// Informing users they can start the ticket booking process.
	fmt.Println("Get your tickets here to attend")

	// Infinite loop to continuously book tickets until they're sold out or maximum bookings are reached.
	for remainingTickets > 0 && len(bookings) < 50 {

		// Variables to store user input for the booking process.
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Prompting the user for their details and desired number of tickets.
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Checking if the number of tickets requested by the user is available.
		if userTickets <= remainingTickets {
			// Subtracting the booked tickets from the total remaining tickets.
			remainingTickets = remainingTickets - userTickets

			// Adding the user's name to the bookings slice.
			bookings = append(bookings, firstName+" "+lastName)

			// Thanking the user for their booking and providing a confirmation.
			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will receive a confirmation email at", email)

			// Informing users about the number of remaining tickets.
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// Creating a slice to store the first names of all the attendees.
			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			// Displaying the first names of all the attendees.
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// Ending the booking process if all tickets are sold out.
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Sorry about your luck!")
				break
			}
		} else {
			// Informing the user if they've requested more tickets than available.
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}
}
```
## Validating User Input
``` go
// Declare the main package. Every executable Go application must have one 'main' package.
package main

// Import the necessary packages.
import (
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"strings"  // This package provides simple functions to manipulate UTF-8 encoded strings.
)// main function is the entry point of every Go program.
func main() {

	// Define and initialize a string variable for the name of the conference
	conferenceName := "Go Conference"

	// Define a constant integer for the total number of tickets available for the conference
	const conferencetTickets = 50

	// Define and initialize a uint (unsigned integer) variable for the number of tickets remaining
	var remainingTickets uint = 50

	// Declare a slice of strings to keep track of bookings (names of attendees)
	var bookings []string


	// bookings := []string{} This line is an alternative to the one above for initializing the slice

	// Print the type of each variable using printf
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	// Print a welcome message
	fmt.Println("Welcome to", conferenceName, "booking applicaiton")

	// Print the total number of tickets and how many are remaining
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")

	// Inform the user they can book tickets
	fmt.Println("Get your tickets here to attend")

	// Run a loop until all tickets are booked or until the number of bookings reaches 50
	for remainingTickets > 0 && len(bookings) < 50 {

		// Declare variables to store user inputs
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Prompt user for inputs
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Validate user inputs:
		// Check if both first and last name have at least two characters
		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2

		// Check if the email contains the "@" character
		var isValidEmail bool = strings.Contains(email, "@")

		// Check if the number of tickets requested is valid
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		// If all inputs are valid:
		if isValidName && isValidEmail && isValidTicketNumber {

			// Subtract the booked tickets from the remaining tickets
			remainingTickets = remainingTickets - userTickets

			// Add the name of the attendee to the bookings slice
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole array: %v\n", bookings)
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("The array type: %T\n", bookings)
			//fmt.Printf("The slice type: %T\n", bookings)
			//fmt.Printf("The array length: %v\n", len(bookings))
			//fmt.Printf("The slice length: %v\n", len(bookings))

			// Confirm the booking to the user
			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will receive a confirmation email at", email)

			// Display the number of tickets still available
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// Extract first names of all bookings for display
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			// Display all the first names of the attendees
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// Check if all tickets are booked
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			// If inputs are not valid, display error messages accordingly:

			// If name is invalid:
			if !isValidName {
				fmt.Println("I know you're lying about your name. Try again buckaroo!")
			}

			// If email is invalid:
			if !isValidEmail {
				fmt.Println("Aren't emails supposed to have an @ symbol? Where was yours? Try harder next time")
			}

			// If number of tickets is invalid:
			if !isValidTicketNumber {
				fmt.Printf("Now if we only have %v tickets remaining, how do you expect to receive %v..."+
					"Didn't think about that, did ya?\n", remainingTickets, userTickets)
			}
		}
	}
}

```
## Functions in Go
``` go
// Declare the main package. Every executable Go application must have one 'main' package.
package main

// Import the necessary packages.
import (
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"strings"  // This package provides simple functions to manipulate UTF-8 encoded strings.
)// main function is the entry point of every Go program.
func main() {

	// Define and initialize a string variable for the name of the conference
	conferenceName := "Go Conference"

	// Define a constant integer for the total number of tickets available for the conference
	const conferencetTickets = 50

	// Define and initialize a uint (unsigned integer) variable for the number of tickets remaining
	var remainingTickets uint = 50

	// Declare a slice of strings to keep track of bookings (names of attendees)
	var bookings []string

	// Greet users and show the number of tickets available.
	greetUsers(conferenceName, conferencetTickets, remainingTickets)

	// Continue booking while there are tickets available and the number of bookings is less than 50.
	for remainingTickets > 0 && len(bookings) < 50 {
		// Obtain user details and desired ticket number.
		firstName, lastName, email, userTickets := getUserInput()

		// Validate the entered details.
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Check if all entered details are valid.
		if isValidName && isValidEmail && isValidTicketNumber {
			// Update the number of remaining tickets and the bookings list.
			remainingTickets, bookings = bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// Extract the first names from the bookings list and display them.
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// If no tickets are left, display a message and end the loop.
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			// Display appropriate error messages based on the validation results.
			if !isValidName {
				fmt.Println("I know you're lying about your name. Try again buckaroo!")
			}
			if !isValidEmail {
				fmt.Println("Aren't emails supposed to have an @ symbol? Where was yours? Try harder next time")
			}
			if !isValidTicketNumber {
				fmt.Printf("Now if we only have %v tickets remaining, how do you expect to receive %v? Didn't think about that did ya?\n", remainingTickets, userTickets)
			}
		}
	}
}

// Function to greet the user and display available tickets.
func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

// Function to extract and return the first names from the bookings list.
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) // Split the booking string into a slice of words.
		firstNames = append(firstNames, names[0]) // Add the first word (first name) to the list.
	}
	return firstNames
}

// Function to validate the user's input and return the results.
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// Validate that both first and last names have more than 2 characters.
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2

	// Validate that the email contains "@".
	var isValidEmail bool = strings.Contains(email, "@")

	// Validate that the user is requesting a valid number of tickets.
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	
	return isValidName, isValidEmail, isValidTicketNumber
}

// Function to capture user input and return it.
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Obtain inputs from the user.
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Function to update the bookings and remaining tickets.
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) (uint, []string) {
	// Update the number of remaining tickets.
	remainingTickets = remainingTickets - userTickets
	
	// Update the bookings list.
	bookings = append(bookings, firstName+" "+lastName)

	// Display a confirmation message.
	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}
```
## Packages in Go
#### Uses a helper.go package that we create
#### When running this you will need to do: go run main.go helper.go
#### Main go file:
``` go
// Declare the main package. Every executable Go application must have one 'main' package.
package main

// Import the necessary packages.
import (
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"strings"  // This package provides simple functions to manipulate UTF-8 encoded strings.
)
// main function is the entry point of every Go program.
func main() {

	// Variable initialization:
	conferenceName := "Go Conference"            // A string variable for the conference name.
	const conferencetTickets = 50                // A constant value indicating the total number of tickets available.
	var remainingTickets uint = 50               // A uint variable indicating the number of tickets that are still unbooked.
	var bookings []string                        // A slice to store the names of individuals who have booked tickets.

	// Call the greetUsers function to display a welcome message and available tickets.
	greetUsers(conferenceName, conferencetTickets, remainingTickets)

	// Continuously loop to allow ticket booking until no tickets remain or 50 bookings are reached.
	for remainingTickets > 0 && len(bookings) < 50 {

		// Call the getUserInput function to get details from the user.
		firstName, lastName, email, userTickets := getUserInput()

		// Call validateUserInput to ensure the user has provided valid information.
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// If all the provided details are valid:
		if isValidName && isValidEmail && isValidTicketNumber {
			// Update the remainingTickets and bookings list by booking the ticket.
			remainingTickets, bookings = bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// Extract the first names from the bookings and display them.
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// If no tickets remain, inform the user and break out of the loop.
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			// In case the details are not valid, show relevant error messages:

			// Name validation message.
			if !isValidName {
				fmt.Println("I know you're lying about your name. Try again buckaroo!")
			}
			// Email validation message.
			if !isValidEmail {
				fmt.Println("Aren't emails supposed to have an @ symbol? Where was yours? Try harder next time")
			}
			// Ticket number validation message.
			if !isValidTicketNumber {
				fmt.Printf("Now if we only have %v tickets remaining, how do you expect to receive %v? Didn't think about that did ya?\n", remainingTickets, userTickets)
			}
		}
	}
}

// Function to display a welcome message and the number of tickets available.
func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)                  // Display a welcome message.
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available") // Show total and available tickets.
	fmt.Println("Get your tickets here to attend")                               // Prompt to purchase tickets.
}

// Function to extract the first names from the bookings list.
func getFirstNames(bookings []string) []string {
	firstNames := []string{}             // Initialize a slice to hold the first names.
	for _, booking := range bookings {   // Loop through each booking.
		var names = strings.Fields(booking)   // Split the full name into words.
		firstNames = append(firstNames, names[0]) // Add the first name (which is the first word) to the list.
	}
	return firstNames   // Return the list of first names.
}

// Function to gather user details and the number of tickets they want to book.
func getUserInput() (string, string, string, uint) {
	var firstName string      // Declare a variable for the first name.
	var lastName string       // Declare a variable for the last name.
	var email string          // Declare a variable for the email address.
	var userTickets uint      // Declare a variable for the number of tickets.

	// Prompt and capture each detail:
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets   // Return the captured details.
}

// Function to process the ticket booking and update the remaining tickets and bookings list.
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets  // Deduct the booked tickets from the total.
	bookings = append(bookings, firstName+" "+lastName) // Add the user's name to the bookings list.

	// Display a confirmation message to the user.
	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
	// Show the number of tickets still available.
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings  // Return the updated number of tickets and bookings list.
}

```
#### helper.go package:
``` go
package main // This line declares the main package. It indicates that the code belongs to the main module of the Go program.

import "strings" // This line imports the strings package, which provides functions to manipulate and query strings.

// Here's a function named `validateUserInput` which accepts five parameters and returns three boolean values.
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	// Check if the length of both `firstName` and `lastName` is greater than 2.
	// This is a simple check to ensure that the names aren't too short.
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2 

	// Check if the provided `email` contains the "@" character, which is a very basic validation for email addresses.
	var isValidEmail bool = strings.Contains(email, "@")

	// Ensure that the `userTickets` is a positive number (greater than 0) and that the number of tickets the user wants 
	// doesn't exceed the number of `remainingTickets`.
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	// The function returns the three boolean results to indicate the validity of the name, email, and ticket number, respectively.
	return isValidName, isValidEmail, isValidTicketNumber 
}

```
## Maps
``` go
// Declares the main package which is the entry point for an executable Go program.
package main 

// Importing necessary packages.
import (
	"fmt"      // Package for formatted I/O functions.
	"strconv"  // String conversion package.
	"strings"  // Package with string utility functions.
)

// Global constants and variables declaration.
const conferenceTickets int = 50  // The maximum number of tickets available for the conference.
var remainingTickets uint = 50    // The current number of tickets that haven't been booked.
var conferenceName = "Go Conference"  // Name of the conference.
var bookings = make([]map[string]string, 0)  // Slice to hold booking details. Each booking is a map with string keys and values.

// main function is the entry point of every Go program.
func main() {

	greetUsers()  // Function call to display a welcome message to users.

	// This loop runs indefinitely until broken out of.
	for {

		// Fetches user input details: first name, last name, email, and desired number of tickets.
		firstName, lastName, email, userTickets := getUserInput() 
		
		// Validates the fetched user input and returns boolean values indicating the validity of each field.
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// If all user inputs are valid:
		if isValidName && isValidEmail && isValidTicketNumber {

			// Books the ticket for the user and updates global data.
			bookTicket(userTickets, firstName, lastName, email) 

			// Fetches the first names of all attendees and stores them in the 'firstNames' slice.
			firstNames := printFirstNames() 
			// Prints the collected first names.
			fmt.Printf("The first names %v\n", firstNames) 

			// If there are no tickets left, exit the loop.
			if remainingTickets == 0 { 
				break 
			}
		} else {  // If any of the user inputs are not valid:
			// Display error messages based on which inputs were invalid.
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// Move to the next iteration of the loop without processing the rest of the loop body.
			continue 
		}
	}
}

// Function to collect first names from the global 'bookings' slice.
func printFirstNames() []string {
	firstNames := []string{}  // Empty slice to store first names.

	// Loop through each booking in the 'bookings' slice.
	for _, booking := range bookings {
		// Fetch the 'firstName' field from each booking and append it to the 'firstNames' slice.
		firstNames = append(firstNames, booking["firstName"]) 
	}
	return firstNames  // Return the slice containing first names.
}

// Function to get user input from the command line.
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Prompting and getting user input for each field.
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets  // Returns the collected user input.
}

// Validates the user's input.
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// Checks if the length of both first and last names is at least 2 characters.
	isValidName := len(firstName) >= 2 && len(lastName) >= 2  
	// Checks if the email string contains the "@" symbol.
	isValidEmail := strings.Contains(email, "@")  
	// Checks if the number of tickets requested is both positive and not greater than the number of available tickets.
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets  
	return isValidName, isValidEmail, isValidTicketNumber  // Returns the results of the validations.
}

// Welcomes the user and provides information about the available tickets.
func greetUsers() {
	// Displays the welcome message and information about ticket availability.
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

// Books tickets for a user based on their input.
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Reduces the number of remaining tickets by the number the user booked.
	remainingTickets = remainingTickets - userTickets  

	// Creates a map to store the user's details.
	var user = make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email
	// Converts the 'userTickets' value from uint to a string and stores it in the map.
	user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)  

	// Appends the user map to the global 'bookings' slice.
	bookings = append(bookings, user)  
	
	// Prints the updated list of bookings and confirms the user's booking.
	fmt.Printf("List of Bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// Displays the updated number of remaining tickets.
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

```
## Structs
``` go
// Declares the main package, the entry point for an executable Go program.
package main 

// Importing necessary packages.
import (
	"fmt"      // Provides functions for formatted I/O operations.
	"strings"  // Provides functions to perform operations on strings.
)

// Declares a constant for the total number of conference tickets.
const conferenceTickets int = 50 

// Declares global variables.
var remainingTickets uint = 50         // The number of tickets left.
var conferenceName = "Go Conference"   // The name of the conference.
var bookings = make([]User, 0)         // A slice that will store booking details using the User type.

// Defines a new type, User, which is a struct that will store booking details.
type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint  // The number of tickets a user has booked.
}

// The entry function for the program.
func main() {

	// Calls the greetUsers function to display a welcome message.
	greetUsers()

	// Infinite loop to repeatedly prompt the user for input.
	for {

		// Gets the user input through the getUserInput function.
		firstName, lastName, email, userTickets := getUserInput() 
		
		// Validates the user input through the validateUserInput function.
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// Checks if all the user inputs are valid.
		if isValidName && isValidEmail && isValidTicketNumber {

			// Processes the user's request to book tickets.
			bookTicket(userTickets, firstName, lastName, email)

			// Fetches and displays the first names of all users who have booked tickets.
			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			// Checks if there are no more tickets left. If so, exits the loop.
			if remainingTickets == 0 {
				break
			}
		} else {  // If any of the user's inputs are invalid:
			// Error messages are displayed based on which inputs were invalid.
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// Skips the rest of the loop iteration and starts the loop from the beginning.
			continue
		}
	}
}

// Function that returns a slice containing the first names of all users in the 'bookings' slice.
func printFirstNames() []string {
	// Initializes an empty slice to store first names.
	firstNames := []string{}

	// Iterates over each booking in the 'bookings' slice.
	for _, booking := range bookings {
		// Appends the first name from each booking to the 'firstNames' slice.
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames  // Returns the slice containing first names.
}

// Prompts the user for their details and desired number of tickets.
func getUserInput() (string, string, string, uint) {
	// Variables to hold user input.
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Prompts and collects user input for each detail.
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	// Returns the collected user input.
	return firstName, lastName, email, userTickets  
}

// Validates the provided user input and returns results of the validations.
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// Checks if both the first name and last name are at least 2 characters long.
	isValidName := len(firstName) >= 2 && len(lastName) >= 2  
	// Checks if the provided email contains an "@" character.
	isValidEmail := strings.Contains(email, "@")  
	// Checks if the requested number of tickets is both greater than 0 and doesn't exceed available tickets.
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets  

	return isValidName, isValidEmail, isValidTicketNumber  // Returns the results of the validations.
}

// Displays a welcome message to users with details about ticket availability.
func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

// Processes the booking based on valid user input.
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Deducts the number of booked tickets from the total remaining tickets.
	remainingTickets = remainingTickets - userTickets  

	// Initializes a new User struct with the provided details.
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Appends the newly created User struct to the global 'bookings' slice.
	bookings = append(bookings, user)  

	// Confirms the user's booking and provides them with a summary.
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// Updates the user on the number of remaining tickets.
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)  
}

```
## ConCurrency
``` go
// Declares the main package which is the entry point for a Go application.
package main

// Importing necessary packages.
import (
	"fmt"      // Provides I/O functions.
	"strings"  // Provides string manipulation functions.
	"sync"     // Provides synchronization primitives like WaitGroup.
	"time"     // Provides time-related functions.
)

// Declares a constant indicating the total number of tickets for the conference.
const conferenceTickets int = 50

// Declares global variables.
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)   // A slice storing bookings of type UserData.

// Defines the structure for a user's data.
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Declares a WaitGroup to handle synchronization for goroutines.
var wg = sync.WaitGroup{}

// The main function; the starting point of the program.
func main() {

	// Calls the greetUsers function to display introductory information.
	greetUsers()

	// Infinite loop to repeatedly prompt the user for input.
	for {
		// Gets the user input and stores in corresponding variables.
		firstName, lastName, email, userTickets := getUserInput()
		// Validates the provided user details and tickets quantity.
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// Checks if all validations passed.
		if isValidName && isValidEmail && isValidTicketNumber {
			// Books the ticket for the user.
			bookTicket(userTickets, firstName, lastName, email)

			// Increment the counter in the WaitGroup by 1.
			wg.Add(1)
			// Start a goroutine to simulate sending a ticket to the user. This operates asynchronously.
			go sendTicket(userTickets, firstName, lastName, email)

			// Retrieves the list of first names from the bookings.
			firstNames := getFirstNames()
			// Prints the first names.
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Checks if all tickets have been sold out.
			if remainingTickets == 0 {
				// Ends the program after notifying the user.
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else { 
			// If any validations fail, informs the user about the specific issues.
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
	// Waits for all goroutines to complete before proceeding.
	wg.Wait()
}

// Function to greet the user and provide information on the conference and tickets.
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Validates the provided user details.
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2        // Validates the lengths of first and last names.
	var isValidEmail bool = strings.Contains(email, "@")                   // Validates the email format.
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets  // Validates the ticket quantity.
	return isValidName, isValidEmail, isValidTicketNumber
}

// Fetches and returns first names from the bookings slice.
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// Prompts the user for their details and number of tickets.
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Getting user input.
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets  // Returns the user's input.
}

// Books a ticket based on the user's input and updates the global bookings and remainingTickets variables.
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Deducts the number of booked tickets from the available tickets.
	remainingTickets = remainingTickets - userTickets

	// Constructs a UserData struct with the user's details.
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Adds the new booking to the global bookings slice.
	bookings = append(bookings, userData)

	// Notifies the user of their successful booking.
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Simulates the process of sending a ticket to the user.
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)  // Simulates a delay before sending the ticket.
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	// Notifies that the ticket is being sent.
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()  // Decrement the counter in the WaitGroup.
}

```
