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
	// bookings := []string{} Same as the line above

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		//fmt.Printf("The whole array: %v\n", bookings)
		//fmt.Printf("The whole slice: %v\n", bookings)
		//fmt.Printf("The first value: %v\n", bookings[0])
		//fmt.Printf("The array type: %T\n", bookings)
		//fmt.Printf("The slice type: %T\n", bookings)
		//fmt.Printf("The array length: %v\n", len(bookings))
		//fmt.Printf("The slice length: %v\n", len(bookings))

		fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
			"tickets. You will reviece a confirmation email at", email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}
}
```
## If Else Statements
#### Break statement terminates the for loop and continues with the code right after the for loop
#### Continue statement causes loop to skip the remainder of its body and immediately retesting its condition
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
	// bookings := []string{} Same as the line above

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
------------------------------------------------------------------------------------------------ If Statement
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

if userTickets > remainingTickets {
	fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets dumbass\n",
remainingTickets, userTickets)
	continue
}

remainingTickets = remainingTickets - userTickets
bookings = append(bookings, firstName+" "+lastName)

//fmt.Printf("The whole array: %v\n", bookings)
//fmt.Printf("The whole slice: %v\n", bookings)
//fmt.Printf("The first value: %v\n", bookings[0])
//fmt.Printf("The array type: %T\n", bookings)
//fmt.Printf("The slice type: %T\n", bookings)
//fmt.Printf("The array length: %v\n", len(bookings))
//fmt.Printf("The slice length: %v\n", len(bookings))

fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
	"tickets. You will reviece a confirmation email at", email)
fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

firstNames := []string{}
for _, booking := range bookings {
	var names = strings.Fields(booking)
	firstNames = append(firstNames, names[0])
}
fmt.Printf("These are all our bookings: %v\n", firstNames)

if remainingTickets == 0 {
	// end program
	fmt.Println("Our conference is booked up. Sorry bout ya luck")
	break
}
```
## If Else Statement
``` go
// ------------------------------------------------------------------------------------------------ If Else Statement
// Or you can do something like this for an is else statement:
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole array: %v\n", bookings)
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("The array type: %T\n", bookings)
			//fmt.Printf("The slice type: %T\n", bookings)
			//fmt.Printf("The array length: %v\n", len(bookings))
			//fmt.Printf("The slice length: %v\n", len(bookings))

			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will reviece a confirmation email at", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets dumbass\n",
		remainingTickets, userTickets)
		}
	}
}

You can also do this with the for loop
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
	// bookings := []string{} Same as the line above

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings)< 50{
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole array: %v\n", bookings)
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("The array type: %T\n", bookings)
			//fmt.Printf("The slice type: %T\n", bookings)
			//fmt.Printf("The array length: %v\n", len(bookings))
			//fmt.Printf("The slice length: %v\n", len(bookings))

			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will reviece a confirmation email at", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets dumbass\n",
		remainingTickets, userTickets)
		}
	}
}
```
## Validating User Input
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
	// bookings := []string{} Same as the line above

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
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

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole array: %v\n", bookings)
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("The array type: %T\n", bookings)
			//fmt.Printf("The slice type: %T\n", bookings)
			//fmt.Printf("The array length: %v\n", len(bookings))
			//fmt.Printf("The slice length: %v\n", len(bookings))

			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will reviece a confirmation email at", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("I know your lying about your name. Try again buckaroo!")
			}
			if !isValidEmail {
				fmt.Println("Aren't emails suppose to have and @ symbol. Where was yours? Try harder next time")
			}
			if !isValidTicketNumber {
				fmt.Printf("Now if we only have %v tickets remaining, how do you expect to recieve %v..."+
					"Didnt think about that did ya?\n", remainingTickets, userTickets)
			}
		}
	}
}
```
## Functions in Go
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
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
