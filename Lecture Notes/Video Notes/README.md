# Video Notes for GoLang

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
    var conferenceName = "Go Conference"
    
    // Declare a constant named 'conferencetTickets' with the value 50.
    // Constants cannot be changed after they are initialized.
    const conferencetTickets = 50
    
    // Declare and initialize a variable named 'remainingTickets' with the value 50.
    var remainingTickets = 50

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
    // fmt.Printf("The whole slice: %v\n", bookings)
    // fmt.Printf("The first value: %v\n", bookings[0])
    // fmt.Printf("The slice type: %T\n", bookings)
    // fmt.Printf("The slice length: %v\n", len(bookings))

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

		//fmt.Printf("The whole array: %v\n", bookings)
		//fmt.Printf("The whole slice: %v\n", bookings)
		//fmt.Printf("The first value: %v\n", bookings[0])
		//fmt.Printf("The array type: %T\n", bookings)
		//fmt.Printf("The slice type: %T\n", bookings)
		//fmt.Printf("The array length: %v\n", len(bookings))
		//fmt.Printf("The slice length: %v\n", len(bookings))
		
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
#### Main go file:
``` go
package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferencetTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets, bookings = bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
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

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

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

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}
```
#### hlper.go package:
``` go
package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
```
## Maps
``` go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create user map
	var user = make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email
	user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, user)
	fmt.Printf("List of Bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
```
## Structs
``` go
package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]User, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create user map
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
```
## ConCurrency
``` go
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
```
