# Video Notes for GoLang

## Example One
``` go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
## Variables and Constants
``` go
package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferencetTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
```
## Data Types
#### You need to tell Go Compiler, the data type when declaring the variable
#### Type Inference: BUT, Go can infer the type when you assign a value
``` go
package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
 conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
 "are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Println("User", userName, "booked", userTickets, "tickets.")
}
```
## Getting User Input Data
``` go
package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
 conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
 "are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Println("User", userName, "booked", userTickets, "tickets.")
}
```
## Book Ticket Logic
``` go
package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
		"tickets. You will reviece a confirmation email at", email)
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
package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string
	// bookings := []string{} Same as the line above

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
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
## If Else Statment
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
